// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "errors"

type IfStatement struct {
	Type     string      `json:"type"`
	StmtList []Statement `json:"statements_list"`
	Line     int         `json:"line"` // Line number of the statement relatively to the function.
}

// NewIfStatement creates a new IfStatement from a generic map.
func NewIfStatement(m map[string]interface{}) (*IfStatement, error) {
	ifstmt := IfStatement{}

	if typ, ok := m["Type"]; !ok || typ != "IF" {
		return nil, errors.New("the generic map supplied is not a IfStatement")
	}

	ifstmt.Type = m["Type"].(string)

	if line, ok := m["Line"]; ok {
		ifstmt.Line = int(line.(float64))
	}

	if stmts, ok := m["StmtList"]; ok && stmts != nil {
		ifstmt.StmtList = make([]Statement, 0)

		for _, stmt := range m["StmtList"].([]interface{}) {
			castStmt, err := castToStatement(stmt.(map[string]interface{}))
			if err != nil {
				return nil, err
			}

			ifstmt.StmtList = append(ifstmt.StmtList, castStmt)
		}
	}

	return &ifstmt, nil
}
