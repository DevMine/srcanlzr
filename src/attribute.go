// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"fmt"
	"reflect"
)

type Attribute struct {
	Variable
	Visibility string `json:"visibility"`
	Constant   bool   `json:"constant"`
	Static     bool   `json:"static"`
}

func newAttribute(m map[string]interface{}) (*Attribute, error) {
	var err error
	errPrefix := "src/attribute"
	attr := Attribute{}

	var v *Variable
	if v, err = newVariable(m); err != nil {
		return nil, addDebugInfo(err)
	}
	attr.Variable = *v

	if attr.Visibility, err = extractStringValue("visibility", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if attr.Constant, err = extractBoolValue("constant", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if attr.Static, err = extractBoolValue("static", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &attr, nil
}

func newAttributesSlice(key, errPrefix string, m map[string]interface{}) ([]*Attribute, error) {
	var err error
	var s reflect.Value

	attrsMap, ok := m[key]
	if !ok {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	}

	if s = reflect.ValueOf(attrsMap); s.Kind() != reflect.Slice {
		return nil, addDebugInfo(fmt.Errorf(
			"%s: field '%s' is supposed to be a slice", errPrefix, key))
	}

	attrs := make([]*Attribute, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		attr := s.Index(i).Interface()

		switch attr.(type) {
		case map[string]interface{}:
			if attrs[i], err = newAttribute(attr.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}", errPrefix, key))
		}
	}

	return attrs, nil
}
