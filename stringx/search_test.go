package stringx

import (
	"strings"
	"testing"
)

func TestBMH(t *testing.T) {
	haystack := "do you like seafood?"
	for _, pattern := range []string{
		"foo", "bar", "u like", "?", "??", haystack, "", "like an arrow",
	} {
		comp := CompileBMH(pattern)
		if s := comp.String(); s != pattern {
			t.Errorf("expected %q, got %q from String", pattern, s)
		}
		i := strings.Index(haystack, pattern)
		if got := comp.Index(haystack); got != i {
			t.Errorf("expected %d for %q, got %d", i, pattern, got)
		}
	}
}

func BenchmarkBMH(b *testing.B) {
	b.StopTimer()
	haystack := "foobarbaz Boyer-Moore"
	for i := 0; i < 1000; i++ {
		haystack += "t266813975i61275916572819"
	}
	needle := "Boyer-Moore-Horspool"
	haystack += needle

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CompileBMH(needle).Index(haystack)
		//strings.Index(haystack, needle) // About 4× slower.
	}
}
