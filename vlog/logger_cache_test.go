package vlog

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	logCache := newLogCache()
	logger1 := logCache.Load("logger1")
	logger2 := logCache.Load("logger2")
	logger3 := logCache.Load("logger1")

	assert.Equal(t, logger1, logger3)
	assert.NotEqual(t, logger1, logger2)
}

func TestLoggerCache_Filter(t *testing.T) {
	logCache := newLogCache()
	_ = logCache.Load("package0")
	_ = logCache.Load("package1")
	_ = logCache.Load("gopkg.in/package1")
	_ = logCache.Load("github.com/user1/package1")
	_ = logCache.Load("github.com/user2/package2")
	_ = logCache.Load("github.com/user3/package3")

	loggers := logCache.filter("")
	assert.Equal(t, 6, len(loggers))

	loggers = logCache.filter("github.com")
	assert.Equal(t, 3, len(loggers))

	loggers = logCache.filter("github.com/user1/package1")
	assert.Equal(t, 1, len(loggers))

	loggers = logCache.filter("gopkg.in/")
	assert.Equal(t, 1, len(loggers))
	assert.Equal(t, "gopkg.in/package1", loggers[0].Name())
}

func TestCache_SetPrefix(t *testing.T) {
	defer os.RemoveAll("logs/")
	defer os.Unsetenv("VLOG_LEVEL")
	os.Setenv("VLOG_LEVEL", "package1=Warn;github.com/user1=debug")
	logCache := initLogCache()
	logger1 := logCache.Load("package1")
	logger2 := logCache.Load("gopkg.in/package1")
	logger3 := logCache.Load("github.com/user1/package1")

	assert.Equal(t, Warn.Name(), logger1.Level().Name())
	assert.Equal(t, Info.Name(), logger2.Level().Name())
	assert.Equal(t, Debug.Name(), logger3.Level().Name())
}
