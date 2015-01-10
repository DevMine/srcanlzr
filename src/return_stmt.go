// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

// A ReturnStmt represents a return statement.
type ReturnStmt struct {
	Type    string `json:"type"`    // TODO rename this field into StmtName
	Results []Expr `json:"results"` // result expressions; or nil
	Line    int64  `json:"line"`
}

func newReturnStmt(m map[string]interface{}) (*ReturnStmt, error) {
	var err error
	errPrefix := "src/return_stmt"
	retstmt := ReturnStmt{}

	// should never happen
	if typ, ok := m["type"]; !ok || typ != ReturnStmtName {
		return nil, addDebugInfo(fmt.Errorf(
			"%s: the generic map supplied is not a ReturnStmt", errPrefix))
	}

	if retstmt.Type, err = extractStringValue("type", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if retstmt.Results, err = newExprsSlice("results", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if retstmt.Line, err = extractInt64Value("line", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &retstmt, nil
}
