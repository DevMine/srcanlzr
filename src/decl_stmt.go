// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

// Kind of declarations
const (
	ConstDecl = "CONSTANT" // constant
	VarDecl   = "VAR"      // variable
)

type DeclStmt struct {
	AssignStmt
	Kind string `json:"kind"`
}

// newDeclStmt creates a new DeclStmt from a generic map.
func newDeclStmt(m map[string]interface{}) (*DeclStmt, error) {
	var err error
	errPrefix := "src/decl_stmt"
	declstmt := DeclStmt{}

	if typ, err := extractStringValue("statement_name", errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	} else if typ != DeclStmtName {
		return nil, fmt.Errorf("invalid type: expected 'DeclStmt', found '%s'", typ)
	}

	declstmt.StmtName = DeclStmtName

	if declstmt.LHS, err = newExprsSlice("left_hand_side", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if declstmt.RHS, err = newExprsSlice("right_hand_side", errPrefix, m); err != nil {
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
