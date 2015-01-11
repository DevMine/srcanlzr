// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

type IfStmt struct {
	StmtName  string `json:"statement_name"`
	Init      Stmt   `json:"initialization"`
	Cond      Expr   `json:"condition"`
	StmtsList []Stmt `json:"statements_list"`
	Else      []Stmt `json:"else"`
	Line      int64  `json:"line"` // Line number of the statement relatively to the function.
}

// newIfStmt creates a new IfStmt from a generic map.
func newIfStmt(m map[string]interface{}) (*IfStmt, error) {
	var err error
	errPrefix := "src/if_stmt"
	ifstmt := IfStmt{}

	// should never happen
	if typ, ok := m["statement_name"]; !ok || typ != IfStmtName {
		return nil, addDebugInfo(fmt.Errorf(
			"%s: the generic map supplied is not a IfStmt", errPrefix))
	}

	ifstmt.StmtName = IfStmtName

	initMap, err := extractMapValue("initialization", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if ifstmt.Cond, err = newStmt(initMap); err != nil {
		return nil, addDebugInfo(err)
	}

	condMap, err := extractMapValue("condition", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if ifstmt.Cond, err = newExpr(condMap); err != nil {
		return nil, addDebugInfo(err)
	}

	if ifstmt.StmtsList, err = newStmtsSlice("statements_list", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if ifstmt.Else, err = newStmtsSlice("else", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if ifstmt.Line, err = extractInt64Value("line", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &ifstmt, nil
}
