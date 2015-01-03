// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"errors"
	"fmt"
)

type AssignStatement struct {
	Type     string `json:"type"`
	VarName  string `json:"var_name"`
	VarValue string `json:"var_value"` // TODO handle case where value is a literal, function call, etc.
	Line     int64  `json:"line"`      // Line number of the statement relatively to the function.
}

// newAssignStatement creates a new AssignStatement from a generic map.
func newAssignStatement(m map[string]interface{}) (*AssignStatement, error) {
	var err error
	errPrefix := "src/assign_statement"
	assignstmt := AssignStatement{}

	// should never happen
	if typ, ok := m["type"]; !ok || typ != IfStmtName {
		return nil, errors.New(fmt.Sprintf("%s: the generic map supplied is not a AssignStatement",
			errPrefix))
	}

	if assignstmt.Type, err = extractStringValue("type", errPrefix, m); err != nil {
		return nil, err
	}

	if assignstmt.VarName, err = extractStringValue("var_name", errPrefix, m); err != nil {
		return nil, err
	}

	if assignstmt.VarValue, err = extractStringValue("var_value", errPrefix, m); err != nil {
		return nil, err
	}

	if assignstmt.Line, err = extractInt64Value("line", errPrefix, m); err != nil {
		return nil, err
	}

	return &assignstmt, nil
}
