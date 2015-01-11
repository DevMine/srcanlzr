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

	if typ, err := extractStringValue("statement_name", errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	} else if typ != OtherStmtName {
		return nil, fmt.Errorf("invalid type: expected 'OtherStmt', found '%s'", typ)
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
