// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

type AssignStmt struct {
	Type     string `json:"type"`
	VarName  string `json:"var_name"`
	VarValue string `json:"var_value"` // TODO handle case where value is a literal, function call, etc.
	Line     int64  `json:"line"`      // Line number of the statement relatively to the function.
}

// newAssignStmt creates a new AssignStmt from a generic map.
func newAssignStmt(m map[string]interface{}) (*AssignStmt, error) {
	var err error
	errPrefix := "src/assign_stmt"
	assignstmt := AssignStmt{}

	// should never happen
	if typ, ok := m["type"]; !ok || typ != AssignStmtName {
		return nil, addDebugInfo(fmt.Errorf(
			"%s: the generic map supplied is not a AssignStmt", errPrefix))
	}

	if assignstmt.Type, err = extractStringValue("type", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if assignstmt.VarName, err = extractStringValue("var_name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if assignstmt.VarValue, err = extractStringValue("var_value", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if assignstmt.Line, err = extractInt64Value("line", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &assignstmt, nil
}
