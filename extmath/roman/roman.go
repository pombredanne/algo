// Copyright 2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

// Package roman implements routines for handling Roman numerals.
package roman

import "bytes"

// Format n as a Roman numeral.
//
// Subtractive notation is used, so 40 becomes "XL", not "XXXX".
//
// This function uses ASCII uppercase letters, not the Unicode range for Roman
// numerals, which exist for backward compatibility only.
func Itoa(n uint) []byte {
	var buf bytes.Buffer

	for n >= 1000 {
		buf.Write([]byte("M"))
		n -= 1000
	}
	if n >= 900 {
		buf.Write([]byte("CM"))
		n -= 900
	}
	if n >= 500 {
		buf.Write([]byte("D"))
		n -= 500
	}
	if n >= 400 {
		buf.Write([]byte("CD"))
		n -= 400
	}
	for n >= 100 {
		buf.Write([]byte("C"))
		n -= 100
	}
	if n >= 90 {
		buf.Write([]byte("XC"))
		n -= 90
	}
	if n >= 50 {
		buf.Write([]byte("L"))
		n -= 50
	}
	if n >= 40 {
		buf.Write([]byte("XL"))
		n -= 40
	}
	for n >= 10 {
		buf.Write([]byte("X"))
		n -= 10
	}
	var rest []byte
	switch n {
	case 1:
		rest = []byte("I")
	case 2:
		rest = []byte("II")
	case 3:
		rest = []byte("III")
	case 4:
		rest = []byte("IV")
	case 5:
		rest = []byte("V")
	case 6:
		rest = []byte("VI")
	case 7:
		rest = []byte("VII")
	case 8:
		rest = []byte("VIII")
	case 9:
		rest = []byte("IX")
	}
	buf.Write(rest)
	return buf.Bytes()
}
