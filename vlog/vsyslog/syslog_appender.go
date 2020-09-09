package vsyslog

import (
	"errors"
	"github.com/hsiafan/glow/vlog"
	"log/syslog"
	"strconv"
)

var _ vlog.Appender = (*Appender)(nil)

// Appender write log to syslogd, using go syslog package.
// This appender always send only raw log message, Appender will not take effect.
//
// Appender will map log levels from vlog to syslog by the following rules:
// TRACE		-- LOG_DEBUG
// DEBUG		-- LOG_DEBUG
// INFO			-- LOG_INFO
// WARN			-- LOG_WARNING
// ERROR		-- LOG_ERR
// CRITICAL		-- LOG_CRIT
type Appender struct {
	log         *syslog.Writer
	levelMap    map[vlog.Level]syslog.Priority
	transformer sysLogTransformer
}

// defaultLevelMap is the default level map from vlog to syslog
var defaultLevelMap = map[vlog.Level]syslog.Priority{
	vlog.Trace:    syslog.LOG_DEBUG,
	vlog.Debug:    syslog.LOG_DEBUG,
	vlog.Info:     syslog.LOG_INFO,
	vlog.Warn:     syslog.LOG_WARNING,
	vlog.Error:    syslog.LOG_ERR,
	vlog.Critical: syslog.LOG_CRIT,
}

type sysLogTransformer struct {
}

func (st sysLogTransformer) Transform(record vlog.LogRecord) vlog.AppendEvent {
	return vlog.AppendEvent{Message: record.Message}
}

// New create syslog appender, to system syslog daemon.
func New(tag string) (*Appender, error) {
	log, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL0, tag)
	if err != nil {
		return nil, err
	}
	return &Appender{log: log, levelMap: defaultLevelMap}, nil
}

// NewToAddress create syslog appender, to a log daemon connected by network address.
func NewToAddress(network string, address string, tag string) (vlog.Appender, error) {
	log, err := syslog.Dial(network, address, syslog.LOG_INFO|syslog.LOG_LOCAL0, tag)
	if err != nil {
		return nil, err
	}
	return &Appender{log: log, levelMap: defaultLevelMap}, nil
}

// SetLevelMap set level map from vlog to syslog, replace the default log level map.
// This method should be called before appender start to work.
func (sa *Appender) SetLevelMap(levelMap map[vlog.Level]syslog.Priority) {
	sa.levelMap = levelMap
}

// Append write one log entry to syslog
func (sa *Appender) Append(event vlog.AppendEvent) error {
	var level = event.Level
	if priority, ok := sa.levelMap[level]; ok {
		switch priority {
		case syslog.LOG_DEBUG:
			return sa.log.Debug(event.Message)
		case syslog.LOG_INFO:
			return sa.log.Info(event.Message)
		case syslog.LOG_NOTICE:
			return sa.log.Notice(event.Message)
		case syslog.LOG_WARNING:
			return sa.log.Warning(event.Message)
		case syslog.LOG_ERR:
			return sa.log.Err(event.Message)
		case syslog.LOG_CRIT:
			return sa.log.Crit(event.Message)
		case syslog.LOG_ALERT:
			return sa.log.Alert(event.Message)
		case syslog.LOG_EMERG:
			return sa.log.Emerg(event.Message)
		default:
			return errors.New("unknown syslog level: " + strconv.Itoa(int(priority)))
		}
	}

	_, err := sa.log.Write([]byte(event.Message))
	return err
}

// Transformer always return the default, non-
func (sa *Appender) Transformer() vlog.Transformer {
	return sa.transformer
}

// SetTransformer not take effect for Appender, which always only send log message
func (sa *Appender) SetTransformer(transformer vlog.Transformer) {

}

// Close the syslog connection
func (sa *Appender) Close() error {
	return sa.log.Close()
}
