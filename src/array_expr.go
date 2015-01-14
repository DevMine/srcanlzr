// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

type ArrayExpr struct {
	ExprName string     `json:"expression_name"`
	Type     *ArrayType `json:"type"`
}

func newArrayExpr(m map[string]interface{}) (*ArrayExpr, error) {
	var err error
	errPrefix := "src/array_expr"
	aryexpr := ArrayExpr{}

	if typ, err := extractStringValue("expression_name", errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	} else if typ != ArrayExprName {
		return nil, fmt.Errorf("invalid type: expected 'ArrayExpr', found '%s'", typ)
	}

	aryexpr.ExprName = ArrayExprName

	typeMap, err := extractMapValue("type", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if aryexpr.Type, err = newArrayType(typeMap); err != nil {
		return nil, addDebugInfo(err)
	}

	return &aryexpr, nil
}
