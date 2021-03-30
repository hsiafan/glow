package vlog

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// LogRecord is one log message
type LogRecord struct {
	LoggerName string    // logger name
	Level      Level     // the level of this logger record
	LogTime    time.Time // Time
	Message    string    // the log message
}

// Transformer convert one log record to byte array data.
// Transformer should can be share across goroutines, and user Should always reuse transformers.
type Transformer interface {
	Transform(record LogRecord) AppendEvent
}

var defaultTransformer = NewDefaultPatternTransformer()

// DefaultTransformer the default transformer used if not set
func DefaultTransformer() Transformer {
	return defaultTransformer
}

var _ Transformer = (*PatternTransformer)(nil)

// PatternTransformer transform one log record using pattern, to string
type PatternTransformer struct {
	pattern string
	items   []patternItem
}

type kind int32

const (
	text             kind = 0
	goPackage        kind = 1
	goFile           kind = 2
	goFunction       kind = 3
	lineNum          kind = 4
	loggerName       kind = 10
	loggerLevel      kind = 11
	loggerLevelUpper kind = 12
	loggerLevelLower kind = 13
	timestamp        kind = 20
	logMessage       kind = 21
)

type patternItem struct {
	kind   kind   // item type
	str    string // for text item, hold the content of text
	filter string // the filter
}

// MustNewPatternTransformer create new pattern transformer, just as NewPatternTransformer do.
// But when an error occurred, MustNewPatternTransformer panic while NewPatternTransformer return the error.
func MustNewPatternTransformer(pattern string) *PatternTransformer {
	t, err := NewPatternTransformer(pattern)
	if err != nil {
		panic(err)
	}

	return t
}

// NewPatternTransformer create new pattern transformer
// below variables can be used in format string:
// {file} filename
// {package} package name
// {line} line number
// {function} function name
// {time} time
// {logger} the logger name
// {Level}/{level}/{LEVEL} the logger level, with different character case
// {message} the log message
// use {{ to escape  {, use }} to escape }
// {time} can set custom format via filter, by {time|2006-01-02 15:04:05.000}
func NewPatternTransformer(pattern string) (*PatternTransformer, error) {
	type State int
	const (
		normalState      State = 0
		leftFirst        State = 1
		inVariableName   State = 2
		rightFirst       State = 3
		inVariableFilter State = 4
	)

	state := normalState
	var buffer []rune
	var items []patternItem
	for idx, r := range []rune(pattern) {
		switch state {
		case normalState:
			if r == '{' {
				state = leftFirst
			} else if r == '}' {
				state = rightFirst
			} else {
				buffer = append(buffer, r)
			}
		case leftFirst:
			if r == '{' {
				buffer = append(buffer, '{')
				state = normalState
			} else {
				if str := string(buffer); len(str) > 0 {
					items = append(items, patternItem{kind: text, str: str})
				}
				buffer = buffer[:0]
				buffer = append(buffer, r)
				state = inVariableName
			}
		case inVariableName:
			if r == '}' || r == '|' {
				name := string(buffer)
				buffer = buffer[:0]
				if name == "file" {
					items = append(items, patternItem{kind: goFile})
				} else if name == "package" {
					items = append(items, patternItem{kind: goPackage})
				} else if name == "function" {
					items = append(items, patternItem{kind: goFunction})
				} else if name == "line" {
					items = append(items, patternItem{kind: lineNum})
				} else if name == "time" {
					items = append(items, patternItem{kind: timestamp})
				} else if name == "logger" {
					items = append(items, patternItem{kind: loggerName})
				} else if name == "message" {
					items = append(items, patternItem{kind: logMessage})
				} else if name == "Level" {
					items = append(items, patternItem{kind: loggerLevel})
				} else if name == "level" {
					items = append(items, patternItem{kind: loggerLevelLower})
				} else if name == "LEVEL" {
					items = append(items, patternItem{kind: loggerLevelUpper})
				} else {
					return nil, errors.New("unknown variable name: " + name)
				}
				if r == '}' {
					state = normalState
				} else if r == '|' {
					state = inVariableFilter
				}
			} else {
				buffer = append(buffer, r)
			}
		case rightFirst:
			if r == '}' {
				buffer = append(buffer, '}')
				state = normalState
			} else {
				return nil, errors.New("unexpected } at " + strconv.Itoa(idx-1))
			}
		case inVariableFilter:
			if r == '}' {
				if str := string(buffer); len(str) > 0 {
					(&items[len(items)-1]).filter = str
				}
				buffer = buffer[:0]
				state = normalState
			} else {
				buffer = append(buffer, r)
			}
		default:
			return nil, errors.New("unhandled state: " + strconv.Itoa(int(state)))
		}
	}
	if state != normalState {
		return nil, errors.New("format str does not finish rightly")
	}
	// the last part
	if str := string(buffer); len(str) > 0 {
		items = append(items, patternItem{kind: text, str: str})
	}

	return &PatternTransformer{pattern: pattern, items: items}, nil
}

// NewDefaultPatternTransformer return formatter with default format
func NewDefaultPatternTransformer() *PatternTransformer {
	formatter, err := NewPatternTransformer("{time} [{Level}] {logger} - {message}\n")
	if err != nil {
		panic(err)
	}
	return formatter
}

// Transform format log data to byte array data
func (f *PatternTransformer) Transform(record LogRecord) AppendEvent {

	var logItems []string
	var caller *caller
	depth := 5
	for _, item := range f.items {
		switch item.kind {
		case text:
			logItems = append(logItems, item.str)
		case timestamp:
			var timeFormat string
			if item.filter == "" {
				timeFormat = "2006-01-02 15:04:05.000"
			} else {
				timeFormat = item.filter
			}
			logItems = append(logItems, record.LogTime.Format(timeFormat))
		case loggerName:
			logItems = append(logItems, record.LoggerName)
		case loggerLevel:
			logItems = append(logItems, record.Level.Name())
		case loggerLevelUpper:
			logItems = append(logItems, strings.ToUpper(record.Level.Name()))
		case loggerLevelLower:
			logItems = append(logItems, strings.ToLower(record.Level.Name()))
		case logMessage:
			logItems = append(logItems, record.Message)
		case goPackage:
			if caller == nil {
				caller = getCaller(depth)
			}
			logItems = append(logItems, caller.packageName)
		case goFile:
			if caller == nil {
				caller = getCaller(depth)
			}
			logItems = append(logItems, caller.fileName)
		case goFunction:
			if caller == nil {
				caller = getCaller(depth)
			}
			logItems = append(logItems, caller.functionName)
		case lineNum:
			if caller == nil {
				caller = getCaller(depth)
			}
			logItems = append(logItems, strconv.Itoa(caller.line))
		default:
			panic("unsupported type: " + strconv.Itoa(int(item.kind)))
		}
	}

	var message = strings.Join(logItems, "")
	return AppendEvent{Message: message}
}
