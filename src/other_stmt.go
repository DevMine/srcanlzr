// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

type OtherStmt struct {
	StmtName  string `json:"statement_name"`
	StmtsList []Stmt `json:"statements_list,omitempty"`
	Line      int64  `json:"line"` // Line number of the statement relatively to the function.
}

// newOtherStmt creates a new OtherStmt from a generic map.
func newOtherStmt(m map[string]interface{}) (*OtherStmt, error) {
	var err error
	errPrefix := "src/other_stmt"
	otherstmt := OtherStmt{}

	// should never happen
	if typ, ok := m["statement_name"]; !ok || typ != OtherStmtName {
		return nil, addDebugInfo(fmt.Errorf(
			"%s: the generic map supplied is not a OtherStmt", errPrefix))
	}

	otherstmt.StmtName = OtherStmtName

	if otherstmt.Line, err = extractInt64Value("line", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if otherstmt.StmtsList, err = newStmtsSlice("statements_list", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	return &otherstmt, nil
}
