// Copyright 2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

// Package netstring implements D.J. Bernstein's netstring encoding.
//
// Netstring is a simple encoding of byte strings for network applications.
// See http://cr.yp.to/proto/netstrings.txt for the format.
package netstring

import (
	"errors"
	"fmt"
	"io"
)

// Decode decodes a netstring at the start of src.
//
// Decode reports the represented string, the number of characters consumed
// from src, and/or an error. The returned string s will be a slice of src.
//
// If it is known that the start of src holds a netstring, it can be parsed
// out using:
//
//	s, n, err := netstring.Decode(src)
//	if err != nil {
//		// Handle error
//	}
//	src = src[n:]
func Decode(src []byte) (s []byte, n int, err error) {
	var length int
	for {
		c := src[n]
		n++
		if c == ':' {
			break
		}
		if length, err = add(length, c); err != nil {
			return
		}
	}
	if n == 1 {
		err = EmptyLen
	}
	if err != nil {
		return
	}
	src = src[n:]
	if len(src) < length+1 {
		err = TooShort
		return
	}
	if src[length] != ',' {
		err = NoComma
		return
	}
	s = src[:length]
	n += length + 1 // +1 for ','
	return
}

// Like Decode, but operates on a string and returns a string.
func DecodeString(src string) (s string, n int, err error) {
	var length int
	for {
		c := src[n]
		n++
		if c == ':' {
			break
		}
		if length, err = add(length, c); err != nil {
			return
		}
	}
	if n == 1 {
		err = EmptyLen
	}
	if err != nil {
		return
	}
	src = src[n:]
	if len(src) < length+1 {
		err = TooShort
		return
	}
	if src[length] != ',' {
		err = NoComma
		return
	}
	s = src[:length]
	n += length + 1 // +1 for ','
	return
}

// Read reads a netstring-encoded string from r.
//
// If maxlen > 0 and the decoded string is longer than maxlen, an error of
// type TooLong is returned after the length and the initial ':' is read from
// r. maxlen <= 0 means that the decoded string may have any length.
//
// If buf is non-nil, then Read may use it as scratch space to construct s.
func Read(r io.Reader, maxlen int, buf []byte) (s []byte, n int, err error) {
	if maxlen <= 0 {
		maxlen = maxint
	}

	var readByte func() (byte, error)
	if br, ok := r.(io.ByteReader); ok {
		readByte = br.ReadByte
	} else {
		readByte = func() (c byte, err error) {
			var buf [1]byte
			_, err = r.Read(buf[:])
			c = buf[0]
			return
		}
	}

	var length int
	for {
		var c byte
		c, err = readByte()
		if err == io.EOF && n > 0 {
			err = io.ErrUnexpectedEOF
		}
		if err != nil {
			return
		}

		n++
		if c == ':' {
			break
		}

		if length, err = add(length, c); err != nil {
			return
		}
	}
	if n == 1 {
		err = EmptyLen
		return
	}
	if length > maxlen {
		err = TooLong(length)
		return
	}

	if cap(buf) >= length+1 {
		s = buf[:length+1]
	} else {
		s = make([]byte, length+1)
	}

	consumed, err := io.ReadFull(r, s)
	n += consumed
	if err != nil {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
		return
	}
	if s[length] != ',' {
		err = NoComma
	}
	s = s[:length]
	return
}

func add(length int, digit byte) (int, error) {
	if digit < '0' || digit > '9' {
		return -1, NonDigit
	}
	oldlength := length
	length = 10*length + int(digit-'0')
	if length < oldlength {
		return -1, Overflow
	}
	return length, nil
}

// Encode encodes p in the netstring format.
func Encode(p []byte) []byte {
	return []byte(fmt.Sprintf("%d:%s,", len(p), p))
}

// Encode encodes s in the netstring format.
func EncodeString(s string) string {
	return fmt.Sprintf("%d:%s,", len(s), s)
}

// Write p to w in netstring encoding. Reports the number of bytes written.
func Write(p []byte, w io.Writer) (n int, err error) {
	return fmt.Fprintf(w, "%d:%s,", len(p), p)
}

// Like Write, but operates on a string.
func WriteString(s string, w io.Writer) (n int, err error) {
	return fmt.Fprintf(w, "%d:%s,", len(s), s)
}

// Errors of this type are returned by Read when the decoded input is longer
// than maxlen. The integer value is the actual length of the input.
type TooLong int

func (e TooLong) Error() string {
	return fmt.Sprintf("%d > maxlen", int(e))
}

// Decoding errors. Instead of TooShort, Read returns io.ErrUnexpectedEOF.
var (
	EmptyLen = errors.New("empty length field")
	NoComma  = errors.New("no comma terminator on netstring")
	NonDigit = errors.New("non-digit in length field")
	Overflow = fmt.Errorf("string length > %d", maxint)
	TooShort = errors.New("input too short")
	maxint   = int(^uint(0) >> 1)
)
