package osx

import (
	"context"
	"os/exec"
	"time"
)

// RunCommandTimeout run the command, with a timeout value.
// If command does not finish before timeout, the process would be destroyed, a timeout error will be return.
func RunCommandTimeout(cmd *exec.Cmd, timeout time.Duration) error {
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
		return context.DeadlineExceeded
	case err := <-done:
		timer.Stop()
		return err
	}
}

// RunCommandContext run the command, with a context, which may cancel before command finished..
// If so the process would be destroyed, a error will be return.
func RunCommandContext(cmd *exec.Cmd, ctx context.Context) error {
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
		return ctx.Err()
	case err := <-done:
		return err
	}
}
