// Copyright 2014-2015 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"errors"
	"fmt"
	"io"
)

const bufsize = 32 * 1024

type scanToken int

const (
	scanIllegalToken scanToken = iota
	scanBeginObject
	scanEndObject
	scanBeginArray
	scanEndArray
	scanStringLit
	scanBoolLit
	scanIntLit
	scanFloatLit
	scanNullVal
	scanComma
)

var scanTokenToString = map[scanToken]string{
	scanIllegalToken: "illegal token",
	scanBeginObject:  "{",
	scanEndObject:    "}",
	scanBeginArray:   "[",
	scanEndArray:     "]",
	scanStringLit:    "string literal",
	scanBoolLit:      "boolean literal",
	scanIntLit:       "integer literal",
	scanFloatLit:     "integer literal",
	scanNullVal:      "null",
	scanComma:        "comma",
}

func (tok scanToken) String() string {
	if str, ok := scanTokenToString[tok]; ok {
		return str
	}
	return "invalid scan token"
}

type scanner struct {
	r   io.Reader
	err error

	row int64
	col int64

	buf []byte
	pos int // position inside the buffer; must be -1 by default

	eof bool // true when the reader has been reached
}

func newScanner(r io.Reader) *scanner {
	return &scanner{r: r, buf: make([]byte, bufsize), pos: -1}
}

func (scan *scanner) nextKey() (string, error) {
	if err := scan.ignoreWhitespaces(); err != nil {
		if err == io.EOF {
			return "", errors.New("expected key, found EOF")
		}
		return "", err
	}

	if c, err := scan.read(); err != nil {
		return "", err
	} else if c != '"' {
		return "", fmt.Errorf("expected '\"', found '%c'", c)
	}

	var key string
	origPos := scan.pos
	for {
		c, err := scan.read()
		if err == io.EOF {
			return "", errors.New("expected key, found EOF")
		} else if err != nil {
			return "", err
		}

		// TODO: check character

		if c == '"' {
			break
		}

		if scan.pos >= len(scan.buf)-1 {
			key += string(scan.buf[origPos:])
			origPos = 0
		}
	}

	key += string(scan.buf[origPos : scan.pos-1])

	if err := scan.ignoreWhitespaces(); err != nil {
		if err == io.EOF {
			return "", errors.New("expected ':', found EOF")
		}
		return "", err
	}

	if c, err := scan.read(); err == io.EOF {
		return "", errors.New("expected ':', found EOF")
	} else if err != nil {
		return "", err
	} else if c != ':' {
		return "", fmt.Errorf("expected ':', found '%c'", c)
	}

	return key, nil
}

func (scan *scanner) nextValue() ([]byte, scanToken, error) {
	if err := scan.ignoreWhitespaces(); err != nil {
		if err == io.EOF {
			return nil, scanIllegalToken, errors.New("expected value, found EOF")
		}
		return nil, scanIllegalToken, err
	}

	c, err := scan.read()
	if err != nil {
		if err == io.EOF {
			return nil, scanIllegalToken, errors.New("expected value, found EOF")
		}
		return nil, scanIllegalToken, err
	}

	switch {
	case isDigit(c):
		// TODO: read digit
	case c == 't' || c == 'f':
		// TODO: read bool
	case c == '{': // beginning of an object literal
		return nil, scanBeginObject, nil
	case c == '}': // ending of an object literal
		return nil, scanEndObject, nil
	case c == '[': // beginning of an array literal
		return nil, scanBeginArray, nil
	case c == ']': // ending of an array literal
		return nil, scanEndArray, nil
	case c == '"': // beginning or ending of a string literal
		str, err := scan.readString()
		return str, scanStringLit, err
	case c == 'n': // null
		// TODO: read null
	case c == ',':
		return nil, scanComma, nil
	}
	return nil, scanIllegalToken, fmt.Errorf("illegal value '%c'", c)
}

// peek reads the next value without consuming it.
func (scan *scanner) peek() (byte, error) {
	c, err := scan.read()
	if err != nil {
		return 0, err
	}
	scan.back()
	return c, nil
}

func (scan *scanner) read() (byte, error) {
	if scan.pos == -1 || (scan.pos > len(scan.buf)-1 && !scan.eof) {
		n, err := scan.r.Read(scan.buf)
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			scan.eof = true
		}
		if n < bufsize {
			scan.buf = scan.buf[:n]
		}
		scan.pos = 0
	}
	if scan.pos >= len(scan.buf) {
		return 0, io.EOF
	}
	b := scan.buf[scan.pos]
	scan.pos++
	return b, nil
}

// readString reads a string literal.
func (scan *scanner) readString() ([]byte, error) {
	// length of the read string
	// XXX: overflow?
	var strLen int

	buf := make([]byte, bufsize)
	for {
		c, err := scan.read()
		if err != nil {
			return nil, err
		}
		if c == '"' {
			if strLen == 0 {
				// empty string
				return []byte{}, nil
			}
			if buf[strLen-1] != '\\' {
				break
			}
		}
		if strLen == len(buf) {
			buf = append(buf, c)
		} else {
			buf[strLen] = c
		}
		strLen++
	}

	return buf[:strLen], nil
}

func (scan *scanner) back() {
	if scan.pos > 0 {
		scan.pos--
	}
}

// ignoreWhitespaces consumes all whitespaces.
func (scan *scanner) ignoreWhitespaces() error {
	var c byte
	var err error
	// XXX: refactor when the tests pass
	for {
		if c, err = scan.read(); err != nil || !isWhitespace(c) {
			break
		}
		// Nothing, we just skip whitespaces.
	}

	// Since a non-whitespace has been read, we have to put it back to the
	// buffer so that it can be read again.
	scan.back()

	if err != io.EOF {
		return err
	}
	return nil
}

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\n' || c == '\t' || c == '\r'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// TODO: implement
func isUnicodeChar(c byte) bool {
	return false
}
