// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

type ExprStmt struct {
	StmtName string `json:"statement_name"`
	X        Expr   `json:"expression"` // expression
}

func newExprStmt(m map[string]interface{}) (*ExprStmt, error) {
	var err error
	errPrefix := "src/expr_stmt"
	exprstmt := ExprStmt{}

	// should never happen
	if typ, ok := m["statement_name"]; !ok || typ != ExprStmtName {
		return nil, addDebugInfo(fmt.Errorf(
			"%s: the generic map supplied is not a ExprStmt", errPrefix))
	}

	exprstmt.StmtName = ExprStmtName

	exprMap, err := extractMapValue("expression", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if exprstmt.X, err = newExpr(exprMap); err != nil {
		return nil, addDebugInfo(err)
	}

	return &exprstmt, nil
}
