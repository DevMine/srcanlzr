// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

type ArrayLit struct {
	ExprName string     `json:"expression_name"`
	Type     *ArrayType `json:"type"`
	Elts     []Expr     `json:"elements"`
}

func newArrayLit(m map[string]interface{}) (*ArrayLit, error) {
	var err error
	errPrefix := "src/array_lit"
	arylit := ArrayLit{}

	if typ, err := extractStringValue("expression_name", errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	} else if typ != ArrayLitName {
		return nil, fmt.Errorf("invalid type: expected 'ArrayLit', found '%s'", typ)
	}

	arylit.ExprName = ArrayLitName

	typeMap, err := extractMapValue("type", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if arylit.Type, err = newArrayType(typeMap); err != nil {
		return nil, addDebugInfo(err)
	}

	if arylit.Elts, err = newExprsSlice("elements", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &arylit, nil
}