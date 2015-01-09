// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "errors"

type LoopStmt struct {
	Type      string `json:"type"`
	StmtsList []Stmt `json:"statements_list"`
	Line      int64  `json:"line"` // Line number of the statement relatively to the function.
}

// newLoopStmt creates a new LoopStmt from a generic map.
func newLoopStmt(m map[string]interface{}) (*LoopStmt, error) {
	var err error
	errPrefix := "src/loop_stmt"
	loopstmt := LoopStmt{}

	// should never happen
	if typ, ok := m["type"]; !ok || typ != LoopStmtName {
		// FIXME use errPrefix
		return nil, addDebugInfo(errors.New(
			"src/loop_statement: the generic map supplied is not a LoopStmt"))
	}

	if loopstmt.Type, err = extractStringValue("type", errPrefix, m); err != nil {
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
