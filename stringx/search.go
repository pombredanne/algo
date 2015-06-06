package stringx

// A Pattern represents a preprocessed string for use in substring searching.
type Pattern interface {
	// Returns the index of the first occurrence of pattern.String() in s,
	// or -1 if it does not occur.
	Index(s string) int

	// Returns the substring being searched for.
	String() string
}

type bmh struct {
	needle string
	table  [256]int
}

// Preprocess pattern for string searching using the Boyer-Moore-Horspool
// algorithm.
func CompileBMH(pattern string) Pattern {
	compiled := &bmh{needle: pattern}
	table := &compiled.table

	for i := range table {
		table[i] = len(pattern)
	}
	if len(pattern) > 0 {
		for i, c := range pattern[:len(pattern)-1] {
			table[c] = len(pattern) - i - 1
		}
	}
	return compiled
}

func (pattern *bmh) Index(s string) int {
	needle := pattern.needle
	patlen, slen := len(needle), len(s)

	if patlen == 0 {
		return 0
	}

	last := needle[patlen-1]
	for i := patlen; i <= slen; {
		c := s[i-1]
		if c == last && s[i-patlen:i] == needle {
			return i - patlen
		}
		i += pattern.table[c]
	}
	return -1
}

// Returns the string that was passed to Compile to produce pattern.
func (pattern *bmh) String() string {
	return pattern.needle
}
