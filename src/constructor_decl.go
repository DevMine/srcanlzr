// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"fmt"
	"reflect"
)

type ConstructorDecl struct {
	Doc        string   `json:"doc,omitempty"`
	Name       string   `json:"name"`
	Params     []*Field `json:"parameters,omitempty"`
	Body       []Stmt   `json:"body,omitempty"`
	Visibility string   `json:"visibility"`
	LoC        int64    `json:"loc"`
}

func newConstructorDecl(m map[string]interface{}) (*ConstructorDecl, error) {
	var err error
	errPrefix := "src/constructor_decl"
	construc := ConstructorDecl{}

	if construc.Doc, err = extractStringValue("doc", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if construc.Name, err = extractStringValue("name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if construc.Params, err = newFieldsSlice("parameters", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if construc.Body, err = newStmtsSlice("body", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if construc.Visibility, err = extractStringValue("visibility", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if construc.LoC, err = extractInt64Value("loc", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &construc, nil
}

func newConstructorDeclsSlice(key, errPrefix string, m map[string]interface{}) ([]*ConstructorDecl, error) {
	var err error
	var s *reflect.Value

	if s, err = reflectSliceValue(key, errPrefix, m); err != nil {
		return nil, err
	}

	construcs := make([]*ConstructorDecl, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		construc := s.Index(i).Interface()

		switch construc.(type) {
		case map[string]interface{}:
			if construcs[i], err = newConstructorDecl(construc.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}", errPrefix, key))
		}
	}

	return construcs, nil
}
