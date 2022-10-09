package ch11

import (
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev := Reverse(tc.in)
		require.Equal(t, tc.want, rev)
	}
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, t := range testcases {
		f.Add(t)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
