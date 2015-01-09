// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "errors"

// TODO rename this type into LoopStatement and the loop stop condition
type LoopStatement struct {
	Type     string      `json:"type"`
	StmtList []Statement `json:"statements_list"`
	Line     int64       `json:"line"` // Line number of the statement relatively to the function.
}

// newLoopStatement creates a new LoopStatement from a generic map.
func newLoopStatement(m map[string]interface{}) (*LoopStatement, error) {

	var err error
	errPrefix := "src/loop_statement"
	loopstmt := LoopStatement{}

	// should never happen
	if typ, ok := m["type"]; !ok || typ != LoopStmtName {
		return nil, addDebugInfo(errors.New(
			"src/loop_statement: the generic map supplied is not a LoopStatement"))
	}

	if loopstmt.Type, err = extractStringValue("type", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if loopstmt.Line, err = extractInt64Value("line", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if loopstmt.StmtList, err = newStatementsSlice("statements_list", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &loopstmt, nil
}
