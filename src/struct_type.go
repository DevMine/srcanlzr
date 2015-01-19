// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"fmt"
	"reflect"
)

// StructType represents a structured type. Most of the Object Oriented
// languages use a Class or a Trait instead.
//
// In Go, a StructType would be something of the form:
//    struct {
//       Bar string
//    }
type StructType struct {
	// This field is only used by the unmarshaller to "guess" the type while it
	// is unmarshalling a generic type. Since the StructType is considered as
	// an expression (which is represented by an interface{}), this is the only
	// way for the unmarshaller to know what type is it.
	//
	// The value of the ExprName for a StructType must always be "STRUCT", as
	// defined by the constant src.StructTypeName.
	ExprName string `json:"expression_name"`

	Doc    []string `json:"doc"`              // associated documentation; or nil
	Name   *Ident   `json:"name,omitempty"`   // name of the struct; or nil
	Fields []*Field `json:"fields,omitempty"` // the fields of the struct; or nil
}

// Field represents a pair name/type.
type Field struct {
	Doc  []string `json:"doc,omitempty"`  // associated documentation; or nil
	Name string   `json:"name,omitempty"` // name of the field; or nil
	Type string   `json:"type,omitempty"` // type of the field; or nil
}

// newStructType creates a new StructType from a generic map.
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

	nameMap, err := extractMapValue("name", errPrefix, m)
	if err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	} else if err == nil {
		if strct.Name, err = newIdent(nameMap); err != nil {
			return nil, addDebugInfo(err)
		}
	}

	if strct.Fields, err = newFieldsSlice("fields", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	return &strct, nil
}

// newStructTypesSlice creates a new slice of StrucType from a generic map.
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

// newField creates a new Field from a generic map.
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

// newFieldsSlice creates a new slice of fields from a generic map.
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
