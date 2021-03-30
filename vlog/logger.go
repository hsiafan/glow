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
func (l *Logger) AtTrace() LoggerContext {
	return LoggerContext{logger: l, level: Trace}
}

// Debug log message with debug level
func (l *Logger) AtDebug() LoggerContext {
	return LoggerContext{logger: l, level: Debug}
}

// Info log message with info level
func (l *Logger) AtInfo() LoggerContext {
	return LoggerContext{logger: l, level: Info}
}

// Warn log message with warn level
func (l *Logger) AtWarn() LoggerContext {
	return LoggerContext{logger: l, level: Warn}
}

// log message with error level
func (l *Logger) AtError() LoggerContext {
	return LoggerContext{logger: l, level: Error}
}

// Critical log message with critical level
func (l *Logger) AtCritical() LoggerContext {
	return LoggerContext{logger: l, level: Critical}
}

// LoggerContext for logger call
type LoggerContext struct {
	logger *Logger
	level  Level
}

// Log log message
func (l LoggerContext) Log(firstArg interface{}, args ...interface{}) {
	l.logger.log(l.level, firstArg, args...)
}

// Log log message with format and args
func (l LoggerContext) LogFormat(format string, args ...interface{}) {
	l.logger.logFormat(l.level, format, args...)
}

// LogLazy log message, it call func to get log message only when log is performed.
func (l LoggerContext) LogLazy(f func() string) {
	if l.logger.Level() <= l.level {
		l.logger.logString(l.level, f())
	}
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
