package logs

import (
	"fmt"
	"os"
	"strings"

	"github.com/op/go-logging"
)

// ==============================================================================================================================
//                                      LOGS
// ==============================================================================================================================
var (
	modulename string = "iDirectStat"
	Log = logging.MustGetLogger(modulename)
	LogPrinter *logPrinterType
	Debug bool = false
	Verbose bool = false
)

func Init(debug, verbose bool) error {
	Debug, Verbose = debug, verbose
	return nil
}

type Password string

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}



func init() {
	var syslogformat = logging.MustStringFormatter(
		"[%{shortpkg}: %{shortfile}: %{shortfunc}] %{message}",
	)
	syslog, _ := logging.NewSyslogBackend(modulename)
	syslogF := logging.NewBackendFormatter(syslog, syslogformat)
	syslogL := logging.AddModuleLevel(syslogF)
	syslogL.SetLevel(logging.WARNING, "")

	var logformat = logging.MustStringFormatter(
		"%{color}%{time:15:04:05} [%{shortpkg}: %{shortfile}: %{shortfunc}()] ▶ %{level:.4s} ◀  %{color:reset} %{message}",
	)
	console := logging.NewLogBackend(os.Stderr, "", 0)
	consoleF := logging.NewBackendFormatter(console, logformat)

	logging.SetBackend(syslogL, consoleF)

	LogPrinter = NewLogPrinter()
	go LogPrinter.run()
}


func LogDebug(format string, args ...interface{}) {
	if Verbose {
		Log.Debug(format, args...)
	} else if Debug {
		Log.Debug(format, args...)
	}
}

// ==============================================================================================================================
//                                      CONSOLE
// ==============================================================================================================================

type LogString struct {
	raw string
	fmt string
}

func (l *LogString) AddF(format string, args ...interface{}) {
	l.raw = l.raw + fmt.Sprintf(format, args...)
}

func (l *LogString) AddS(s string) {
	l.raw = l.raw + s 
}

func (l *LogString) AddSR(s string) {
	l.raw = l.raw + s + "\n" 
}

func (l *LogString) Box(w int) string {
	out := "\n"
	ss := strings.Split(l.raw, "\n")
	ls := len(ss)
	for i, ln := range ss {
		if i == 0 {
			x := ((w - len(ln)) / 2) - 1
			out += fmt.Sprintf("\u2554%s %s %s\n", strings.Repeat("\u2550", x), ln, strings.Repeat("\u2550", x))
		} else if i == (ls - 1) && len(ln) == 0 {
			continue
		} else {
			out += fmt.Sprintf("\u2551%s\n", strings.Replace(ln, "\n", "\n\u2551", -1))
		}
	}
	out += fmt.Sprintf("\u255A%s\n", strings.Repeat("\u2550", w))
	l.fmt = out
	return l.fmt
}

func (l *LogString) BCon(w int) {
	LogPrinter.Con(l.Box(w))
}

func (l *LogString) Raw() string {
	return l.raw
}



type logPrinterType struct {
	todo chan string
}

func NewLogPrinter() (*logPrinterType) {
	Log.Debug("NewLogPrinter()... ")
	l := new(logPrinterType)
	l.todo = make(chan string, 100)
	return l
}

func (l *logPrinterType) Con(s string) {
	l.todo <- s
}

func (l *logPrinterType) run() {
	Log.Debug("logPrinterType.run()... ")
	for msg := range l.todo {
		fmt.Println(msg)
	}
}


