// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the goauthors-LICENSE file.

package outside

import (
	"syscall" // for EINVAL (exists in windows and linux)
	"unicode/utf16"
)

func utf16PtrFromString(s string) (*uint16, error) {
	a, err := utf16FromString(s)
	if err != nil {
		return nil, err
	}
	return &a[0], nil
}

// Copied from syscall_windows.go
// UTF16FromString returns the UTF-16 encoding of the UTF-8 string
// s, with a terminating NUL added. If s contains a NUL byte at any
// location, it returns (nil, EINVAL).
func utf16FromString(s string) ([]uint16, error) {
	for i := 0; i < len(s); i++ {
		if s[i] == 0 {
			return nil, syscall.EINVAL
		}
	}
	return utf16.Encode([]rune(s + "\x00")), nil
}

// Copied from syscall_windows.go
// UTF16ToString returns the UTF-8 encoding of the UTF-16 sequence s,
// with a terminating NUL removed.
func utf16ToString(s []uint16) string {
	for i, v := range s {
		if v == 0 {
			s = s[0:i]
			break
		}
	}
	return string(utf16.Decode(s))
}

// Copied from syscall_windows.go
// BytePtrFromString returns a pointer to a NUL-terminated array of
// bytes containing the text of s. If s contains a NUL byte at any
// location, it returns (nil, EINVAL).
func bytePtrFromString(s string) (*byte, error) {
	a, err := byteSliceFromString(s)
	if err != nil {
		return nil, err
	}
	return &a[0], nil
}

// Copied from syscall_windows.go
// ByteSliceFromString returns a NUL-terminated slice of bytes
// containing the text of s. If s contains a NUL byte at any
// location, it returns (nil, EINVAL).
func byteSliceFromString(s string) ([]byte, error) {
	for i := 0; i < len(s); i++ {
		if s[i] == 0 {
			return nil, syscall.EINVAL
		}
	}
	a := make([]byte, len(s)+1)
	copy(a, s)
	return a, nil
}
