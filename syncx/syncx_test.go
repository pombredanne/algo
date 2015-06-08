package syncx_test

import (
	"github.com/larsmans/algo/syncx"
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
