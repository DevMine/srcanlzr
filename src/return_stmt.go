// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

// A ReturnStmt represents a return statement.
type ReturnStmt struct {
	StmtName string `json:"statement_name"`
	Results  []Expr `json:"results"` // result expressions; or nil
	Line     int64  `json:"line"`
}

func newReturnStmt(m map[string]interface{}) (*ReturnStmt, error) {
	var err error
	errPrefix := "src/return_stmt"
	retstmt := ReturnStmt{}

	// should never happen
	if typ, ok := m["statement_name"]; !ok || typ != ReturnStmtName {
		return nil, addDebugInfo(fmt.Errorf(
			"%s: the generic map supplied is not a ReturnStmt", errPrefix))
	}

	retstmt.StmtName = ReturnStmtName

	if retstmt.Results, err = newExprsSlice("results", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if retstmt.Line, err = extractInt64Value("line", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &retstmt, nil
}
