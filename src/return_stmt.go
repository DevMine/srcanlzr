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

	if typ, err := extractStringValue("statement_name", errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	} else if typ != ReturnStmtName {
		return nil, fmt.Errorf("invalid type: expected 'ReturnStmt', found '%s'", typ)
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
