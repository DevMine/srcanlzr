// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"errors"
	"fmt"
	"reflect"
)

type Constant struct {
	Name  string `json:"name"`
	Type  string `json:"type"`  // TODO rename into TypeName or use a type Type
	Value string `json:"value"` // TODO use an Expr instead of string value
	Doc   string `json:"doc"`
}

func newConstant(m map[string]interface{}) (*Constant, error) {
	var err error
	errPrefix := "src/constant"
	cst := Constant{}

	if cst.Name, err = extractStringValue("name", errPrefix, m); err != nil {
		return nil, err
	}

	if cst.Type, err = extractStringValue("type", errPrefix, m); err != nil {
		return nil, err
	}

	if cst.Value, err = extractStringValue("value", errPrefix, m); err != nil {
		return nil, err
	}

	if cst.Doc, err = extractStringValue("doc", errPrefix, m); err != nil {
		return nil, err
	}

	return &cst, nil
}

func newConstantsSlice(key, errPrefix string, m map[string]interface{}) ([]*Constant, error) {
	var err error
	var s reflect.Value

	cstsMap, ok := m[key]
	if !ok {
		return nil, errNotExist
	}

	if s = reflect.ValueOf(cstsMap); s.Kind() != reflect.Slice {
		return nil, errors.New(fmt.Sprintf("%s: field '%s' is supposed to be a slice",
			errPrefix, key))
	}

	csts := make([]*Constant, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		cst := s.Index(i).Interface()

		switch cst.(type) {
		case map[string]interface{}:
			if csts[i], err = newConstant(cst.(map[string]interface{})); err != nil {
				return nil, err
			}
		default:
			return nil, errors.New(fmt.Sprintf("%s: '%s' must be a map[string]interface{}",
				errPrefix, key))
		}
	}

	return csts, nil
}
