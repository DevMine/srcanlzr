// Copyright 2014-2015 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"bytes"
	"errors"
	"testing"
)

func TestIgnoreWhitespaces(t *testing.T) {
	buf := bytes.NewBufferString("     x")
	scan := newScanner(buf)
	if err := scan.ignoreWhitespaces(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c := scan.buf[scan.pos:]; len(c) != 1 || c[0] != 'x' {
		t.Errorf("scan.ignoreWhitespace: found '%s', expected 'x'", c)
	}
}

func TestIsWhitespace(t *testing.T) {
	ws := []byte{' ', '\n', '\t', '\r'}
	for _, c := range ws {
		if !isWhitespace(c) {
			t.Errorf("'%c' should be considered as a whitespace", c)
		}
	}

	nows := []byte{'a', ',', ';', ':', '0', '\v'}
	for _, c := range nows {
		if isWhitespace(c) {
			t.Errorf("'%c' should not be considered as a whitespace", c)
		}
	}
}

func TestIsDigit(t *testing.T) {
	digits := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for _, c := range digits {
		if !isDigit(c) {
			t.Errorf("'%c' should be considered as a digit", c)
		}
	}

	var c byte
	for c = 'a'; c <= 'z'; c++ {
		if isDigit(c) {
			t.Errorf("'%c' should not be considered as a digit", c)
		}
	}
}

func TestNextKey(t *testing.T) {
	validKeys := []string{
		`"foo":`,
		`"foo"  :`,
		`"foo": "bar"`,
	}
	for _, keyInput := range validKeys {
		buf := bytes.NewBufferString(keyInput)
		scan := newScanner(buf)
		key, err := scan.nextKey()
		if err != nil {
			t.Error(err)
		}
		if key != "foo" {
			t.Errorf("nextKey: found '%s', expected 'foo'", key)
		}
	}

	invalidKeys := map[string]error{
		`"foo"`: errors.New("expected ':', found EOF"),
		`"foo`:  errors.New("expected key, found EOF"),
		`foo`:   errors.New("expected '\"', found 'f'"),
	}
	for keyInput, expectedErr := range invalidKeys {
		buf := bytes.NewBufferString(keyInput)
		scan := newScanner(buf)
		_, err := scan.nextKey()
		if err == nil {
			t.Errorf("nextKey: '%s' is expected to return the error '%v'", keyInput, expectedErr)
			continue
		}
		if err.Error() != expectedErr.Error() {
			t.Errorf("nextKey: found error \"%v\", expected error \"%v\"", err, expectedErr)
		}
	}
}
