// Package name declaration
package logger

// Import packages
import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
	"sync/atomic"
	"time"
)

var (
	// Map for te various codes of colors
	colors map[string]string

	// Contains color strings for stdout
	logNo uint64
)

// Color numbers for stdout
const (
	Black = (iota + 30)
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// Worker class, Worker is a log object used to log messages and Color specifies
// if colored output is to be produced
type Worker struct {
	Minion *log.Logger
	Color  int
}

// Info class, Contains all the info on what has to logged, time is the current time, Module is the specific module
// For which we are logging, level is the state, importance and type of message logged,
// Message contains the string to be logged, format is the format of string to be passed to sprintf
type Info struct {
	Id       uint64
	GID      uint64
	Time     string
	Module   string
	Level    string
	Line     int
	Filename string
	Message  string
	format   string
}

// Logger class that is an interface to user to log messages, Module is the module for which we are testing
// worker is variable of Worker class that is used in bottom layers to log the message
type Logger struct {
	Module string
	worker *Worker
}

// Returns a proper string to be outputted for a particular info
func (r *Info) Output() string {
	msg := fmt.Sprintf(r.format, r.Id, r.Time, r.Filename, r.Line, r.Level, r.GID, r.Message)
	return msg
}

// Returns an instance of worker class, prefix is the string attached to every log,
// flag determine the log params, color parameters verifies whether we need colored outputs or not
func NewWorker(prefix string, flag int, color int, out io.Writer) *Worker {
	return &Worker{Minion: log.New(out, prefix, flag), Color: color}
}

// Function of Worker class to log a string based on level
func (w *Worker) Log(level string, calldepth int, info *Info) error {
	if w.Color != 0 {
		buf := &bytes.Buffer{}
		buf.Write([]byte(colors[level]))
		buf.Write([]byte(info.Output()))
		buf.Write([]byte("\033[0m"))
		return w.Minion.Output(calldepth+1, buf.String())
	} else {
		return w.Minion.Output(calldepth+1, info.Output())
	}
}

// Returns a proper string to output for colored logging
func colorString(color int) string {
	return fmt.Sprintf("\033[%dm", int(color))
}

// Initializes the map of colors
func initColors() {
	colors = map[string]string{
		"CRITICAL": colorString(Magenta),
		"ERROR":    colorString(Red),
		"WARNING":  colorString(Yellow),
		"NOTICE":   colorString(Green),
		"DEBUG":    colorString(Cyan),
		"INFO":     colorString(White),
	}
}

// Returns a new instance of logger class, module is the specific module for which we are logging
// , color defines whether the output is to be colored or not, out is instance of type io.Writer defaults
// to os.Stderr
func New(args ...interface{}) (*Logger, error) {
	initColors()

	var module string = "DEFAULT"
	var color int = 1
	var out io.Writer = os.Stderr

	for _, arg := range args {
		switch t := arg.(type) {
		case string:
			module = t
		case int:
			color = t
		case io.Writer:
			out = t
		default:
			panic("logger: Unknown argument")
		}
	}
	newWorker := NewWorker("", 0, color, out)
	return &Logger{Module: module, worker: newWorker}, nil
}

// The log commnand is the function available to user to log message, lvl specifies
// the degree of the messagethe user wants to log, message is the info user wants to log
func (l *Logger) Log(lvl string, message interface{}) {
	l.log_internal(lvl, message, 2)
}

func InterfacesToString(raw_slice ...interface{}) string {
	message := ""
	for _, raw := range raw_slice {
		switch v := raw.(type) {
		case string:
			message += v
		default:
			message += fmt.Sprint(v)
		}
		message += " "
	}
	return message
}

func (l *Logger) log_internal(lvl string, raw_message interface{}, pos int) {
	message := InterfacesToString(raw_message)
	var formatString string = "#%d %s %s:%d\t ▶ %.3s{%d} %s"
	_, filename, line, _ := runtime.Caller(pos)
	filename = path.Base(filename)
	info := &Info{
		Id:       atomic.AddUint64(&logNo, 1),
		GID:      getGID(),
		Time:     time.Now().Format("2006-01-02 15:04:05"),
		Module:   l.Module,
		Level:    lvl,
		Message:  message,
		Filename: filename,
		Line:     line,
		format:   formatString,
	}
	l.worker.Log(lvl, 2, info)
}

// Fatal is just like func l.Critical logger except that it is followed by exit to program
func (l *Logger) Fatal(messages ...interface{}) {
	l.log_internal("CRITICAL", InterfacesToString(messages...), 2)
	os.Exit(1)
}

// FatalF is just like func l.CriticalF logger except that it is followed by exit to program
func (l *Logger) FatalF(format string, a ...interface{}) {
	l.log_internal("CRITICAL", fmt.Sprintf(format, a...), 2)
	os.Exit(1)
}

// FatalNF is just like FatalF, but the n parameter indicates how many stack levels to go back when printing file name and line number info
func (l *Logger) FatalNF(n int, format string, a ...interface{}) {
	l.log_internal("CRITICAL", fmt.Sprintf(format, a...), 2+n)
	os.Exit(1)
}

// Panic is just like func l.Critical except that it is followed by a call to panic
func (l *Logger) Panic(messages ...interface{}) {
	l.log_internal("CRITICAL", InterfacesToString(messages...), 2)
	panic(InterfacesToString(messages...))
}

// PanicF is just like func l.CriticalF except that it is followed by a call to panic
func (l *Logger) PanicF(format string, a ...interface{}) {
	l.log_internal("CRITICAL", fmt.Sprintf(format, a...), 2)
	panic(fmt.Sprintf(format, a...))
}

// PanicNF is just like PanicF, but the n parameter indicates how many stack levels to go back when printing file name and line number info
func (l *Logger) PanicNF(n int, format string, a ...interface{}) {
	l.log_internal("CRITICAL", fmt.Sprintf(format, a...), 2+n)
	panic(fmt.Sprintf(format, a...))
}

// Critical logs a message at a Critical Level
func (l *Logger) Critical(messages ...interface{}) {
	l.log_internal("CRITICAL", InterfacesToString(messages...), 2)
}

// CriticalF logs a message at Critical level using the same syntax and options as fmt.Printf
func (l *Logger) CriticalF(format string, a ...interface{}) {
	l.log_internal("CRITICAL", fmt.Sprintf(format, a...), 2)
}

// CriticalNF is just like CriticalF, but the n parameter indicates how many stack levels to go back when printing file name and line number info
func (l *Logger) CriticalNF(n int, format string, a ...interface{}) {
	l.log_internal("CRITICAL", fmt.Sprintf(format, a...), 2+n)
}

// Error logs a message at Error level
func (l *Logger) Error(messages ...interface{}) {
	l.log_internal("ERROR", InterfacesToString(messages...), 2)
}

// ErrorF logs a message at Error level using the same syntax and options as fmt.Printf
func (l *Logger) ErrorF(format string, a ...interface{}) {
	l.log_internal("ERROR", fmt.Sprintf(format, a...), 2)
}

// ErrorNF is just like ErrorF, but the n parameter indicates how many stack levels to go back when printing file name and line number info
func (l *Logger) ErrorNF(n int, format string, a ...interface{}) {
	l.log_internal("ERROR", fmt.Sprintf(format, a...), 2+n)
}

// Warning logs a message at Warning level
func (l *Logger) Warning(messages ...interface{}) {
	l.log_internal("WARNING", InterfacesToString(messages...), 2)
}

// WarningF logs a message at Warning level using the same syntax and options as fmt.Printf
func (l *Logger) WarningF(format string, a ...interface{}) {
	l.log_internal("WARNING", fmt.Sprintf(format, a...), 2)
}

// WarningNF is just like WarningF, but the n parameter indicates how many stack levels to go back when printing file name and line number info
func (l *Logger) WarningNF(n int, format string, a ...interface{}) {
	l.log_internal("WARNING", fmt.Sprintf(format, a...), 2+n)
}

// Notice logs a message at Notice level
func (l *Logger) Notice(messages ...interface{}) {
	l.log_internal("NOTICE", InterfacesToString(messages...), 2)
}

// NoticeF logs a message at Notice level using the same syntax and options as fmt.Printf
func (l *Logger) NoticeF(format string, a ...interface{}) {
	l.log_internal("NOTICE", fmt.Sprintf(format, a...), 2)
}

// NoticeNF is just like NoticeF, but the n parameter indicates how many stack levels to go back when printing file name and line number info
func (l *Logger) NoticeNF(n int, format string, a ...interface{}) {
	l.log_internal("NOTICE", fmt.Sprintf(format, a...), 2+n)
}

// Info logs a message at Info level
func (l *Logger) Info(messages ...interface{}) {
	l.log_internal("INFO", InterfacesToString(messages...), 2)
}

// InfoF logs a message at Info level using the same syntax and options as fmt.Printf
func (l *Logger) InfoF(format string, a ...interface{}) {
	l.log_internal("INFO", fmt.Sprintf(format, a...), 2)
}

// InfoNF is just like InfoF, but the n parameter indicates how many stack levels to go back when printing file name and line number info
func (l *Logger) InfoNF(n int, format string, a ...interface{}) {
	l.log_internal("INFO", fmt.Sprintf(format, a...), 2+n)
}

// Debug logs a message at Debug level
func (l *Logger) Debug(messages ...interface{}) {
	l.log_internal("DEBUG", InterfacesToString(messages...), 2)
}

// DebugF logs a message at Debug level using the same syntax and options as fmt.Printf
func (l *Logger) DebugF(format string, a ...interface{}) {
	l.log_internal("DEBUG", fmt.Sprintf(format, a...), 2)
}

// DebugNF is just like DebugF, but the n parameter indicates how many stack levels to go back when printing file name and line number info
func (l *Logger) DebugNF(n int, format string, a ...interface{}) {
	l.log_internal("DEBUG", fmt.Sprintf(format, a...), 2+n)
}

// Prints this goroutine's execution stack as an error with an optional message at the begining
func (l *Logger) StackAsError(messages interface{}) {
	l.log_internal("ERROR", InterfacesToString(messages, "\n", Stack()), 2)
}

// Prints this goroutine's execution stack as critical with an optional message at the begining
func (l *Logger) StackAsCritical(messages interface{}) {
	l.log_internal("CRITICAL", InterfacesToString(messages, "\n", Stack()), 2)
}

// Prints this goroutine's execution stack as debug with an optional message at the begining
func (l *Logger) StackAsDebug(messages interface{}) {
	l.log_internal("DEBUG", InterfacesToString(messages, "\n", Stack()), 2)
}

// Returns a string with the execution stack for this goroutine
func Stack() string {
	buf := make([]byte, 1000000)
	runtime.Stack(buf, false)
	return string(buf)
}

// Returns the id of the current goroutine. This function is not exported on purpose.
func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
