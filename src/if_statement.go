// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"errors"
	"fmt"
)

type IfStatement struct {
	Type     string      `json:"type"`
	StmtList []Statement `json:"statements_list"`
	Line     int64       `json:"line"` // Line number of the statement relatively to the function.
}

// newIfStatement creates a new IfStatement from a generic map.
func newIfStatement(m map[string]interface{}) (*IfStatement, error) {
	var err error
	errPrefix := "src/if_statement"
	ifstmt := IfStatement{}

	// should never happen
	if typ, ok := m["type"]; !ok || typ != IfStmtName {
		return nil, errors.New(fmt.Sprintf("%s: the generic map supplied is not a IfStatement",
			errPrefix))
	}

	if ifstmt.Type, err = extractStringValue("type", errPrefix, m); err != nil {
		return nil, err
	}

	if ifstmt.Line, err = extractInt64Value("line", errPrefix, m); err != nil {
		return nil, err
	}

	if ifstmt.StmtList, err = newStatementsSlice("statements_list", errPrefix, m); err != nil {
		return nil, err
	}

	return &ifstmt, nil
}
