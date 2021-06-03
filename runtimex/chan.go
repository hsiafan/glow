package runtimex

import (
	"sync/atomic"
	"unsafe"
)

type myChan struct {
	queueCount    uint           // total data in the queue
	dataQueueSize uint           // size of the circular queue
	buf           unsafe.Pointer // points to an array of dataqsiz elements
	elemSize      uint16
	closed        uint32
}

// Chan the channel type, to be generified
type Chan = interface{}

// ChanIsClosed return if channel is closed, no read from channel required.
func ChanIsClosed(c Chan) bool {
	hc := (*myChan)(InterfaceValuePtr(c))
	return atomic.LoadUint32(&hc.closed) != 0
}
