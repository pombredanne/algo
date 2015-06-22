package syncx_test

import (
	"github.com/larsmans/algo/syncx"
	"runtime"
	"sync"
	"testing"
)

func TestTryMutex(t *testing.T) {
	var mu syncx.TryMutex
	if !mu.TryLock() {
		t.Fatal("expected TryLock to work")
	}
	wait := make(chan struct{})
	go func() {
		for i := 0; i < 2; i++ {
			if mu.TryLock() {
				t.Fatal("expected TryLock to fail")
			}
		}
		wait <- struct{}{}
	}()
	<-wait

	mu.Unlock()
	if !mu.TryLock() {
		t.Fatal("expected TryLock to work")
	}
	mu.Unlock()

	defer func() {
		err := recover()
		if err == nil {
			t.Fatal("expected Unlock to panic")
		}
		t.Logf("got %T %q", err, err)
	}()
	mu.Unlock()
}

func TestTryMutexLocker(t *testing.T) {
	var mu syncx.TryMutex
	var l sync.Locker = &mu

	l.Lock()
	ch := make(chan struct{})
	go func() {
		l.Lock()
		ch <- struct{}{}
	}()
	runtime.Gosched()

	if mu.TryLock() {
		t.Fatal("mu should be locked")
	}
	l.Unlock()
	<-ch
	l.Unlock()
	if !mu.TryLock() {
		t.Fatal("mu should be unlocked")
	}
}
