// Copyright 2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

package extsort

import (
	"github.com/larsmans/algo/extmath/intmath"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
)

// Sort data using multiple goroutines.
//
// The implementation is recursive and switches to a serial implementation
// when it has fewer than cutoff items to sort.
//
// Parallel(data, cutoff) is equivalent to sort.Sort(data) if the
// sort.Interface methods on data are thread-safe. It uses a global pool of
// goroutines, so invoking this function from multiple goroutines in parallel
// does not use more than a constant times runtime.GOMAXPROCS(-1) goroutines.
func Parallel(data sort.Interface, cutoff int) {
	wg := new(sync.WaitGroup)
	n := data.Len()
	if cutoff < 3 {
		cutoff = 3 // medianOfThree cannot handle fewer items.
	}
	pqsort(data, 0, n, cutoff, 2*intmath.Log2(n), wg)
	wg.Wait()
}

var availableProcs int32

func init() {
	atomic.StoreInt32(&availableProcs, int32(runtime.GOMAXPROCS(-1)) * 2)
}

func pqsort(data sort.Interface, lo, hi, cutoff, depth int, wg *sync.WaitGroup) {
	for hi-lo > cutoff && depth > 0 {
		p := medianOfThree(data, lo, hi)
		// TODO: implement a parallel partition function.
		p = partition(data, lo, hi, p)

		depth--
		var losmall, hismall int
		if p < (hi-lo)/2 {
			losmall, hismall = lo, p
			lo = p + 1
		} else {
			losmall, hismall = p+1, hi
			hi = p
		}

		// Do the "small" recursion in a freshly spawned goroutine, if we
		// don't already have maxworkers running.
		if atomic.AddInt32(&availableProcs, -1) >= 0 {
			wg.Add(1)
			go func() {
				pqsort(data, losmall, hismall, cutoff, depth, wg)
				wg.Done()
				atomic.AddInt32(&availableProcs, 1)
			}()
		} else {
			atomic.AddInt32(&availableProcs, 1)
			pqsort(data, losmall, hismall, cutoff, depth, wg)
		}
	}
	if hi-lo > 1 {
		sort.Sort(islice{data, lo, hi})
	}
}

// Slice of sort.Interface.
type islice struct {
	a      sort.Interface
	lo, hi int
}

func (s islice) Len() int           { return s.hi - s.lo }
func (s islice) Less(i, j int) bool { return s.a.Less(i+s.lo, j+s.lo) }
func (s islice) Swap(i, j int)      { s.a.Swap(i+s.lo, j+s.lo) }
