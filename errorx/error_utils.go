package errorx

import (
	"errors"
	"strings"
)

// WrapError is an interface for wrapped error.
// The error generated by fmt.Errorf with an origin error also implement this interface
type WrapError interface {
	error
	// Unwrap return the wrapped error
	Unwrap() error
}

var _ WrapError = (*wrapError)(nil)

// a wrap error impl
type wrapError struct {
	message string
	err     error
}

func (e *wrapError) Error() string {
	var sb strings.Builder
	e.writeTo(&sb)
	return sb.String()
}

func (e *wrapError) Unwrap() error {
	return e.err
}

func (e *wrapError) writeTo(sb *strings.Builder) {
	sb.WriteString(e.message)
	sb.WriteString("\n    caused by: ")
	we, ok := e.Unwrap().(*wrapError)
	if ok {
		we.writeTo(sb)
	} else {
		sb.WriteString(e.Unwrap().Error())
	}
}

// Wrap return a new error wrapped the origin error, with an additional message.
func Wrap(message string, err error) error {
	if err == nil {
		panic("error is nil")
	}
	return &wrapError{
		message: message,
		err:     err,
	}
}

// UnwrapUtil unwrap the err chain, util the err meet the predicate func, and then return the err.
// If not found meet err, return nil
func UnwrapUtil(err error, f func(e error) bool) error {
	for {
		if f(err) {
			return err
		}
		ne := errors.Unwrap(err)
		if ne == nil {
			return nil
		}
		err = ne
	}
}

// UnwrapToRoot recursively unwrap the err chain, util the err is not a WrapError, and return this err.
func UnwrapToRoot(err error) error {
	for {
		ne := errors.Unwrap(err)
		if ne == nil {
			return err
		}
		err = ne
	}
}

// PanicOnError panic when err is not nil
func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
