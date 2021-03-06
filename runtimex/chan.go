package runtimex

import (
	"github.com/hsiafan/glow/internal"
	"sync/atomic"
)

// Chan the channel type, to be generified
type Chan = interface{}

// ChanClosed return if channel is closed, no read from channel required.
func ChanClosed(c Chan) bool {
	hc := (*internal.Chan)(InterfaceValuePtr(c))
	return atomic.LoadUint32(&hc.Closed) != 0
}
