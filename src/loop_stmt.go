// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

type LoopStmt struct {
	StmtName   string `json:"statement_name"`
	Init       []Stmt `json:"initialization,omitempty"`
	Cond       Expr   `json:"condition"`
	Post       []Stmt `json:"post_iteration_statement,omitempty"`
	Body       []Stmt `json:"body"`
	Else       []Stmt `json:"else,omitempty"`
	IsPostEval bool   `json:"is_post_evaluated"`
	Line       int64  `json:"line"` // Line number of the statement relatively to the function.
}

// newLoopStmt creates a new LoopStmt from a generic map.
func newLoopStmt(m map[string]interface{}) (*LoopStmt, error) {
	var err error
	errPrefix := "src/loop_stmt"
	loopstmt := LoopStmt{}

	if typ, err := extractStringValue("statement_name", errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	} else if typ != LoopStmtName {
		return nil, fmt.Errorf("invalid type: expected 'LoopStmt', found '%s'", typ)
	}

	loopstmt.StmtName = LoopStmtName

	if loopstmt.Init, err = newStmtsSlice("initialization", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	condMap, err := extractMapValue("condition", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if loopstmt.Cond, err = newExpr(condMap); err != nil {
		return nil, addDebugInfo(err)
	}

	if loopstmt.Post, err = newStmtsSlice("post_iteration_statement", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if loopstmt.IsPostEval, err = extractBoolValue("is_post_evaluated", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if loopstmt.Line, err = extractInt64Value("line", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if loopstmt.Body, err = newStmtsSlice("body", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if loopstmt.Else, err = newStmtsSlice("else", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &loopstmt, nil
}
