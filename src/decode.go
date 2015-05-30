// Copyright 2014-2015 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/DevMine/repotool/model"
)

func Decode(r io.Reader) (*Project, error) {
	dec := newDecoder(r)
	return dec.decode()
}

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

func newDecoder(r io.Reader) *decoder {
	return &decoder{scan: newScanner(r)}
}

// TODO: implement
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

func (dec *decoder) decodeProject() *Project {
	prj := &Project{}
	pv := reflect.ValueOf(prj)

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

		switch {
		case key == "packages":
			prj.Packages = dec.decodePackages()
		case key == "languages":
			prj.Langs = dec.decodeLanguages()
		case key == "repository":
			prj.Repo = dec.decodeRepository()
		case tok == scanIntLit:
			var num int64
			num, dec.err = dec.unmarshalInt(val)
			f := pv.FieldByName(dec.tagToFieldName(key))
			f.SetInt(num)
		case tok == scanStringLit:
			var str string
			str, dec.err = dec.unmarshalString(val)
			f := pv.FieldByName(dec.tagToFieldName(key))
			f.SetString(str)
		default:
			dec.err = errors.New("unexpected value for project object")
		}

		if dec.err != nil {
			return nil
		}
	}
	return prj
}

func (dec *decoder) decodePackages() []*Package {
	pkgs := make([]*Package, 0)
	for {
		if dec.err != nil {
			return nil
		}

		var tok scanToken
		_, tok, dec.err = dec.scan.nextValue()
		if dec.err != nil {
			return nil
		}
		if tok == scanEndArray {
			break
		}
		if tok != scanBeginObject {
			dec.err = fmt.Errorf("expected an object, found %v", tok)
		}
		pkgs = append(pkgs, dec.decodePackage())
	}

	return nil
}

// TODO: implement
func (dec *decoder) decodePackage() *Package {
	pkg := &Package{}
	pv := reflect.ValueOf(pkg)
	for {
		if dec.err != nil {
			return nil
		}

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

		switch {
		case key == "source_files":
			pkg.SrcFiles = dec.decodeSrcFiles()
		case key == "doc":
			pkg.Doc = dec.decodeStringsList()
		case tok == scanIntLit:
			var num int64
			num, dec.err = dec.unmarshalInt(val)
			f := pv.FieldByName(dec.tagToFieldName(key))
			f.SetInt(num)
		case tok == scanStringLit:
			var str string
			str, dec.err = dec.unmarshalString(val)
			if dec.err != nil {
				return nil
			}
			f := pv.FieldByName(dec.tagToFieldName(key))
			f.SetString(str)
		default:
			dec.err = errors.New("unexpected value for project object")
		}

		if dec.err != nil {
			return nil
		}
	}
	return pkg
}

// TODO: implement
func (dec *decoder) decodeSrcFiles() []*SrcFile {
	return nil
}

func (dec *decoder) decodeStringsList() []string {
	_, tok, err := dec.scan.nextValue()
	if err != nil {
		dec.err = err
		return nil
	}
	if tok != scanBeginArray {
		dec.err = fmt.Errorf("expected array, found %v", tok)
		return nil
	}

	sl := []string{}
	for {
		val, tok, err := dec.scan.nextValue()
		if err != nil {
			dec.err = err
			return nil
		}
		if tok == scanEndArray {
			if len(sl) > 0 {
				dec.err = errors.New("unexpected ']'")
				return nil
			}
			// empty array
			break
		}
		if tok != scanStringLit {
			dec.err = fmt.Errorf("expected string, found %v", tok)
			return nil
		}
		sl = append(sl, string(val))

		val, tok, err = dec.scan.nextValue()
		if err != nil {
			dec.err = err
			return nil
		}
		if tok == scanEndArray {
			// empty array
			break
		}
		if tok != scanComma {
			dec.err = fmt.Errorf("expected ',', found '%s'", val)
			return nil
		}
	}
	return sl
}

// TODO: implement
func (dec *decoder) decodeLanguages() []*Language {
	_, tok, err := dec.scan.nextValue()
	if err != nil {
		dec.err = err
		return nil
	}
	if tok != scanBeginArray {
		dec.err = fmt.Errorf("expected array, found %v", tok)
		return nil
	}

	ls := []*Language{}
	for {
		val, tok, err := dec.scan.nextValue()
		if err != nil {
			dec.err = err
			return nil
		}
		if tok == scanEndArray {
			if len(ls) > 0 {
				dec.err = errors.New("unexpected ']'")
				return nil
			}
			// empty array
			break
		}
		if tok != scanBeginObject {
			dec.err = fmt.Errorf("expected object, found %v", tok)
			return nil
		}

		// Since we have consumed the opening '{', we need to go back
		// before decoding the Language object.
		dec.scan.back()

		lang := dec.decodeLanguage()
		if dec.err != nil {
			return nil
		}

		ls = append(ls, lang)

		val, tok, err = dec.scan.nextValue()
		if err != nil {
			dec.err = err
			return nil
		}
		if tok == scanEndArray {
			// empty array
			break
		}
		if tok != scanComma {
			dec.err = fmt.Errorf("expected ',', found '%s'", val)
			return nil
		}
	}
	return ls
}

// decodeLanguage decode a src.Language object.
func (dec *decoder) decodeLanguage() *Language {
	// Since Language is a JSON Object, we expect to find a '{' character.
	_, tok, err := dec.scan.nextValue()
	if err != nil {
		dec.err = err
		return nil
	}
	if tok != scanBeginObject {
		dec.err = fmt.Errorf("expected object, found %v", tok)
		return nil
	}

	lang := Language{}

	// The object can be empty, so we have to check for that and without
	// consuming the next byte.
	if b, err := dec.scan.peek(); err != nil {
		if err == io.EOF {
			dec.err = errors.New("unexpected EOF")
		} else {
			dec.err = err
		}
		return nil
	} else if b == '}' {
		return &lang
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

		switch {
		case key == "paradigms":
			// Since the '[' character has been consumed, we need to step back
			// brefore calling decodeStringsList.
			dec.scan.back()
			lang.Paradigms = dec.decodeStringsList()
		case key == "language":
			if tok != scanStringLit {
				dec.err = fmt.Errorf("expected string literal, found %v", tok)
				return nil
			}
			if lang.Lang, dec.err = dec.unmarshalString(val); dec.err != nil {
				return nil
			}
		default:
			dec.err = fmt.Errorf("unexpected value for the key '%s' of a language object", key)
		}

		if dec.err != nil {
			return nil
		}

		// Next token can be either a '}' or a ','.
		_, tok, err = dec.scan.nextValue()
		if err != nil {
			dec.err = err
			return nil
		}
		if tok == scanEndObject {
			break
		}
		if tok != scanComma {
			dec.err = fmt.Errorf("expected comma, found %v", tok)
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

// TODO: implement
func (dec *decoder) unmarshalString(data []byte) (string, error) {
	if data == nil {
		return "", errors.New("unable to unmarshal string: data is nil")
	}
	return string(data), nil
}

// TODO: implement
func (dec *decoder) tagToFieldName(tag string) string {
	var field string
	for _, seg := range strings.Split(tag, "_") {
		if len(seg) == 0 {
			return ""
		}
		field += strings.ToUpper(string(seg[0]))
		if len(seg) > 1 {
			field += seg[1:]
		}
	}
	return field
}
