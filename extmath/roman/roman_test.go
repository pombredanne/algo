package roman

import "testing"

func TestItoa(t *testing.T) {
	for _, c := range []struct {
		n uint
		s string
	}{
		{1, "I"}, {11, "XI"}, {14, "XIV"}, {27, "XXVII"}, {40, "XL"},
		{50, "L"}, {58, "LVIII"}, {63, "LXIII"}, {76, "LXXVI"}, {99, "XCIX"},
		{100, "C"}, {119, "CXIX"}, {405, "CDV"}, {500, "D"}, {901, "CMI"},
		{1000, "M"}, {1632, "MDCXXXII"}, {1954, "MCMLIV"}, {2014, "MMXIV"},
	} {
		if s := string(Itoa(c.n)); s != c.s {
			t.Errorf("expected %s for %d, got %s", c.s, c.n, s)
		}
	}
}

func BenchmarkRoman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := uint(1); n < 1000; n++ {
			Itoa(n)
		}
	}
}
