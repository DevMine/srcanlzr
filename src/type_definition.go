// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"errors"
	"fmt"
	"reflect"
)

type TypeDef struct {
	Name string   `json:"name"`
	Doc  string   `json:"doc"`
	Type ExprType `json:"type"`
}

func newTypeDef(m map[string]interface{}) (*TypeDef, error) {
	var err error
	errPrefix := "src/type_definition"
	typeDef := TypeDef{}

	if typeDef.Name, err = extractStringValue("name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if typeDef.Doc, err = extractStringValue("doc", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if typeDef.Type, err = newExprType(m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &typeDef, nil
}

func newTypeDefsSlice(key, errPrefix string, m map[string]interface{}) ([]*TypeDef, error) {
	var err error
	var s reflect.Value

	typeDefsMap, ok := m[key]
	if !ok {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	}

	if s = reflect.ValueOf(typeDefsMap); s.Kind() != reflect.Slice {
		return nil, addDebugInfo(errors.New(fmt.Sprintf(
			"%s: field '%s' is supposed to be a slice", errPrefix, key)))
	}

	typeDefs := make([]*TypeDef, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		typeDef := s.Index(i).Interface()

		switch typeDef.(type) {
		case map[string]interface{}:
			if typeDefs[i], err = newTypeDef(typeDef.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(errors.New(fmt.Sprintf(
				"%s: '%s' must be a map[string]interface{}", errPrefix, key)))
		}
	}

	return typeDefs, nil
}
