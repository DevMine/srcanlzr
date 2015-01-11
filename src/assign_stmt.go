// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

type AssignStmt struct {
	StmtName string `json:"statement_name"`
	Lhs      []Expr `json:"left_hand_side"`
	Rhs      []Expr `json:"right_hand_side"`
	Line     int64  `json:"line"`
}

// newAssignStmt creates a new AssignStmt from a generic map.
func newAssignStmt(m map[string]interface{}) (*AssignStmt, error) {
	var err error
	errPrefix := "src/assign_stmt"
	assignstmt := AssignStmt{}

	// should never happen
	if typ, ok := m["statement_name"]; !ok || typ != AssignStmtName {
		return nil, addDebugInfo(fmt.Errorf(
			"%s: the generic map supplied is not a AssignStmt", errPrefix))
	}

	assignstmt.StmtName = AssignStmtName

	if assignstmt.Lhs, err = newExprsSlice("left_hand_side", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if assignstmt.Rhs, err = newExprsSlice("right_hand_side", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if assignstmt.Line, err = extractInt64Value("line", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &assignstmt, nil
}