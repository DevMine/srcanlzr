// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

type DeclStmt struct {
	AssignStmt
	Kind string `json:"kind"`
}

// newDeclStmt creates a new DeclStmt from a generic map.
func newDeclStmt(m map[string]interface{}) (*DeclStmt, error) {
	var err error
	errPrefix := "src/decl_stmt"
	declstmt := DeclStmt{}

	// should never happen
	if typ, ok := m["type"]; !ok || typ != DeclStmtName {
		return nil, addDebugInfo(fmt.Errorf(
			"%s: the generic map supplied is not a DeclStmt", errPrefix))
	}

	if declstmt.Type, err = extractStringValue("type", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if declstmt.Lhs, err = newExprsSlice("left_hand_side", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if declstmt.Rhs, err = newExprsSlice("right_hand_side", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if declstmt.Line, err = extractInt64Value("line", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if declstmt.Kind, err = extractStringValue("kind", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &declstmt, nil
}
