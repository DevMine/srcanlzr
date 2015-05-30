// Copyright 2014-2015 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"bytes"
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
