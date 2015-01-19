// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"fmt"
	"reflect"
)

// TypeSpec represents a type declaration. Most of the object oriented languages
// does not have such a node, they use classes and traits instead.
//
// In Go, a TypeSpec would be something of the form:
//    type Foo struct {
//       Bar string
//    }
type TypeSpec struct {
	Doc  []string `json:"doc,omitempty"`  // associated documentation; or nil
	Name *Ident   `json:"name"`           // type name (in the exemple, the name is "Foo")
	Type Expr     `json:"type,omitempty"` // *Ident or any of the *XxxType; or nil
}

// newTypeSpec creates a new TypeSpec from a generic map.
func newTypeSpec(m map[string]interface{}) (*TypeSpec, error) {
	var err error
	errPrefix := "src/type_spec"
	typespec := TypeSpec{}

	if typespec.Doc, err = extractStringSliceValue("doc", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	nameMap, err := extractMapValue("name", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if typespec.Name, err = newIdent(nameMap); err != nil {
		return nil, addDebugInfo(err)
	}

	typeMap, err := extractMapValue("type", errPrefix, m)
	if err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	} else if err == nil {
		if typespec.Type, err = newExpr(typeMap); err != nil {
			return nil, addDebugInfo(err)
		}
	}

	return &typespec, nil
}

// newTypeSpecsSlice creates a new slice of TypeSpec from a generic map.
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
