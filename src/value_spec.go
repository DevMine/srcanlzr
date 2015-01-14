// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"fmt"
	"reflect"
)

type ValueSpec struct {
	ExprName string `json:"expression_name"`
	Name     *Ident `json:"name"`
	Type     *Ident `json:"type"`
}

func newValueSpec(m map[string]interface{}) (*ValueSpec, error) {
	var err error
	errPrefix := "src/value_spec"
	valspec := ValueSpec{}

	if typ, err := extractStringValue("expression_name", errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	} else if typ != ValueSpecName {
		return nil, fmt.Errorf("invalid type: expected 'ValueSpec', found '%s'", typ)
	}

	valspec.ExprName = ValueSpecName

	nameMap, err := extractMapValue("name", errPrefix, m)
	if err != nil {
		return nil, err
	}

	if valspec.Name, err = newIdent(nameMap); err != nil {
		return nil, addDebugInfo(err)
	}

	typeMap, err := extractMapValue("type", errPrefix, m)
	if err != nil {
		return nil, err
	}

	if valspec.Type, err = newIdent(typeMap); err != nil {
		return nil, addDebugInfo(err)
	}

	return &valspec, nil
}

func newValueSpecsSlice(key, errPrefix string, m map[string]interface{}) ([]*ValueSpec, error) {
	var err error
	var s *reflect.Value

	if s, err = reflectSliceValue(key, errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	valspecs := make([]*ValueSpec, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		valspec := s.Index(i).Interface()
		if valspec == nil {
			continue
		}

		switch valspec.(type) {
		case map[string]interface{}:
			if valspecs[i], err = newValueSpec(valspec.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}, found %v",
				errPrefix, key, reflect.TypeOf(valspec)))
		}
	}

	return valspecs, nil
}
