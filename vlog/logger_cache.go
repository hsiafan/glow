package vlog

import (
	"os"
	"strings"
	"sync"
	"unsafe"
)

var loggerCache = initLogCache()

var levelNamesMap = reverseLevelNames(levelNames)

func reverseLevelNames(levelNames map[Level]string) map[string]Level {
	var m = map[string]Level{}
	for level, str := range levelNames {
		m[strings.ToUpper(str)] = level
	}
	return m
}

// create unique global log cache
func initLogCache() *LoggerCache {
	cache := newLogCache()
	return cache
}

func newLogCache() *LoggerCache {
	var loggerConfigs []*loggerConfig
	var levelStr = os.Getenv("VLOG_LEVEL")
	if len(levelStr) > 0 {
		for _, levelPair := range strings.Split(levelStr, ";") {
			levelPair = strings.TrimSpace(levelPair)
			if len(levelPair) == 0 {
				continue
			}
			var idx = strings.IndexByte(levelPair, byte('='))
			if idx == -1 {
				continue
			}
			var prefix = strings.TrimSpace(levelPair[:idx])
			if len(prefix) == 0 {
				continue
			}
			var level, ok = levelNamesMap[strings.ToUpper(strings.TrimSpace(levelPair[idx+1:]))]
			if !ok {
				continue
			}
			loggerConfigs = append(loggerConfigs, &loggerConfig{prefix: prefix, level: level})
		}
	}

	return &LoggerCache{
		loggerMap:  make(map[string]*Logger),
		logConfigs: loggerConfigs,
	}
}

// LoggerCache contains loggers with name as key
type LoggerCache struct {
	logConfigs []*loggerConfig
	loggerMap  map[string]*Logger
	lock       sync.Mutex
}

// Load return logger for with name, using cached one or create new one if logger with name not exist
func (lc *LoggerCache) Load(name string) *Logger {
	lc.lock.Lock()
	defer lc.lock.Unlock()
	logger, ok := lc.loggerMap[name]
	if ok {
		return logger
	}

	var level = DefaultLevel
	var frozen = false
	logConfig := lc.matchConfig(name)
	if logConfig != nil {
		level = logConfig.level
		frozen = true
	}

	appenders := &[]Appender{defaultAppender}
	logger = &Logger{
		name:      name,
		level:     int32(level),
		appenders: unsafe.Pointer(appenders),
		frozen:    frozen,
	}
	lc.loggerMap[name] = logger
	return logger
}

func (lc *LoggerCache) matchConfig(name string) *loggerConfig {
	var maxMatchLen = 0
	var config *loggerConfig
	for _, logConfig := range lc.logConfigs {
		if matchPrefix(name, logConfig.prefix) {
			matchLen := len(logConfig.prefix)
			if matchLen >= maxMatchLen {
				config = logConfig
				maxMatchLen = matchLen
			}
		}
	}
	return config
}

func (lc *LoggerCache) filter(prefix string) []*Logger {
	var loggers []*Logger
	for _, logger := range lc.loggerMap {
		if matchPrefix(logger.Name(), prefix) {
			loggers = append(loggers, logger)
		}
	}
	return loggers
}

func matchPrefix(name string, prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	if len(prefix) > len(name) {
		return false
	}
	i := 0
	for ; i < len(prefix); i++ {
		if prefix[i] != name[i] {
			return false
		}
	}
	if prefix[i-1] == '/' || len(prefix) == len(name) || name[i] == '/' {
		return true
	}
	return false
}

// loggerConfig used to config logger levels
type loggerConfig struct {
	prefix string
	level  Level
}
