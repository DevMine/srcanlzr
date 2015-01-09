// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"errors"
	"fmt"
	"reflect"
)

type Variable struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
	Doc   string `json:"doc"`
}

func newVariable(m map[string]interface{}) (*Variable, error) {
	var err error
	errPrefix := "src/variable"
	v := Variable{}

	if v.Name, err = extractStringValue("name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if v.Type, err = extractStringValue("type", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if v.Value, err = extractStringValue("value", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if v.Doc, err = extractStringValue("doc", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &v, nil
}

func newVariablesSlice(key, errPrefix string, m map[string]interface{}) ([]*Variable, error) {
	var err error
	var s reflect.Value

	varsMap, ok := m[key]
	if !ok {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	}

	if s = reflect.ValueOf(varsMap); s.Kind() != reflect.Slice {
		return nil, addDebugInfo(errors.New(fmt.Sprintf(
			"%s: field '%s' is supposed to be a slice", errPrefix, key)))
	}

	vars := make([]*Variable, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		v := s.Index(i).Interface()

		switch v.(type) {
		case map[string]interface{}:
			if vars[i], err = newVariable(v.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(errors.New(fmt.Sprintf(
				"%s: '%s' must be a map[string]interface{}", errPrefix, key)))
		}
	}

	return vars, nil
}
