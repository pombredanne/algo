// Copyright 2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

// Package syncx provides synchronization algorithms.
package syncx

import "sync/atomic"

type TryMutex struct {
	v int32
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
