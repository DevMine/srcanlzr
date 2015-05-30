// Copyright 2014-2015 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"bytes"
	"errors"
	"testing"
)

func TestDecodeStringsList(t *testing.T) {
	expected := []string{"foo", "bar"}
	buf := bytes.NewBufferString(`["foo","bar"]`)
	dec := newDecoder(buf)
	sl := dec.decodeStringsList()
	if dec.err != nil {
		t.Fatal(dec.err)
	}
	if !stringsSlicesEquals(sl, expected) {
		t.Errorf("decodeStringsList: found %v, expected %v", sl, expected)
	}
}

func TestDecodeLanguages(t *testing.T) {
	expected := []*Language{
		&Language{
			Lang:      Go,
			Paradigms: []string{Compiled, Concurrent, Imperative, Structured},
		},
		&Language{
			Lang:      Java,
			Paradigms: []string{ObjectOriented, Structured, Imperative, Generic, Reflective, Concurrent},
		},
	}
	buf := bytes.NewBufferString(`[
		{
			"language": "go",
			"paradigms": [
				"compiled",
				"concurrent",
				"imperative",
				"structured"
			]
		},
		{
			"language": "java",
			"paradigms": [
				"object oriented",
				"structured",
				"imperative",
				"generic",
				"reflective",
				"concurrent"
			]
		}
	]`)
	dec := newDecoder(buf)
	langs := dec.decodeLanguages()
	if dec.err != nil {
		t.Fatal(dec.err)
	}
	if l := len(langs); l != 2 {
		t.Fatalf("decodeLanguages: found %d languages, expected 2", l)
	}
	compareLanguages(t, langs[0], expected[0])
	compareLanguages(t, langs[1], expected[1])
}

func TestDecodeLanguage(t *testing.T) {
	expected := &Language{
		Lang:      Go,
		Paradigms: []string{Compiled, Concurrent, Imperative, Structured},
	}
	buf := bytes.NewBufferString(`{"language": "go", "paradigms": ["compiled", "concurrent", "imperative", "structured"]}`)
	dec := newDecoder(buf)
	lang := dec.decodeLanguage()
	if dec.err != nil {
		t.Fatal(dec.err)
	}
	compareLanguages(t, lang, expected)
}

func TestUnmarshalString(t *testing.T) {
	dec := newDecoder(new(bytes.Buffer))
	if str, err := dec.unmarshalString([]byte("foo")); err != nil {
		t.Error(err)
	} else if str != "foo" {
		t.Errorf("unmarshalString 'foo': found '%s', expected 'foo'", str)
	}

	expectedErr := errors.New("unable to unmarshal string: data is nil")
	if _, err := dec.unmarshalString(nil); err == nil {
		t.Errorf("unmarshalString nil: found no error, expected \"%v\"", expectedErr)
	} else if err.Error() != expectedErr.Error() {
		t.Errorf("unmarshalString nil: found error \"%v\", expected \"%v\"", err, expectedErr)
	}
}

func TestAssertNewObject(t *testing.T) {
	type returnStatus struct {
		status bool
		err    error
	}
	input := map[string]returnStatus{
		"{": returnStatus{true, nil},
		"":  returnStatus{false, errors.New("expected value, found EOF")},
		"[": returnStatus{false, errors.New("expected object, found '['")},
	}

	for in, ret := range input {
		buf := bytes.NewBufferString(in)
		dec := newDecoder(buf)
		if status := dec.assertNewObject(); status != ret.status {
			t.Errorf("assertNewObject '%s': found %v, expected %v", in, status, ret.status)
		}
		if err := dec.err; err != nil && err.Error() != ret.err.Error() {
			t.Errorf("assertNewObject '%s': found error \"%v\", expected error \"%v\"", in, dec.err, ret.err)
		}
	}
}

func TestAssertNewArray(t *testing.T) {
	type returnStatus struct {
		status bool
		err    error
	}
	input := map[string]returnStatus{
		"[": returnStatus{true, nil},
		"":  returnStatus{false, errors.New("expected value, found EOF")},
		"{": returnStatus{false, errors.New("expected array, found '{'")},
	}

	for in, ret := range input {
		buf := bytes.NewBufferString(in)
		dec := newDecoder(buf)
		if status := dec.assertNewArray(); status != ret.status {
			t.Errorf("assertNewArray '%s': found %v, expected %v", in, status, ret.status)
		}
		if err := dec.err; err != nil && err.Error() != ret.err.Error() {
			t.Errorf("assertNewArray '%s': found error \"%v\", expected error \"%v\"", in, dec.err, ret.err)
		}
	}
}

