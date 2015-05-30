// Copyright 2014-2015 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/DevMine/repotool/model"
)

// Decode a JSON encoded src.Project read from r.
func Decode(r io.Reader) (*Project, error) {
	dec := newDecoder(r)
	return dec.decode()
}

// Decode a JSON encoded src.Project read from a given file.
func DecodeFile(path string) (*Project, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Decode(f)
}

type decoder struct {
	scan *scanner
	buf  []byte
	err  error
}

// newDecoder creates a new JSON decoder that reads from r.
func newDecoder(r io.Reader) *decoder {
	return &decoder{scan: newScanner(r)}
}

// decode decodes JSON input into a src.Project structure.
func (dec *decoder) decode() (*Project, error) {
	if _, tok, err := dec.scan.nextValue(); err != nil {
		return nil, dec.errorf(err)
	} else if tok != scanBeginObject {
		return nil, dec.errorf("expected '{' as first character")
	}

	prj := dec.decodeProject()
	if dec.err != nil {
		return nil, dec.errorf(dec.err)
	}

	if _, tok, err := dec.scan.nextValue(); err != nil {
		return nil, dec.errorf(err)
	} else if tok != scanEndObject {
		return nil, dec.errorf("expected '}' as last character")
	}

	return prj, nil
}

func (dec *decoder) errorf(v interface{}) error {
	return fmt.Errorf("malformed json: %v", v)
}

// decodeProject decodes a project object.
func (dec *decoder) decodeProject() *Project {
	prj := &Project{}

	for {
		key, err := dec.scan.nextKey()
		if err != nil {
			if err == io.EOF {
				break
			}
			dec.err = err
			return nil
		}
		if key == "" {
			dec.err = errors.New("empty key")
			return nil
		}

		val, tok, err := dec.scan.nextValue()
		if err != nil {
			dec.err = err
			return nil
		}

		switch key {
		case "packages":
			prj.Packages = dec.decodePackages()
		case "languages":
			prj.Langs = dec.decodeLanguages()
		case "repository":
			prj.Repo = dec.decodeRepository()
		case "loc":
			if tok != scanIntLit {
				dec.err = fmt.Errorf("expected integer literal, found %v", tok)
				return nil
			}
			prj.LoC, dec.err = dec.unmarshalInt(val)
		case "name":
			if tok != scanStringLit {
				dec.err = fmt.Errorf("expected string literal, found %v", tok)
				return nil
			}
			prj.Name, dec.err = dec.unmarshalString(val)
		default:
			dec.err = errors.New("unexpected value for project object")
		}

		if dec.err != nil {
			return nil
		}
	}
	return prj
}

// decodePackages decodes a list of package objects.
func (dec *decoder) decodePackages() []*Package {
	if !dec.assertNewArray() {
		return nil
	}

	pkgs := []*Package{}

	if dec.isEmptyArray() {
		return pkgs
	}
	if dec.err != nil {
		return nil
	}

	for {
		pkg := dec.decodePackage()
		if dec.err != nil {
			return nil
		}
		pkgs = append(pkgs, pkg)

		if dec.isEndArray() {
			break
		}
		if dec.err != nil {
			return nil
		}
	}

	return nil
}

// decoderPackage decodes a package object.
func (dec *decoder) decodePackage() *Package {
	if !dec.assertNewObject() {
		return nil
	}

	pkg := Package{}

	if dec.isEmptyObject() {
		return &pkg
	}
	if dec.err != nil {
		return nil
	}

	for {
		var key string
		key, dec.err = dec.scan.nextKey()
		if dec.err != nil {
			return nil
		}

		val, tok, err := dec.scan.nextValue()
		if err != nil {
			dec.err = err
			return nil
		}

		switch key {
		case "source_files":
			pkg.SrcFiles = dec.decodeSrcFiles()
		case "doc":
			pkg.Doc = dec.decodeStringsList()
		case "loc":
			if tok != scanIntLit {
				dec.err = fmt.Errorf("expected integer literal, found %v", tok)
				return nil
			}
			if pkg.LoC, dec.err = dec.unmarshalInt(val); dec.err != nil {
				return nil
			}
		case "name":
			if tok != scanStringLit {
				dec.err = fmt.Errorf("expected string literal, found %v", tok)
				return nil
			}
			if pkg.Name, dec.err = dec.unmarshalString(val); dec.err != nil {
				return nil
			}
		default:
			dec.err = errors.New("unexpected value for project object")
		}

		if dec.err != nil {
			return nil
		}

		if dec.isEndObject() {
			break
		}
		if err != nil {
			return nil
		}
	}

	return &pkg
}

// TODO: implement
func (dec *decoder) decodeSrcFiles() []*SrcFile {
	return nil
}

// decoderStringsList decodes a list of strings.
func (dec *decoder) decodeStringsList() []string {
	if !dec.assertNewArray() {
		return nil
	}

	sl := []string{}

	if dec.isEmptyArray() {
		return sl
	}
	if dec.err != nil {
		return nil
	}

	for {
		val, tok, err := dec.scan.nextValue()
		if err != nil {
			dec.err = err
			return nil
		}
		if tok != scanStringLit {
			dec.err = fmt.Errorf("expected string, found %v", tok)
			return nil
		}
		sl = append(sl, string(val))

		if dec.isEndArray() {
			break
		}
		if dec.err != nil {
			return nil
		}
	}
	return sl
}

