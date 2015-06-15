package syncx

import "sync"

// A Barrier is a rendezvous point for a set of goroutines.
//
// The number of goroutines that are going to use the barrier must be known
// at construction time.
//
// A Barrier differs from a sync.WaitGroup in that a Barrier can be reused.
// A WaitGroup wg can be used as a one-off barrier by having goroutines
// synchronize using
//	wg.Done()
//	wg.Wait()
// However, when the last goroutines arrives, all goroutines immediately resume
// execution, without waiting for the WaitGroup to be reset. If any goroutine
// reaches its Done call before the reset, it panics. Barrier does not suffer
// from this problem and will makes goroutines Wait until the reset is
// complete.
type Barrier struct {
	cond   *sync.Cond
	lock   sync.Mutex
	max, n int
}

// Construct a new barrier for n goroutines.
func NewBarrier(n int) *Barrier {
	if n < 1 {
		panic("n must be ≥ 1")
	}
	b := &Barrier{max: n, n: n}
	b.cond = sync.NewCond(&b.lock)
	return b
}

// Returns the number of goroutines for which b is a barrier.
func (b *Barrier) Num() int {
	return b.max
}

// Wait until all n goroutines have arrived at the barrier.
//
// The barrier is reset to its initial state when the last goroutine has
// arrived, and before any of the waiting goroutines has resumed execution.
func (b *Barrier) Wait() {
	b.lock.Lock()
	b.n--
	if b.n == 0 {
		b.n = b.max
		b.cond.Broadcast()
	} else {
		b.cond.Wait()
	}
	b.lock.Unlock()
}
