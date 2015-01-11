// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

type FuncLit struct {
	Type    string   `json:"type"`
	Name    string   `json:"name"`
	Params  []*Field `json:"parameters,omitempty"`
	Results []*Field `json:"results,omitempty"`
	Body    []Stmt   `json:"body,omitempty"`
	LoC     int64    `json:"loc"` // Lines of Code
}

func newFuncLit(m map[string]interface{}) (*FuncLit, error) {
	var err error
	errPrefix := "src/func_lit"
	fct := FuncLit{}

	if typ, err := extractStringValue("type", errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	} else if typ != FuncLitName {
		return nil, fmt.Errorf("invalid type: expected 'FuncLit', found '%s'", typ)
	}

	fct.Type = FuncLitName

	if fct.Name, err = extractStringValue("name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if fct.Params, err = newFieldsSlice("parameters", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if fct.Results, err = newFieldsSlice("results", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if fct.Body, err = newStmtsSlice("body", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if fct.LoC, err = extractInt64Value("loc", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &fct, nil
}