func TestIsEndObject(t *testing.T) {
	type returnStatus struct {
		status bool
		err    error
	}
	input := map[string]returnStatus{
		"}":  returnStatus{true, nil},
		",":  returnStatus{false, nil},
		"42": returnStatus{false, errors.New("expected 'comma', found 'integer literal'")},
	}

	for in, ret := range input {
		buf := bytes.NewBufferString(in)
		dec := newDecoder(buf)
		if status := dec.isEndObject(); status != ret.status {
			t.Errorf("isEndObject '%s': found %v, expected %v", in, status, ret.status)
		}
		if err := dec.err; err != nil && err.Error() != ret.err.Error() {
			t.Errorf("isEndObject '%s': found error \"%v\", expected error \"%v\"", in, dec.err, ret.err)
		}
	}
}

func TestIsEndArray(t *testing.T) {
	type returnStatus struct {
		status bool
		err    error
	}
	input := map[string]returnStatus{
		"]":  returnStatus{true, nil},
		",":  returnStatus{false, nil},
		"42": returnStatus{false, errors.New("expected 'comma', found 'integer literal'")},
	}

	for in, ret := range input {
		buf := bytes.NewBufferString(in)
		dec := newDecoder(buf)
		if status := dec.isEndArray(); status != ret.status {
			t.Errorf("isEndArray '%s': found %v, expected %v", in, status, ret.status)
		}
		if err := dec.err; err != nil && err.Error() != ret.err.Error() {
			t.Errorf("isEndArray '%s': found error \"%v\", expected error \"%v\"", in, dec.err, ret.err)
		}
	}
}

func TestIsEmptyObject(t *testing.T) {
	type returnStatus struct {
		status bool
		err    error
	}
	input := map[string]returnStatus{
		"}":        returnStatus{true, nil},
		"":         returnStatus{false, errors.New("unexpected EOF")},
		"\"foo\":": returnStatus{false, nil},
	}

	for in, ret := range input {
		buf := bytes.NewBufferString(in)
		dec := newDecoder(buf)
		if status := dec.isEmptyObject(); status != ret.status {
			t.Errorf("isEmptyObject '%s': found %v, expected %v", in, status, ret.status)
		}
		if err := dec.err; err != nil && err.Error() != ret.err.Error() {
			t.Errorf("isEmptyObject '%s': found error \"%v\", expected error \"%v\"", in, dec.err, ret.err)
		}
	}
}

func TestIsEmptyArray(t *testing.T) {
	type returnStatus struct {
		status bool
		err    error
	}
	input := map[string]returnStatus{
		"]":  returnStatus{true, nil},
		"":   returnStatus{false, errors.New("unexpected EOF")},
		"42": returnStatus{false, nil},
	}

	for in, ret := range input {
		buf := bytes.NewBufferString(in)
		dec := newDecoder(buf)
		if status := dec.isEmptyArray(); status != ret.status {
			t.Errorf("isEmptyArray '%s': found %v, expected %v", in, status, ret.status)
		}
		if err := dec.err; err != nil && err.Error() != ret.err.Error() {
			t.Errorf("isEmptyArray '%s': found error \"%v\", expected error \"%v\"", in, dec.err, ret.err)
		}
	}
}

func compareLanguages(t *testing.T, found, expected *Language) {
	if found == nil {
		t.Fatal("decodeLanguage: should not be nil")
	}

	if expected.Lang != found.Lang {
		t.Errorf("decodeLanguage.Lang: found '%s', expected '%s'", found.Lang, expected.Lang)
	}
	if !stringsSlicesEquals(expected.Paradigms, found.Paradigms) {
		t.Errorf("decodeLanguage.Paradigms: found '%v', expected '%v'", found.Paradigms, expected.Paradigms)
	}
}

func stringsSlicesEquals(sl1, sl2 []string) bool {
	if len(sl1) != len(sl2) {
		return false
	}
	for i := 0; i < len(sl1); i++ {
		if sl1[i] != sl2[i] {
			return false
		}
	}
	return true
}
