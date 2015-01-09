// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

const (
	// logical operators
	AND = "AND"
	OR  = "OR"
	XOR = "XOR"

	// comparison operators
	NEQ = "NEQ" // not equal
	LEQ = "LEQ" // less or equal
	GEQ = "GEQ" // greater or equal
	EQ  = "EQ"  // equal
)

type BinaryExpr struct {
	LeftExpr  Expr   `json:"left_expression"`  // left operand
	Op        string `json:"operator"`         // operator
	RightExpr Expr   `json:"right_expression"` // right operand
}

func newBinaryExpr(m map[string]interface{}) (*BinaryExpr, error) {
	var err error
	errPrefix := "src/binary_expr"
	binexpr := BinaryExpr{}

	// should never happen
	if typ, ok := m["type"]; !ok || typ != BinaryExprName {
		return nil, addDebugInfo(fmt.Errorf(
			"%s: the generic map supplied is not a BinaryExpr", errPrefix))
	}

	exprMap, err := extractMapValue("left_expression", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if binexpr.LeftExpr, err = newExprStmt(exprMap); err != nil {
		return nil, addDebugInfo(err)
	}

	if binexpr.Op, err = extractStringValue("operator", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if exprMap, err = extractMapValue("right_expression", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if binexpr.RightExpr, err = newExpr(exprMap); err != nil {
		return nil, addDebugInfo(err)
	}

	return &binexpr, nil
}
