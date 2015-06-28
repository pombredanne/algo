package rangemin

import (
	"regexp"
	"sort"
	"testing"
)

func TestRMQ(t *testing.T) {
	// https://en.wikipedia.org/wiki/Range_minimum_query
	data := sort.IntSlice{0, 5, 2, 5, 4, 3, 1, 6, 3}
	rmq := New(data)

	for _, c := range []struct {
		from, to, min int
	}{
		{0, 1, 0},
		{0, 4, 0},
		{0, len(data), 0},
		{1, 2, 1},
		{1, 3, 2},
		{3, 8, 6},
	} {
		if got := rmq.Min(c.from, c.to); got != c.min {
			t.Errorf("expected Min=%d for [%d,%d], got %d",
				c.min, c.from, c.to, got)
		}
	}

	matchPanic(func() {
		New(sort.IntSlice{})
	}, "Len", t)

	for _, offset := range []int{0, -1} {
		matchPanic(func() {
			rmq.Min(5, 5+offset)
		}, "i >= j", t)
	}
}

func matchPanic(f func(), pattern string, t *testing.T) {
	re := regexp.MustCompile(pattern)
	defer func() {
		switch x := recover().(type) {
		case nil:
			t.Fatal("no panic")
		case string:
			if !re.MatchString(x) {
				t.Errorf("%q does not match %q", x, pattern)
			}
		default:
			t.Fatal("wrong type %T, expected string", x)
		}
	}()

	f()
}