// decodeLanguages decodes a list of languages.
func (dec *decoder) decodeLanguages() []*Language {
	if !dec.assertNewArray() {
		return nil
	}

	ls := []*Language{}

	if dec.isEmptyArray() {
		return ls
	}
	if dec.err != nil {
		return nil
	}

	for {
		lang := dec.decodeLanguage()
		if dec.err != nil {
			return nil
		}

		ls = append(ls, lang)

		if dec.isEndArray() {
			break
		}
		if dec.err != nil {
			return nil
		}
	}

	return ls
}

// decodeLanguage decode a src.Language object.
func (dec *decoder) decodeLanguage() *Language {
	if !dec.assertNewObject() {
		return nil
	}

	lang := Language{}

	if dec.isEmptyObject() {
		return &lang
	}
	if dec.err != nil {
		return nil
	}

	for {
		key, err := dec.scan.nextKey()
		if err != nil {
			if err == io.EOF {
				break
			}
			dec.err = err
			return nil
		}
		if key == "" {
			dec.err = errors.New("empty key")
			return nil
		}

		val, tok, err := dec.scan.nextValue()
		if err != nil {
			dec.err = err
			return nil
		}

		switch key {
		case "paradigms":
			// Since the '[' character has been consumed, we need to step back
			// brefore calling decodeStringsList.
			dec.scan.back()
			lang.Paradigms = dec.decodeStringsList()
		case "language":
			if tok != scanStringLit {
				dec.err = fmt.Errorf("expected string literal, found %v", tok)
				return nil
			}
			lang.Lang, dec.err = dec.unmarshalString(val)
		default:
			dec.err = fmt.Errorf("unexpected value for the key '%s' of a language object", key)
		}

		if dec.err != nil {
			return nil
		}

		if dec.isEndObject() {
			break
		}
		if err != nil {
			return nil
		}
	}
	return &lang
}

// TODO: implement
func (dec *decoder) decodeRepository() *model.Repository {
	return nil
}

// TODO: implement
func (dec *decoder) unmarshalInt(data []byte) (int64, error) {
	return 0, nil
}

// unmarshalString unmarshals a bytes slice into a string.
func (dec *decoder) unmarshalString(data []byte) (string, error) {
	if data == nil {
		return "", errors.New("unable to unmarshal string: data is nil")
	}
	return string(data), nil
}

// assertNewObject makes sure that the next value is a new object. In other
// words, the next value must begin with a '{'. If it is not, it will set
// dec.err and return false.
func (dec *decoder) assertNewObject() bool {
	// Since Language is a JSON Object, we expect to find a '{' character.
	_, tok, err := dec.scan.nextValue()
	if err != nil {
		dec.err = err
		return false
	}
	if tok != scanBeginObject {
		dec.err = fmt.Errorf("expected object, found '%v'", tok)
		return false
	}
	return true
}

// assertNewArray makes sure that the next value is a new array. In order
// words, the next value must begin with a '['. If it is not, it will set
// dec.err and return false.
func (dec *decoder) assertNewArray() bool {
	_, tok, err := dec.scan.nextValue()
	if err != nil {
		dec.err = err
		return false
	}
	if tok != scanBeginArray {
		dec.err = fmt.Errorf("expected array, found '%v'", tok)
		return false
	}
	return true
}

// isEndObject returns true if the next value marks the end of the object
// ('}') and false otherwise. If it is false, the next value must be a
// comma. If not, it will set dec.err accordingly.
func (dec *decoder) isEndObject() bool {
	_, tok, err := dec.scan.nextValue()
	if err != nil {
		dec.err = err
		return false
	}
	if tok == scanEndObject {
		return true
	}
	if tok != scanComma {
		dec.err = fmt.Errorf("expected 'comma', found '%v'", tok)
	}
	return false
}

// isEndArray returns true if the next value marks the end of the array
// (']') and false otherwise. If it is false, the next value must be a
// comma. If not, it will set dec.err accordingly.
func (dec *decoder) isEndArray() bool {
	_, tok, err := dec.scan.nextValue()
	if err != nil {
		dec.err = err
		return false
	}
	if tok == scanEndArray {
		return true
	}
	if tok != scanComma {
		dec.err = fmt.Errorf("expected 'comma', found '%s'", tok)
	}
	return false
}

// isEmptyObject tests if the object is empty (no key/value pairs inside).
//
// This method does not consume any byte.
//
// If an error occurs, it returns false and set dec.err.
func (dec *decoder) isEmptyObject() bool {
	// The object can be empty, so we have to check for that and without
	// consuming the next byte.
	if b, err := dec.scan.peek(); err != nil {
		if err == io.EOF {
			dec.err = errors.New("unexpected EOF")
		} else {
			dec.err = err
		}
		return false
	} else if b == '}' {
		return true
	}
	return false
}

// isEmptyArray tests if the object is empty (no values inside).
//
// This method does not consume any byte.
//
// If an error occurs, it returns false and set dec.err.
func (dec *decoder) isEmptyArray() bool {
	if b, err := dec.scan.peek(); err != nil {
		if err == io.EOF {
			dec.err = errors.New("unexpected EOF")
		} else {
			dec.err = err
		}
		return false
	} else if b == ']' {
		return true
	}
	return false
}
