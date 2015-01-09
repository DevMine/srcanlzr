// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"fmt"
	"reflect"
)

type TypeSpec struct {
	Name string `json:"name"`
	Doc  string `json:"doc"`
}

func newTypeSpec(m map[string]interface{}) (*TypeSpec, error) {
	var err error
	errPrefix := "src/type_specifier"
	typespec := TypeSpec{}

	if typespec.Name, err = extractStringValue("name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if typespec.Doc, err = extractStringValue("doc", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &typespec, nil
}

func newTypeSpecsSlice(key, errPrefix string, m map[string]interface{}) ([]*TypeSpec, error) {
	var err error
	var s reflect.Value

	typespecsMap, ok := m[key]
	if !ok {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	}

	if s = reflect.ValueOf(typespecsMap); s.Kind() != reflect.Slice {
		return nil, addDebugInfo(fmt.Errorf(
			"%s: field '%s' is supposed to be a slice", errPrefix, key))
	}

	typespecs := make([]*TypeSpec, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		typespec := s.Index(i).Interface()

		switch typespec.(type) {
		case map[string]interface{}:
			if typespecs[i], err = newTypeSpec(typespec.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}", errPrefix, key))
		}
	}

	return typespecs, nil
}
