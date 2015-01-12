// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

// Unary operators
const (
	NOT  = "NOT"
	ADDR = "ADDR" // memory address (&foo)
	STAR = "STAR" // dereference operator (*foo)
)

type UnaryExpr struct {
	ExprName string `json:"expression_name"`
	Op       string `json:"operator"` // operator
	X        Expr   `json:"operand"`  // operand
}

func newUnaryExpr(m map[string]interface{}) (*UnaryExpr, error) {
	var err error
	errPrefix := "src/unary_expr"
	unaryexpr := UnaryExpr{}

	if typ, err := extractStringValue("expression_name", errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	} else if typ != UnaryExprName {
		return nil, fmt.Errorf("invalid type: expected 'UnaryExpr', found '%s'", typ)
	}

	unaryexpr.ExprName = UnaryExprName

	if unaryexpr.Op, err = extractStringValue("operator", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	exprMap, err := extractMapValue("operand", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if unaryexpr.X, err = newExpr(exprMap); err != nil {
		return nil, addDebugInfo(err)
	}

	return &unaryexpr, nil
}
