// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

type TernaryExpr struct {
	ExprName string `json:"expression_name"`
	Cond     Expr   `json:"condition"`
	Then     Expr   `json:"then"`
	Else     Expr   `json:"else"`
}

func newTernaryExpr(m map[string]interface{}) (*TernaryExpr, error) {
	var err error
	errPrefix := "src/ternary_expr"
	ternexpr := TernaryExpr{}

	if typ, err := extractStringValue("expression_name", errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	} else if typ != TernaryExprName {
		return nil, fmt.Errorf("invalid type: expected 'TernaryExpr', found '%s'", typ)
	}

	exprMap, err := extractMapValue("condition", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if ternexpr.Cond, err = newExprStmt(exprMap); err != nil {
		return nil, addDebugInfo(err)
	}

	if exprMap, err = extractMapValue("then", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if ternexpr.Then, err = newExprStmt(exprMap); err != nil {
		return nil, addDebugInfo(err)
	}

	if exprMap, err = extractMapValue("else", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if ternexpr.Then, err = newExprStmt(exprMap); err != nil {
		return nil, addDebugInfo(err)
	}

	return &ternexpr, nil
}
