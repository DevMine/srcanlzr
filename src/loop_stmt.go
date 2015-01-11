// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

type LoopStmt struct {
	StmtName   string `json:"statement_name"`
	Init       Stmt   `json:"initialization"`
	Cond       Expr   `json:"condition"`
	Post       Stmt   `json:"post_iteration_statement"`
	StmtsList  []Stmt `json:"statements_list"`
	IsPostEval bool   `json:"is_post_evaluated"`
	Line       int64  `json:"line"` // Line number of the statement relatively to the function.
}

// newLoopStmt creates a new LoopStmt from a generic map.
func newLoopStmt(m map[string]interface{}) (*LoopStmt, error) {
	var err error
	errPrefix := "src/loop_stmt"
	loopstmt := LoopStmt{}

	// should never happen
	if typ, ok := m["statement_name"]; !ok || typ != LoopStmtName {
		return nil, addDebugInfo(fmt.Errorf(
			"%s: the generic map supplied is not a LoopStmt", errPrefix))
	}

	loopstmt.StmtName = LoopStmtName

	initMap, err := extractMapValue("initialization", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if loopstmt.Init, err = newStmt(initMap); err != nil {
		return nil, addDebugInfo(err)
	}

	condMap, err := extractMapValue("condition", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if loopstmt.Cond, err = newExpr(condMap); err != nil {
		return nil, addDebugInfo(err)
	}

	postMap, err := extractMapValue("post_iteration_statement", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if loopstmt.Post, err = newStmt(postMap); err != nil {
		return nil, addDebugInfo(err)
	}

	if loopstmt.IsPostEval, err = extractBoolValue("is_post_evaluated", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if loopstmt.Line, err = extractInt64Value("line", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if loopstmt.StmtsList, err = newStmtsSlice("statements_list", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &loopstmt, nil
}