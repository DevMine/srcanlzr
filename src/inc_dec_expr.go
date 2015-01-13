// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

// Increment/Decrement operators
const (
	INC = "INC"
	DEC = "DEC"
)

type IncDecExpr struct {
	ExprName string `json:"expression_name"`
	X        Expr   `json:"operand"`
	Op       string `json:"operator"` // INC or DEC
}

func newIncDecExpr(m map[string]interface{}) (*IncDecExpr, error) {
	var err error
	errPrefix := "src/inc_dec_expr"
	incdecexpr := IncDecExpr{}

	if typ, err := extractStringValue("expression_name", errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	} else if typ != IncDecExprName {
		return nil, fmt.Errorf("invalid type: expected 'IncDecExpr', found '%s'", typ)
	}

	incdecexpr.ExprName = IncDecExprName

	exprMap, err := extractMapValue("operand", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if incdecexpr.X, err = newExpr(exprMap); err != nil {
		return nil, addDebugInfo(err)
	}

	if incdecexpr.Op, err = extractStringValue("operator", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &incdecexpr, nil
}
