package concurrencyutils

import (
	"sync"
)

type CountDownLatch struct {
	m       sync.Mutex
	counter int
	cond    *sync.Cond
}

func NewCountDownLatch(counter int) *CountDownLatch {
	cdl := &CountDownLatch{
		counter: counter,
	}
	cdl.cond = sync.NewCond(&cdl.m)
	return cdl
}

func (cdl *CountDownLatch) CountDown() {
	cdl.m.Lock()
	defer cdl.m.Unlock()

	// Decrement the counter
	cdl.counter--
	if cdl.counter < 0 {
		panic("counter should never be less than zero")
	}

	// If counter is zero, broadcast to all waiting goroutines
	if cdl.counter == 0 {
		cdl.cond.Broadcast()
	}
}

func (cdl *CountDownLatch) Wait() {
	cdl.m.Lock()
	defer cdl.m.Unlock()

	// Wait until the counter is zero
	for cdl.counter > 0 {
		cdl.cond.Wait()
	}
}
