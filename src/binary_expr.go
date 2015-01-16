// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

// Binary operators
const (
	// numerical operators
	ADD = "ADD"
	SUB = "SUB"
	MUL = "MUL"
	QUO = "QUO"
	MOD = "MOD"

	// logical operators
	AND     = "AND"         // binary and (&)
	OR      = "OR"          // binary or (|)
	XOR     = "XOR"         // binary xor (^)
	SHL     = "SHIFT_LEFT"  // binary left shift <<
	SHR     = "SHIFT_RIGHT" // binary right shift >>
	AND_NOT = "AND_NOT"     // binary and not (&^)

	// comparison operators
	NEQ  = "NEQ"  // not equal
	LEQ  = "LEQ"  // less or equal
	GEQ  = "GEQ"  // greater or equal
	EQ   = "EQ"   // equal
	LSS  = "LSS"  // less
	GTR  = "GTR"  // greater
	LAND = "LAND" // and (&&)
	LOR  = "LOR"  // or (||)
)

type BinaryExpr struct {
	ExprName  string `json:"expression_name"`
	LeftExpr  Expr   `json:"left_expression,omitempty"`  // left operand
	Op        string `json:"operator"`                   // operator
	RightExpr Expr   `json:"right_expression,omitempty"` // right operand
}

func newBinaryExpr(m map[string]interface{}) (*BinaryExpr, error) {
	var err error
	errPrefix := "src/binary_expr"
	binexpr := BinaryExpr{}

	if typ, err := extractStringValue("expression_name", errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	} else if typ != BinaryExprName {
		return nil, fmt.Errorf("invalid type: expected 'BinaryExpr', found '%s'", typ)
	}

	binexpr.ExprName = BinaryExprName

	exprMap, err := extractMapValue("left_expression", errPrefix, m)
	if err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	} else if err == nil {
		if binexpr.LeftExpr, err = newExpr(exprMap); err != nil {
			return nil, addDebugInfo(err)
		}
	}

	if binexpr.Op, err = extractStringValue("operator", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if exprMap, err = extractMapValue("right_expression", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	} else if err == nil {
		if binexpr.RightExpr, err = newExpr(exprMap); err != nil {
			return nil, addDebugInfo(err)
		}
	}

	return &binexpr, nil
}
