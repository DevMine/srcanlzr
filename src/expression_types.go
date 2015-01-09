// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"errors"
	"fmt"
	"reflect"
)

const (
	PrimitiveTypeName = "PRIMITIVE"
	StructTypeName    = "STRUCT"
)

// ExprType can be either a structured type (Struct) or another type (PrimitiveType)
// such as an int, float, etc.
type ExprType interface{}

type PrimitiveType struct {
	Type string `json:"type"`
	Name string `json:"name"` // int, float, string, etc.
}

type StructuredType struct {
	Type   string  `json:"type"`
	Name   string  `json:"name"`
	Doc    string  `json:"doc"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Name string `json:"name"`

	// TODO rename into TypeName or use a type ExprType
	Type string `json:"type"`

	Doc string `json:"doc"`
}

func newExprType(m map[string]interface{}) (ExprType, error) {
	errPrefix := "src/expression_type"

	typ, ok := m["type"]
	if !ok {
		return nil, addDebugInfo(errors.New(fmt.Sprintf(
			"%s: field 'type' does not exist", errPrefix)))
	}

	switch typ {
	case PrimitiveTypeName:
		return newPrimitive(m)
	case StructTypeName:
		return newStruct(m)
	}

	return nil, addDebugInfo(errors.New("unknown expression type"))
}

func newPrimitive(m map[string]interface{}) (*PrimitiveType, error) {
	var err error
	errPrefix := "src/structured_type"
	prim := PrimitiveType{}

	// should never happen
	if typ, ok := m["type"]; !ok || typ != PrimitiveTypeName {
		return nil, addDebugInfo(errors.New(fmt.Sprintf(
			"%s: the generic map supplied is not a PrimitiveType",
			errPrefix)))
	}

	if prim.Type, err = extractStringValue("type", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if prim.Name, err = extractStringValue("name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &prim, nil
}

func newStruct(m map[string]interface{}) (*StructuredType, error) {
	var err error
	errPrefix := "src/structured_type"
	strct := StructuredType{}

	// should never happen
	if typ, ok := m["type"]; !ok || typ != StructTypeName {
		return nil, addDebugInfo(errors.New(fmt.Sprintf(
			"%s: the generic map supplied is not a PrimitiveType",
			errPrefix)))
	}

	if strct.Type, err = extractStringValue("type", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if strct.Name, err = extractStringValue("name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if strct.Doc, err = extractStringValue("doc", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &strct, nil
}

func newStructsSlice(key, errPrefix string, m map[string]interface{}) ([]*StructuredType, error) {
	var err error
	var s reflect.Value

	structsMap, ok := m[key]
	if !ok {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	}

	if s = reflect.ValueOf(structsMap); s.Kind() != reflect.Slice {
		return nil, addDebugInfo(errors.New(fmt.Sprintf(
			"%s: field '%s' is supposed to be a slice", errPrefix, key)))
	}

	structs := make([]*StructuredType, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		strct := s.Index(i).Interface()

		switch strct.(type) {
		case map[string]interface{}:
			if structs[i], err = newStruct(strct.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(errors.New(fmt.Sprintf(
				"%s: '%s' must be a map[string]interface{}", errPrefix, key)))
		}
	}

	return structs, nil
}

func newField(m map[string]interface{}) (*Field, error) {
	var err error
	errPrefix := "src/field"
	field := Field{}

	if field.Name, err = extractStringValue("name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if field.Type, err = extractStringValue("type", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if field.Doc, err = extractStringValue("doc", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &field, nil
}

func newFieldsSlice(key, errPrefix string, m map[string]interface{}) ([]*Field, error) {
	var err error
	var s reflect.Value

	fieldsMap, ok := m[key]
	if !ok {
		return nil, addDebugInfo(errors.New(fmt.Sprintf(
			"%s: field '%s' does not exist", errPrefix, key)))
	}

	if s = reflect.ValueOf(fieldsMap); s.Kind() != reflect.Slice {
		return nil, addDebugInfo(errors.New(fmt.Sprintf(
			"%s: field '%s' is supposed to be a slice", errPrefix, key)))
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
			return nil, addDebugInfo(errors.New(fmt.Sprintf(
				"%s: '%s' must be a map[string]interface{}", errPrefix, key)))
		}
	}

	return fields, nil
}
