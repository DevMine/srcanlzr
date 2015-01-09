// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"errors"
	"fmt"
	"reflect"
)

// TODO it would be nice to only have the raw function prototype instead of the
// whole function.
// TODO proposal: keep only positions instead of a raw string. This would make
// FIXME add the possibility to have multiple return statements
// the parsing faster and the generated JSON smaller.
type Function struct {
	Name     string      `json:"name"`
	Doc      string      `json:"doc,omitempty"` // TODO rename into doc?
	Args     []*Variable `json:"args,omitempty"`
	Return   []*Variable `json:"return,omitempty"` // TODO put return in statements
	StmtList []Statement `json:"statements_list,omitempty"`
	LoC      int64       `json:"loc"`           // Lines of Code
	Raw      string      `json:"raw,omitempty"` // Function raw source code.
}

func newFunction(m map[string]interface{}) (*Function, error) {
	var err error
	errPrefix := "src/function"
	fct := Function{}

	if fct.Name, err = extractStringValue("name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if fct.Doc, err = extractStringValue("doc", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if fct.LoC, err = extractInt64Value("loc", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if fct.Raw, err = extractStringValue("raw", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if fct.Args, err = newVariablesSlice("args", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if fct.Return, err = newVariablesSlice("returns", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if fct.StmtList, err = newStatementsSlice("statements_list", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	return &fct, nil
}

func newFunctionsSlice(key, errPrefix string, m map[string]interface{}) ([]*Function, error) {
	var err error
	var s reflect.Value

	fctsMap, ok := m[key]
	if !ok {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	}

	if s = reflect.ValueOf(fctsMap); s.Kind() != reflect.Slice {
		return nil, addDebugInfo(errors.New(fmt.Sprintf(
			"%s: field '%s' is supposed to be a slice", errPrefix, key)))
	}

	fcts := make([]*Function, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		fct := s.Index(i).Interface()

		switch fct.(type) {
		case map[string]interface{}:
			if fcts[i], err = newFunction(fct.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(errors.New(fmt.Sprintf(
				"%s: '%s' must be a map[string]interface{}", errPrefix, key)))
		}
	}

	return fcts, nil
}
