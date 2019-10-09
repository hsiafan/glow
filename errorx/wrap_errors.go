package errorx

// WrappedError wrapped another error
type WrappedError interface {
	error
	Cause() error // The error wrapped
}

// Wrap wrap an error with a message, to a WrappedError
func Wrap(err error, message string) WrappedError {
	return &wrappedError{err, message}
}

// RootCause return the root cause of this error.
// That means it will unwrap error till is not a WrappedError
func RootCause(err error) error {
	for {
		we, ok := err.(WrappedError)
		if !ok {
			return err
		}
		err = we.Cause()
	}
}

type wrappedError struct {
	error
	message string
}

// Cause implement WrappedError
func (w *wrappedError) Cause() error {
	return w.error
}

// Error implement error
func (w *wrappedError) Error() string {
	return w.message + "; Caused by: " + w.error.Error()
}
