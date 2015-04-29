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

		var tok int
		_, tok, dec.err = dec.scan.nextValue()
		if dec.err != nil {
			return nil
		}
		if tok == scanEndArray {
			break
		}
		if tok != scanBeginObject {
			dec.err = errors.New("expected an object")
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

// TODO: implement
func (dec *decoder) decodeStringsList() []string {
	return nil
}

// TODO: implement
func (dec *decoder) decodeLanguages() []*Language {
	return nil
}

// TODO: implement
func (dec *decoder) decodeLanguage() *Language {
	return nil
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
	return "", nil
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
