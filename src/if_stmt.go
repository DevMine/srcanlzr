// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

type IfStmt struct {
	StmtName string `json:"statement_name"`
	Init     Stmt   `json:"initialization"`
	Cond     Expr   `json:"condition"`
	Body     []Stmt `json:"body"`
	Else     []Stmt `json:"else"`
	Line     int64  `json:"line"` // Line number of the statement relatively to the function.
}

// newIfStmt creates a new IfStmt from a generic map.
func newIfStmt(m map[string]interface{}) (*IfStmt, error) {
	var err error
	errPrefix := "src/if_stmt"
	ifstmt := IfStmt{}

	if typ, err := extractStringValue("statement_name", errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	} else if typ != IfStmtName {
		return nil, fmt.Errorf("invalid type: expected 'IfStmt', found '%s'", typ)
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

	if ifstmt.Body, err = newStmtsSlice("body", errPrefix, m); err != nil {
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
