// Copyright 2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

// Package syncx provides synchronization algorithms.
package syncx

import (
	"runtime"
	"sync/atomic"
)

// A mutex with a TryLock method that returns immediately if it is already
// locked.
//
// The zero value represents an unlocked mutex.
type TryMutex struct {
	v int32
}

// Lock mu.
//
// This spins on TryLock, yielding the processor on every unsuccessful call.
func (mu *TryMutex) Lock() {
	for !mu.TryLock() {
		runtime.Gosched()
	}
}

// Try to lock l.
//
// A call to this method returns true if the call locked the mutex, false if
// it was already locked.
func (mu *TryMutex) TryLock() bool {
	return atomic.CompareAndSwapInt32(&mu.v, 0, 1)
}

// Unlock l.
//
// If l was not locked, a runtime error occurs.
func (mu *TryMutex) Unlock() {
	if !atomic.CompareAndSwapInt32(&mu.v, 1, 0) {
		panic("Unlock called, but mutex was not locked")
	}
}
