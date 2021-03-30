package vlog

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	logger := CurrentPackageLogger()
	assert.Equal(t, "github.com/hsiafan/glow/vlog", logger.Name())
	assert.Equal(t, DefaultLevel, logger.Level())

	appender := NewBytesAppender()
	logger.SetAppenders(appender)
	logger.AtInfo().Log("this is a test")
	assert.True(t, strings.HasSuffix(appender.buffer.String(),
		" [Info] github.com/hsiafan/glow/vlog - this is a test\n"))

	appender = NewBytesAppender()
	logger.SetAppenders(appender)
	logger.AtInfo().LogFormat("this is a test")
	assert.True(t, strings.HasSuffix(appender.buffer.String(),
		" [Info] github.com/hsiafan/glow/vlog - this is a test\n"))

	appender = NewBytesAppender()
	logger.SetAppenders(appender)
	logger.AtInfo().LogLazy(func() string { return "this is a test" })
	assert.True(t, strings.HasSuffix(appender.buffer.String(),
		" [Info] github.com/hsiafan/glow/vlog - this is a test\n"))

	appender = NewBytesAppender()
	transformer, _ := NewPatternTransformer("{time|2006-01-02} {package}/{file} - {message}\n")
	logger.SetTransformer(transformer)
	logger.SetAppenders(appender)
	logger.AtInfo().Log("this is a test")
	date := time.Now().Format("2006-01-02")
	assert.Equal(t, date+" github.com/hsiafan/glow/vlog/logger_test.go - this is a test\n", appender.buffer.String())

	logger2 := CurrentPackageLogger()
	assert.Equal(t, logger, logger2)
}

func TestLoggerJudge(t *testing.T) {
	logger := CurrentPackageLogger()
	logger.SetLevel(Off)
	assert.False(t, logger.TraceEnabled())
	assert.False(t, logger.ErrorEnabled())

	logger.SetLevel(Critical)
	assert.True(t, logger.CriticalEnabled())
	assert.False(t, logger.ErrorEnabled())

	logger.SetLevel(Error)
	assert.True(t, logger.ErrorEnabled())
	assert.False(t, logger.InfoEnabled())

	logger.SetLevel(Trace)
	assert.True(t, logger.TraceEnabled())
	assert.True(t, logger.InfoEnabled())
}

func TestLogger_AddAppender(t *testing.T) {
	logger := CurrentPackageLogger()
	assert.Equal(t, 1, len(logger.Appenders()))
	appender := NewConsole2Appender()
	logger.AddAppenders(appender)
	assert.Equal(t, 2, len(logger.Appenders()))
	assert.Equal(t, appender, logger.Appenders()[1])
}

func TestFormatMessage(t *testing.T) {
	assert.Equal(t, "This is a test", formatMessage("This is a test"), "")
	assert.Equal(t, "This is 1", formatMessage("This is {}", 1), "")
	assert.Equal(t, "This is 1", formatMessage("This is {}", 1, 2), "")
	assert.Equal(t, "This is 1, {}", formatMessage("This is {}, {}", 1), "")
	assert.Equal(t, "1, 2", formatMessage("{}, {}", 1, 2), "")
}

func TestJoinMessage(t *testing.T) {
	assert.Equal(t, "test", joinMessage("test"))
	assert.Equal(t, "test 2", joinMessage("test", "2"))
}
