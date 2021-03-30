package vlog

import (
	"bytes"
	"os"
	"sync/atomic"
)

// Appender write the log to one destination, and can provider a transformer to convert the log message to desired data.
// Appender Should be reused across loggers.
type Appender interface {
	// Append append new data to destination. name is the name of logger, level is the level of logger
	Append(event AppendEvent) error
	// get the transformer of this appender
	Transformer() Transformer
	// set transformer to this appender
	SetTransformer(transformer Transformer)
}

// AppendEvent is a log event passed to Appender
type AppendEvent struct {
	LoggerName string
	Level      Level
	Message    string // the logger message
}

// CanFormattedMixin used for impl Appender Transformer/Name... methods
type CanFormattedMixin struct {
	transformer atomic.Value //*Transformer
}

// NewAppenderMixin create new CanFormattedMixin
func NewAppenderMixin() *CanFormattedMixin {
	return &CanFormattedMixin{}
}

// Transformer return the transformer of this appender. This method is thread-safe
func (am *CanFormattedMixin) Transformer() Transformer {
	iface := am.transformer.Load()
	if iface == nil {
		return DefaultTransformer()
	}
	return *iface.(*Transformer)
}

// SetTransformer set transformer to this appender. This method is thread-safe
func (am *CanFormattedMixin) SetTransformer(transformer Transformer) {
	am.transformer.Store(&transformer)
}

// ConsoleAppender appender write log to stdout
type ConsoleAppender struct {
	*CanFormattedMixin
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
	return &ConsoleAppender{file: os.Stdout, CanFormattedMixin: NewAppenderMixin()}
}

// NewConsole2Appender create console appender, which write log to stderr
func NewConsole2Appender() *ConsoleAppender {
	return &ConsoleAppender{file: os.Stderr, CanFormattedMixin: NewAppenderMixin()}
}

var _ Appender = (*NopAppender)(nil)

// NopAppender discard all logs
type NopAppender struct {
	*CanFormattedMixin
}

// NewNopAppender create nop appender
func NewNopAppender() *NopAppender {
	return &NopAppender{CanFormattedMixin: NewAppenderMixin()}
}

// Append silently discard log data
func (NopAppender) Append(event AppendEvent) error {
	return nil
}

var _ Appender = (*BytesAppender)(nil)

// BytesAppender write log into memory
type BytesAppender struct {
	*CanFormattedMixin
	buffer bytes.Buffer
}

// NewBytesAppender create BytesAppender
func NewBytesAppender() *BytesAppender {
	return &BytesAppender{CanFormattedMixin: NewAppenderMixin()}
}

// Append write log data to byte buffer
func (b *BytesAppender) Append(event AppendEvent) error {
	_, err := b.buffer.WriteString(event.Message)
	return err
}
