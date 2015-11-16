package netstring

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

type testcase struct{ encoded, decoded string }

var cases = []testcase{
	// Examples from http://cr.yp.to/proto/netstrings.txt.
	{"0:,", ""},
	{"12:hello world!,", "hello world!"},
	{"9:naïveté,", "naïveté"},
	// From https://en.wikipedia.org/wiki/Netstring.
	{encoded: "17:5:hello,6:world!,,", decoded: "5:hello,6:world!,"},
}

func init() {
	// Add a largish byte-string to cases.
	var b bytes.Buffer
	n := 9999
	for i := 0; i < n; i++ {
		b.WriteRune(rune(i))
	}
	s := b.String()
	cases = append(cases, testcase{
		encoded: strconv.Itoa(len(s)) + ":" + s + ",",
		decoded: s,
	})
}

func TestDecode(t *testing.T) {
	check := func(input, got, expected string, n int, err error) {
		if err != nil {
			t.Errorf("got error %q", err)
		}
		if n < len(input) {
			t.Errorf("consumed %d bytes instead of whole string (%d bytes)",
				n, len(input))
		}
		if got != expected {
			t.Errorf("expected %q, got %q", expected[:60], got[:60])
		}
	}

	for _, c := range cases {
		s, n, err := DecodeString(c.encoded)
		check(c.encoded, s, c.decoded, n, err)

		var b []byte
		b, n, err = Decode([]byte(c.encoded))
		s = string(b)
		check(c.encoded, s, c.decoded, n, err)
	}
}

func TestErrors(t *testing.T) {
	s := "hello, world!"
	enc := EncodeString("hello, world!")
	for _, n := range []int{6, 10} {
		b := bytes.NewBufferString(enc[:n])
		_, consumed, err := Read(b, 0, make([]byte, len(enc)))
		if err != io.ErrUnexpectedEOF {
			t.Errorf("expected io.ErrUnexpectedEOF, got %T", err)
		}
		if consumed != n {
			t.Errorf("expected to consume %d bytes, got %d", n, consumed)
		}
	}

	b := bytes.NewBufferString(enc)
	_, _, err := Read(b, 1, nil)
	if toolong, ok := err.(TooLong); ok {
		if int(toolong) != len(s) {
			t.Errorf("wrong length: expected %d, got %d",
				len(s), int(toolong))
		}
	} else {
		t.Errorf("expected TooLong, got %T", err)
	}

	for _, c := range []struct {
		err error
		enc string
	}{
		{Overflow, "999999999999999999999999999999999999"},
		{NonDigit, "6x;hello!,"},
		{EmptyLen, ":,"},
		{TooShort, "10:foo!,"},
		{NoComma, "2:...,"},
		{NoComma, "2:...."},
	} {
		for _, f := range []func() error{
			func() (err error) {
				_, _, err = Read(bytes.NewBufferString(c.enc), 0, nil)
				if err == io.ErrUnexpectedEOF {
					err = TooShort
				}
				return
			},
			func() (err error) {
				_, _, err = DecodeString(c.enc)
				return
			},
			func() (err error) {
				_, _, err = Decode([]byte(c.enc))
				return
			},
		} {
			if err := f(); err != c.err {
				t.Errorf("expected %v, got %T %q", c.err, err, err)
			}
		}
	}

	_, _, err = Read(&bytes.Buffer{}, 0, nil)
	if err != io.EOF {
		t.Errorf("expected io.EOF, got %T %q", err, err)
	}
}

type missingByteReader struct {
	buf *bytes.Buffer
}

func (r missingByteReader) Read(p []byte) (int, error) {
	return r.buf.Read(p)
}

func TestNonByteReader(t *testing.T) {
	for _, c := range cases {
		r := missingByteReader{bytes.NewBufferString(c.encoded)}
		dec, _, err := Read(r, len(c.decoded), nil)
		if string(dec) != c.decoded {
			t.Errorf("expected %q, got %q", c.decoded, dec)
		}
		if err != nil {
			t.Errorf("unexpected error %q", err)
		}
	}
}

func TestWrite(t *testing.T) {
	for _, c := range cases {
		var b bytes.Buffer
		n, err := Write([]byte(c.decoded), &b)
		if err != nil {
			t.Error(err)
		}
		if n != len(c.encoded) {
			t.Errorf("expected to write %d bytes, got %d", len(c.encoded), n)
		}

		b.Reset()
		n, err = WriteString(c.decoded, &b)
		if err != nil {
			t.Error(err)
		}
		if n != len(c.encoded) {
			t.Errorf("expected to write %d bytes, got %d", len(c.encoded), n)
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			DecodeString(c.encoded)
		}
	}
}

func BenchmarkRead(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()

	f, err := ioutil.TempFile("", "netstring_test")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(f.Name())

	for i := 0; i < 10; i++ {
		for _, c := range cases {
			Write([]byte(c.encoded), f)
		}
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if _, err = f.Seek(0, os.SEEK_SET); err != nil {
			b.Fatal(err)
		}
		var buf []byte
		for {
			buf, _, err = Read(f, 0, buf)
			if err == io.EOF {
				break
			} else if err != nil {
				b.Fatal(err)
			}
		}
	}
}
