// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

type RangeLoopStmt struct {
	StmtName string       `json:"statement_name"`
	Vars     []*ValueSpec `json:"variables,omitempty"`
	Iterable Expr         `json:"iterable"`
	Body     []Stmt       `json:"body"`
	Line     int64        `json:"line"` // Line number of the statement relatively to the function.
}

// newRangeLoopStmt creates a new RangeLoopStmt from a generic map.
func newRangeLoopStmt(m map[string]interface{}) (*RangeLoopStmt, error) {
	var err error
	errPrefix := "src/range_loop_stmt"
	loopstmt := RangeLoopStmt{}

	if typ, err := extractStringValue("statement_name", errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	} else if typ != RangeLoopStmtName {
		return nil, fmt.Errorf("invalid type: expected 'RangeLoopStmt', found '%s'", typ)
	}

	loopstmt.StmtName = RangeLoopStmtName

	if loopstmt.Vars, err = newValueSpecsSlice("variables", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	iterMap, err := extractMapValue("iterable", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if loopstmt.Iterable, err = newExpr(iterMap); err != nil {
		return nil, addDebugInfo(err)
	}

	if loopstmt.Body, err = newStmtsSlice("body", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &loopstmt, nil
}
