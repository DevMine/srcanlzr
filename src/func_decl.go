// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"fmt"
	"reflect"
)

type FuncDecl struct {
	Doc        []string  `json:"doc,omitempty"`
	Name       string    `json:"name"`
	Type       *FuncType `json:"type"`
	Body       []Stmt    `json:"body,omitempty"`
	Visibility string    `json:"visibility"`
	LoC        int64     `json:"loc"` // Lines of Code
}

func newFuncDecl(m map[string]interface{}) (*FuncDecl, error) {
	var err error
	errPrefix := "src/func_decl"
	fct := FuncDecl{}

	if fct.Doc, err = extractStringSliceValue("doc", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if fct.Name, err = extractStringValue("name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	fctTypeMap, err := extractMapValue("type", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if fct.Type, err = newFuncType(fctTypeMap); err != nil {
		return nil, addDebugInfo(err)
	}

	if fct.Body, err = newStmtsSlice("body", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if fct.Visibility, err = extractStringValue("visibility", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if fct.LoC, err = extractInt64Value("loc", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &fct, nil
}

func newFuncDeclsSlice(key, errPrefix string, m map[string]interface{}) ([]*FuncDecl, error) {
	var err error
	var s reflect.Value

	fctsMap, ok := m[key]
	if !ok {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	}

	if s = reflect.ValueOf(fctsMap); s.Kind() != reflect.Slice {
		return nil, addDebugInfo(fmt.Errorf(
			"%s: field '%s' is supposed to be a slice", errPrefix, key))
	}

	fcts := make([]*FuncDecl, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		fct := s.Index(i).Interface()

		switch fct.(type) {
		case map[string]interface{}:
			if fcts[i], err = newFuncDecl(fct.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}", errPrefix, key))
		}
	}

	return fcts, nil
}
