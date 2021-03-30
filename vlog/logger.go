package vlog

import (
	"fmt"
	"os"
	"strings"
	"sync/atomic"
	"time"
	"unsafe"
)

// Level the logger level
type Level int32

var levelNames = map[Level]string{
	Trace:    "Trace",
	Debug:    "Debug",
	Info:     "Info",
	Warn:     "Warn",
	Error:    "Error",
	Critical: "Critical",
}

// Name return the name of level, using captical form
func (l Level) Name() string {
	return levelNames[l]
}

// log levels
const (
	Trace        Level = 10
	Debug        Level = 20
	Info         Level = 30
	Warn         Level = 40
	Error        Level = 50
	Critical     Level = 60
	Off          Level = 70
	DefaultLevel Level = Info
)

// Logger the logger
type Logger struct {
	name      string
	level     int32          //Level
	appenders unsafe.Pointer //*[]Appender
	frozen    bool           // frozen level. the level is set by env, following level set in code will not take effect
}

// Name the name of this logger
func (l *Logger) Name() string {
	return l.name
}

// SetLevel set new Level to this logger. the default log level is Debug
func (l *Logger) SetLevel(level Level) {
	if l.frozen {
		return
	}
	atomic.StoreInt32(&l.level, int32(level))
}

// Level current level of this logger
func (l *Logger) Level() Level {
	return Level(atomic.LoadInt32(&l.level))
}

// SetAppenders set one or multi appenders for this logger
func (l *Logger) SetAppenders(appenders ...Appender) {
	atomic.StorePointer(&l.appenders, unsafe.Pointer(&appenders))
}

// Appenders return the appenders this logger have
func (l *Logger) Appenders() []Appender {
	return *(*[]Appender)(atomic.LoadPointer(&l.appenders))
}

// AddAppenders add one new appender to logger
func (l *Logger) AddAppenders(appenders ...Appender) {
	if len(appenders) == 0 {
		return
	}

	for {
		p := atomic.LoadPointer(&l.appenders)
		originAppenders := *(*[]Appender)(p)
		newAppenders := make([]Appender, len(originAppenders)+len(appenders))
		copy(newAppenders, originAppenders)
		copy(newAppenders[len(originAppenders):], appenders)
		if atomic.CompareAndSwapPointer(&l.appenders, p, unsafe.Pointer(&newAppenders)) {
			break
		}
	}
}

// SetTransformerForAppenders set transformer, apply to all appenders the logger current have
func (l *Logger) SetTransformerForAppenders(transformer Transformer) {
	for _, appender := range l.Appenders() {
		appender.SetTransformer(transformer)
	}
}

// Trace log message with trace level
func (l *Logger) Trace(firstArg interface{}, args ...interface{}) {
	l.log(Trace, firstArg, args...)
}

// Debug log message with debug level
func (l *Logger) Debug(firstArg interface{}, args ...interface{}) {
	l.log(Debug, firstArg, args...)
}

// Info log message with info level
func (l *Logger) Info(firstArg interface{}, args ...interface{}) {
	l.log(Info, firstArg, args...)
}

// Warn log message with warn level
func (l *Logger) Warn(firstArg interface{}, args ...interface{}) {
	l.log(Warn, firstArg, args...)
}

// log message with error level
func (l *Logger) Error(firstArg interface{}, args ...interface{}) {
	l.log(Error, firstArg, args...)
}

// Critical log message with critical level
func (l *Logger) Critical(firstArg interface{}, args ...interface{}) {
	l.log(Critical, firstArg, args...)
}

// TraceFormat log message with trace level
func (l *Logger) TraceFormat(format string, args ...interface{}) {
	l.logFormat(Trace, format, args...)
}

// DebugFormat log message with debug level
func (l *Logger) DebugFormat(format string, args ...interface{}) {
	l.logFormat(Debug, format, args...)
}

// InfoFormat log message with info level
func (l *Logger) InfoFormat(format string, args ...interface{}) {
	l.logFormat(Info, format, args...)
}

// WarnFormat log message with warn level
func (l *Logger) WarnFormat(format string, args ...interface{}) {
	l.logFormat(Warn, format, args...)
}

// ErrorFormat message with error level
func (l *Logger) ErrorFormat(format string, args ...interface{}) {
	l.logFormat(Error, format, args...)
}

// CriticalFormat log message with critical level
func (l *Logger) CriticalFormat(format string, args ...interface{}) {
	l.logFormat(Critical, format, args...)
}

