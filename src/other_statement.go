// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "errors"

type OtherStatement struct {
	Type     string      `json:"type"`
	StmtList []Statement `json:"statements_list"`
	Line     int         `json:"line"` // Line number of the statement relatively to the function.
}

// CastToOtherStatement "cast" a generic map into a OtherStatement.
func CastToOtherStatement(m map[string]interface{}) (*OtherStatement, error) {
	otherstmt := OtherStatement{}

	if typ, ok := m["Type"]; !ok || typ != "OTHER" {
		return nil, errors.New("the generic map supplied is not a OtherStatement")
	}

	otherstmt.Type = m["Type"].(string)

	if line, ok := m["Line"]; ok {
		otherstmt.Line = int(line.(float64))
	}

	if stmts, ok := m["StmtList"]; ok && stmts != nil {
		otherstmt.StmtList = make([]Statement, 0)

		for _, stmt := range m["StmtList"].([]interface{}) {
			castStmt, err := castToStatement(stmt.(map[string]interface{}))
			if err != nil {
				return nil, err
			}

			otherstmt.StmtList = append(otherstmt.StmtList, castStmt)
		}
	}

	return &otherstmt, nil
}
