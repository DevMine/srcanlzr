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
	prj  *Project
}

func newDecoder(r io.Reader) *decoder {
	return &decoder{scan: newScanner(r), prj: &Project{}}
}

// TODO: implement
func (dec *decoder) decode() (*Project, error) {
	if _, tok, err := dec.scan.nextValue(); err != nil {
		return nil, dec.errorf(err)
	} else if tok != scanBeginObject {
		return nil, dec.errorf("expected '{' as first character")
	}

	if err := dec.decodeProject(); err != nil {
		return nil, dec.errorf(err)
	}

	if _, tok, err := dec.scan.nextValue(); err != nil {
		return nil, dec.errorf(err)
	} else if tok != scanEndObject {
		return nil, dec.errorf("expected '}' as last character")
	}

	return dec.prj, nil
}

func (dec *decoder) errorf(v interface{}) error {
	return fmt.Errorf("malformed json: %v", v)
}

func (dec *decoder) decodeProject() error {
	pv := reflect.ValueOf(dec.prj)

	for {
		key, err := dec.scan.nextKey()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if key == "" {
			return errors.New("empty key")
		}

		val, tok, err := dec.scan.nextValue()
		if err != nil {
			return err
		}

		switch {
		case key == "packages":
			err = dec.decodePackages()
		case key == "languages":
			err = dec.decodeLanguages()
		case key == "repository":
			err = dec.decodeRepository()
		case tok == scanIntLit:
			var num int64
			num, err := dec.unmarshalInt(val)
			if err != nil {
				return err
			}
			f := pv.FieldByName(dec.tagToFieldName(key))
			f.SetInt(num)
		case tok == scanStringLit:
			var str string
			str, err := dec.unmarshalString(val)
			if err != nil {
				return err
			}
			f := pv.FieldByName(dec.tagToFieldName(key))
			f.SetString(str)
		default:
			return errors.New("unexpected value for project object")
		}

		if err != nil {
			return err
		}
	}
	return nil
}

// TODO: implement
func (dec *decoder) decodePackages() error {
	return nil
}

// TODO: implement
func (dec *decoder) decodePackage() error {
	return nil
}

// TODO: implement
func (dec *decoder) decodeLanguages() error {
	return nil
}

// TODO: implement
func (dec *decoder) decodeLanguage() error {
	return nil
}

// TODO: implement
func (dec *decoder) decodeRepository() error {
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
