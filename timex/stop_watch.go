package timex

import "time"

const (
	statusInit    = 0
	statusStarted = 1
	statusStopped = 2
)

// Stopwatch for measure elapsed time. This class is NOT thread-safe.
type Stopwatch struct {
	startedTime time.Time
	stoppedTime time.Time
	status      int
}

// NewStopwatch create new Stopwatch, not started.
func NewStopwatch() *Stopwatch {
	return &Stopwatch{status: statusInit}
}

// NewStopwatchStarted create new Stopwatch, then start it.
func NewStopwatchStarted() *Stopwatch {
	return &Stopwatch{
		startedTime: time.Now(),
		status:      statusStarted,
	}
}

// Start the Stopwatch. If StopWatch is already started, return false.
// A stopped Stopwatch can start again.
func (w *Stopwatch) Start() bool {
	switch w.status {
	case statusStarted:
		return false
	case statusInit:
	case statusStopped:
		w.startedTime = time.Now()
		w.status = statusStopped
		return true
	default:
		panic("unknown status")
	}
	panic("")
}

// Stop the Stopwatch. The elapsed time will not changed after stopped.
// Return true if stop succeed. If Stopwatch is not started, or already stopped, return false.
func (w *Stopwatch) Stop() bool {
	switch w.status {
	case statusStarted:
		w.status = statusStopped
		w.stoppedTime = time.Now()
		return true
	case statusStopped:
	case statusInit:
		return false
	default:
		panic("unknown status")
	}
	panic("")
}

// Elapsed return elapsed time.
// If not started, always return zero;
// If started and not stopped, return time interval from begin time to now;
// If stopped, return elapsed time duration from started time to stopped time.
func (w *Stopwatch) Elapsed() time.Duration {
	switch w.status {
	case statusStarted:
		return time.Now().Sub(w.startedTime)
	case statusStopped:
		return w.stoppedTime.Sub(w.startedTime)
	case statusInit:
		return time.Duration(0)
	}
	panic("")
}

// Elapsed return elapsed time by milli-seconds.
// If not started, always return zero;
// If started and not stopped, return time interval from begin time to now;
// If stopped, return elapsed time duration from started time to stopped time.
func (w *Stopwatch) ElapsedMillis() int64 {
	return int64(w.Elapsed() / time.Millisecond)
}
