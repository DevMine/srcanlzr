// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"fmt"
	"reflect"
)

type StructType struct {
	ExprName string   `json:"expression_name"`
	Doc      []string `json:"doc"`
	Name     string   `json:"name"`
	Type     Expr     `json:"type"`
	Fields   []Field  `json:"fields"`
}

type Field struct {
	Doc  []string `json:"doc,omitempty"`
	Name string   `json:"name,omitempty"`
	Type string   `json:"type,omitempty"`
}

func newStructType(m map[string]interface{}) (*StructType, error) {
	var err error
	errPrefix := "src/structured_type"
	strct := StructType{}

	if typ, err := extractStringValue("expression_name", errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	} else if typ != StructTypeName {
		return nil, fmt.Errorf("invalid type: expected 'StructTypeName', found '%s'", typ)
	}

	strct.ExprName = StructTypeName

	if strct.Doc, err = extractStringSliceValue("doc", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if strct.Type, err = extractStringValue("type", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if strct.Name, err = extractStringValue("name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &strct, nil
}

func newStructTypesSlice(key, errPrefix string, m map[string]interface{}) ([]*StructType, error) {
	var err error
	var s reflect.Value

	structsMap, ok := m[key]
	if !ok {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	}

	if s = reflect.ValueOf(structsMap); s.Kind() != reflect.Slice {
		return nil, addDebugInfo(fmt.Errorf(
			"%s: field '%s' is supposed to be a slice", errPrefix, key))
	}

	structs := make([]*StructType, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		strct := s.Index(i).Interface()

		switch strct.(type) {
		case map[string]interface{}:
			if structs[i], err = newStructType(strct.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}", errPrefix, key))
		}
	}

	return structs, nil
}

func newField(m map[string]interface{}) (*Field, error) {
	var err error
	errPrefix := "src/field"
	field := Field{}

	if field.Doc, err = extractStringSliceValue("doc", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if field.Name, err = extractStringValue("name", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if field.Type, err = extractStringValue("type", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	return &field, nil
}

func newFieldsSlice(key, errPrefix string, m map[string]interface{}) ([]*Field, error) {
	var err error
	var s reflect.Value

	fieldsMap, ok := m[key]
	if !ok {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	}

	if s = reflect.ValueOf(fieldsMap); s.Kind() != reflect.Slice {
		return nil, addDebugInfo(fmt.Errorf(
			"%s: field '%s' is supposed to be a slice", errPrefix, key))
	}

	fields := make([]*Field, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		field := s.Index(i).Interface()

		switch field.(type) {
		case map[string]interface{}:
			if fields[i], err = newField(field.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}", errPrefix, key))
		}
	}

	return fields, nil
}
