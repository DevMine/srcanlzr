// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

type CallStmt struct {
	Type string   `json:"type"`
	Ref  *FuncRef `json:"reference"` // Reference to the function
	Line int64    `json:"line"`      // Line number of the statement relatively to the function.
}

// newCallStmt creates a new CallStmt from a generic map.
func newCallStmt(m map[string]interface{}) (*CallStmt, error) {
	var err error
	errPrefix := "src/call_stmt"
	callstmt := CallStmt{}

	// should never happen
	if typ, ok := m["type"]; !ok || typ != CallStmtName {
		return nil, addDebugInfo(fmt.Errorf(
			"%s: the generic map supplied is not a CallStmt", errPrefix))
	}

	refMap, err := extractMapValue("reference", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if callstmt.Ref, err = newFuncRef(refMap); err != nil {
		return nil, addDebugInfo(err)
	}

	if callstmt.Line, err = extractInt64Value("line", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &callstmt, nil
}
