package vlog

import (
	"bytes"
	"os"
)

// Appender write the log to one destination, and can provider a transformer to convert the log message to desired data.
// Appender Should be reused across loggers.
type Appender interface {
	// Append append new data to destination. name is the name of logger, level is the level of logger
	Append(event AppendEvent) error
}

// AppendEvent is a log event passed to Appender
type AppendEvent struct {
	LoggerName string
	Level      Level
	Message    string // the logger message
}

// ConsoleAppender appender write log to stdout
type ConsoleAppender struct {
	file *os.File
}

// Append log to stdout
func (ca *ConsoleAppender) Append(event AppendEvent) error {
	_, err := ca.file.WriteString(event.Message)
	return err
}

var defaultAppender Appender = NewConsoleAppender()

// DefaultAppender return the default appender all logger use
func DefaultAppender() Appender {
	return defaultAppender
}

// NewConsoleAppender create console appender, which write log to stdout
func NewConsoleAppender() *ConsoleAppender {
	return &ConsoleAppender{file: os.Stdout}
}

// NewConsole2Appender create console appender, which write log to stderr
func NewConsole2Appender() *ConsoleAppender {
	return &ConsoleAppender{file: os.Stderr}
}

var _ Appender = (*NopAppender)(nil)

// NopAppender discard all logs
type NopAppender struct {
}

// NewNopAppender create nop appender
func NewNopAppender() *NopAppender {
	return &NopAppender{}
}

// Append silently discard log data
func (NopAppender) Append(event AppendEvent) error {
	return nil
}

var _ Appender = (*BytesAppender)(nil)

// BytesAppender write log into memory
type BytesAppender struct {
	buffer bytes.Buffer
}

// NewBytesAppender create BytesAppender
func NewBytesAppender() *BytesAppender {
	return &BytesAppender{}
}

// Append write log data to byte buffer
func (b *BytesAppender) Append(event AppendEvent) error {
	_, err := b.buffer.WriteString(event.Message)
	return err
}
