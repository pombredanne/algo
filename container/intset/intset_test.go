package intset

import "testing"

func TestIntSet(t *testing.T) {
	s := New(10)

	check := func(expect []int) {
		for _, k := range expect {
			if !s.Contains(k) {
				t.Errorf("%d expected but missing", k)
			}
		}

		expset := make(map[int]bool)
		for _, k := range expect {
			expset[k] = true
		}
		if len(expset) != s.Len() {
			t.Errorf("Len mismatch: expected %d, got %d", len(expset), s.Len())
		}

		s.Do(func (k int) {
			if !expset[k] {
				t.Errorf("%d not expected in set", k)
			}
			delete(expset, k)
		})
		for k := range expset {
			t.Errorf("%d expected but missed by Do", k)
		}
	}

	check([]int{})

	s.Add(1)
	s.Add(5)
	s.Add(2)
	s.Add(5)
	check([]int{1, 2, 5})

	s.Remove(5)
	s.Remove(0)
	check([]int{1, 2})

	s.Add(7)
	s.Add(0)
	check([]int{0, 1, 2, 7})
}