// TraceEnabled if this logger log trace message
func (l *Logger) TraceEnabled() bool {
	return l.Level() <= Trace
}

// DebugEnabled if this logger log debug message
func (l *Logger) DebugEnabled() bool {
	return l.Level() <= Debug
}

// InfoEnabled if this logger log info message
func (l *Logger) InfoEnabled() bool {
	return l.Level() <= Info
}

// WarnEnabled if this logger log warn level message
func (l *Logger) WarnEnabled() bool {
	return l.Level() <= Warn
}

// ErrorEnabled if this logger log error message
func (l *Logger) ErrorEnabled() bool {
	return l.Level() <= Error
}

// CriticalEnabled if this logger log critical message
func (l *Logger) CriticalEnabled() bool {
	return l.Level() <= Critical
}

// TraceLazy log message with trace level, and call func to get log message only when log is performed.
func (l *Logger) TraceLazy(f func() string) {
	if !l.TraceEnabled() {
		return
	}
	l.logString(Trace, f())
}

// DebugLazy log message with debug level, and call func to get log message only when log is performed.
func (l *Logger) DebugLazy(f func() string) {
	if !l.DebugEnabled() {
		return
	}
	l.logString(Debug, f())
}

// InfoLazy log message with info level, and call func to get log message only when log is performed.
func (l *Logger) InfoLazy(f func() string) {
	if !l.InfoEnabled() {
		return
	}
	l.logString(Info, f())
}

// WarnLazy log message with warn level, and call func to get log message only when log is performed.
func (l *Logger) WarnLazy(f func() string) {
	if !l.WarnEnabled() {
		return
	}
	l.logString(Warn, f())
}

// ErrorLazy log message with error level, and call func to get log message only when log is performed.
func (l *Logger) ErrorLazy(f func() string) {
	if !l.ErrorEnabled() {
		return
	}
	l.logString(Error, f())
}

// CriticalLazy log message with critical level, and call func to get log message only when log is performed.
func (l *Logger) CriticalLazy(f func() string) {
	if !l.CriticalEnabled() {
		return
	}
	l.logString(Critical, f())
}

// log multi messages, delimited with a white space
func (l *Logger) log(level Level, firstArg interface{}, args ...interface{}) {
	appenders := l.Appenders()
	if l.Level() <= level && len(appenders) > 0 {
		message := joinMessage(firstArg, args...)
		if err := l.writeToAppends(level, appenders, message); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "log error", err)
		}
	}
}

// log one string message
func (l *Logger) logString(level Level, message string) {
	appenders := l.Appenders()
	if l.Level() <= level && len(appenders) > 0 {
		if err := l.writeToAppends(level, appenders, message); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "log error", err)
		}
	}
}

// log formated messages as java slf4j style.
func (l *Logger) logFormat(level Level, format string, args ...interface{}) {
	appenders := l.Appenders()
	if l.Level() <= level && len(appenders) > 0 {
		message := formatMessage(format, args...)
		if err := l.writeToAppends(level, appenders, message); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "log error", err)
		}
	}
}

func (l *Logger) writeToAppends(level Level, appenders []Appender, message string) error {
	now := time.Now()
	//TODO: async, parallel write
	for _, appender := range appenders {
		transformer := appender.Transformer()
		appendEvent := transformer.Transform(LogRecord{l.Name(), level, now, message})
		err := appender.Append(appendEvent)
		if err != nil {
			//TODO: collection errors
			return err
		}
	}
	return nil
}

func joinMessage(message interface{}, args ...interface{}) string {
	var results = make([]string, len(args)+1)
	results[0] = argToString(message)
	for idx := 0; idx < len(args); idx++ {
		results[idx+1] = argToString(args[idx])
	}

	return strings.Join(results, " ")
}

func formatMessage(format string, args ...interface{}) string {
	argNum := len(args)
	items := strings.SplitN(format, "{}", argNum+1)

	var results []string
	var minArgNum = len(items) - 1
	if minArgNum > argNum {
		minArgNum = argNum
	}

	for idx, item := range items {
		results = append(results, item)
		if idx < minArgNum {
			results = append(results, argToString(args[idx]))
		}
	}
	return strings.Join(results, "")
}

func argToString(arg interface{}) string {
	return fmt.Sprint(arg)
}

// GetLogger return the logger with name
func GetLogger(name string) *Logger {
	return loggerCache.Load(name)
}

// CurrentPackageLogger return the log of current package, use package name as logger name
func CurrentPackageLogger() *Logger {
	caller := getCaller(2)
	return GetLogger(caller.packageName)
}
