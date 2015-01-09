// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"fmt"
	"reflect"
)

// whole function.
// TODO proposal: keep only positions instead of a raw string. This would make
// the parsing faster and the generated JSON smaller.
type Func struct {
	Name     string `json:"name"`
	Doc      string `json:"doc,omitempty"`
	Args     []*Var `json:"args,omitempty"`
	Return   []*Var `json:"return,omitempty"` // TODO put return in statements
	StmtList []Stmt `json:"statements_list,omitempty"`
	LoC      int64  `json:"loc"`           // Lines of Code
	Raw      string `json:"raw,omitempty"` // Function raw source code.
}

func newFunc(m map[string]interface{}) (*Func, error) {
	var err error
	errPrefix := "src/func"
	fct := Func{}

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

	if fct.Args, err = newVarsSlice("args", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if fct.Return, err = newVarsSlice("returns", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if fct.StmtList, err = newStmtsSlice("statements_list", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	return &fct, nil
}

func newFuncsSlice(key, errPrefix string, m map[string]interface{}) ([]*Func, error) {
	var err error
	var s reflect.Value

	fctsMap, ok := m[key]
	if !ok {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	}

	if s = reflect.ValueOf(fctsMap); s.Kind() != reflect.Slice {
		return nil, addDebugInfo(fmt.Errorf(
			"%s: field '%s' is supposed to be a slice", errPrefix, key))
	}

	fcts := make([]*Func, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		fct := s.Index(i).Interface()

		switch fct.(type) {
		case map[string]interface{}:
			if fcts[i], err = newFunc(fct.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}", errPrefix, key))
		}
	}

	return fcts, nil
}
