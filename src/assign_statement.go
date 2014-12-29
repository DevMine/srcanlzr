// Copyright 2014 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "errors"

type AssignStatement struct {
	Type     string `json:"type"`
	VarName  string `json:"var_name"`
	VarValue string `json:"var_value"` // TODO handle case where value is a literal, function call, etc.
	Line     int    `json:"line"`      // Line number of the statement relatively to the function.
}

// NewAssignStatement creates a new AssignStatement from a generic map.
func NewAssignStatement(m map[string]interface{}) (*AssignStatement, error) {
	assignstmt := AssignStatement{}

	if typ, ok := m["Type"]; !ok || typ != "ASSIGN" {
		return nil, errors.New("the generic map supplied is not a AssignStatement")
	}

	assignstmt.Type = m["Type"].(string)

	if line, ok := m["Line"]; ok {
		// XXX unsafe cast
		assignstmt.Line = int(line.(float64))
	}

	if varName, ok := m["VarName"]; ok {
		assignstmt.VarName = varName.(string)
	}

	if varValue, ok := m["VarValue"]; ok {
		assignstmt.VarValue = varValue.(string)
	}

	return &assignstmt, nil
}
