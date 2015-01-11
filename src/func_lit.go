// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

type FuncLit struct {
	ExprName string    `json:"expression_name"`
	Name     string    `json:"name"`
	Type     *FuncType `json:"type"`
	Body     []Stmt    `json:"body,omitempty"`
	LoC      int64     `json:"loc"` // Lines of Code
}

func newFuncLit(m map[string]interface{}) (*FuncLit, error) {
	var err error
	errPrefix := "src/func_lit"
	fct := FuncLit{}

	if typ, err := extractStringValue("expression_name", errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	} else if typ != FuncLitName {
		return nil, fmt.Errorf("invalid type: expected 'FuncLit', found '%s'", typ)
	}

	fct.ExprName = FuncLitName

	if fct.Name, err = extractStringValue("name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	fctTypeMap, err := extractMapValue("type", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if fct.Type, err = newFuncType(fctTypeMap); err != nil {
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
