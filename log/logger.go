package log

import (
	L "log"
	"os"
)

const (
	// Bits or'ed together to control what's printed. There is no control over the
	// order they appear (the order listed here) or the format they present (as
	// described in the comments).  A colon appears after these items:
	//	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	Ldate         = 1 << iota     // the date: 2009/01/23
	Ltime                         // the time: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

type LvlLogger struct {
	loggers []*L.Logger
	lvl     int
}

var Out *LvlLogger
var Err *LvlLogger

func init() {
	Out.loggers = make([]*L.Logger, 10)
	Err.loggers = make([]*L.Logger, 10)
	Out.SetLevel(1)
	Err.SetLevel(1)
	for i, _ := range Out.loggers {
		Out.loggers[i] = L.New(os.Stdout, "", 0)
		Err.loggers[i] = L.New(os.Stderr, "", 0)
	}
}

func (l *LvlLogger) Print(n int, v ...interface{}) {
	if n < l.lvl {
		l.loggers[n].Print(v)
	}
}

func (l *LvlLogger) Printf(n int, format string, v ...interface{}) {
	if n < l.lvl {
		l.loggers[n].Printf(format, v)
	}
}

func (l *LvlLogger) Println(n int, v ...interface{}) {
	if n < l.lvl {
		l.loggers[n].Println(v)
	}
}

func (l *LvlLogger) SetLevel(n int) {
	l.lvl = n
}

func (l *LvlLogger) Level() int {
	return l.lvl
}
func (l *LvlLogger) SetFlags(n, flag int) {
	l.loggers[n].SetFlags(flag)
}

func (l *LvlLogger) Flags(n int) int {
	return l.loggers[n].Flags()
}

func (l *LvlLogger) SetPrefix(n int, prefix string) {
	l.loggers[n].SetPrefix(prefix)
}

func (l *LvlLogger) Prefix(n int) string {
	return l.loggers[n].Prefix()
}
