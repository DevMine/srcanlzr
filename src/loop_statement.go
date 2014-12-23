// Copyright 2014 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "errors"

// TODO rename this type into LoopStatement and the loop stop condition
type LoopStatement struct {
	Type     string      `json:"type"`
	StmtList []Statement `json:"statements_list"`
	Line     int         `json:"line"` // Line number of the statement relatively to the function.
}

// CastToLoopStatement "cast" a generic map into a LoopStatement.
func CastToLoopStatement(m map[string]interface{}) (*LoopStatement, error) {
	loopstmt := LoopStatement{}

	if typ, ok := m["Type"]; !ok || typ != "LOOP" {
		return nil, errors.New("the generic map supplied is not a LoopStatement")
	}

	loopstmt.Type = m["Type"].(string)

	if line, ok := m["Line"]; ok {
		loopstmt.Line = int(line.(float64))
	}

	if stmts, ok := m["StmtList"]; ok && stmts != nil {
		loopstmt.StmtList = make([]Statement, 0)

		for _, stmt := range m["StmtList"].([]interface{}) {
			castStmt, err := castToStatement(stmt.(map[string]interface{}))
			if err != nil {
				return nil, err
			}

			loopstmt.StmtList = append(loopstmt.StmtList, castStmt)
		}
	}

	return &loopstmt, nil
}
