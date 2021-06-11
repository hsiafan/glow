package osx

import (
	"context"
	"errors"
	"os/exec"
	"time"
)

// timeoutError is a timeout error
type timeoutError struct {
	message string
}

func (t *timeoutError) Error() string {
	return t.message
}

func (t *timeoutError) Timeout() bool {
	return true
}

// RunCommandWithTimeout run the command, with a timeout value.
// If command does not finish before timeout, the process would be destroyed, a timeout error will be return.
func RunCommandWithTimeout(cmd *exec.Cmd, timeout time.Duration) error {
	if err := cmd.Start(); err != nil {
		return err
	}

	timer := time.NewTimer(timeout)

	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-timer.C:
		_ = cmd.Process.Kill()
		return &timeoutError{"command run timeout"}
	case err := <-done:
		timer.Stop()
		return err
	}
}

// RunCommandWithContext run the command, with a context, which may cancel before command finished..
// If so the process would be destroyed, a error will be return.
func RunCommandWithContext(cmd *exec.Cmd, ctx context.Context) error {
	if err := cmd.Start(); err != nil {
		return err
	}

	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-ctx.Done():
		_ = cmd.Process.Kill()
		return errors.New("command canceled")
	case err := <-done:
		return err
	}
}
