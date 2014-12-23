// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "errors"

// TODO rename this type into LoopStatement and the loop stop condition
type ForStatement struct {
	Type     string      `json:"type"`
	StmtList []Statement `json:"statements_list"`
	Line     int         `json:"line"` // Line number of the statement relatively to the function.
}

// CastToForStatement "cast" a generic map into a ForStatement.
func CastToForStatement(m map[string]interface{}) (*ForStatement, error) {
	forstmt := ForStatement{}

	// TODO remove FOR
	if typ, ok := m["Type"]; !ok || (typ != "FOR" && typ != "LOOP") {
		return nil, errors.New("the generic map supplied is not a ForStatement")
	}

	forstmt.Type = m["Type"].(string)

	if line, ok := m["Line"]; ok {
		forstmt.Line = int(line.(float64))
	}

	if stmts, ok := m["StmtList"]; ok && stmts != nil {
		forstmt.StmtList = make([]Statement, 0)

		for _, stmt := range m["StmtList"].([]interface{}) {
			castStmt, err := castToStatement(stmt.(map[string]interface{}))
			if err != nil {
				return nil, err
			}

			forstmt.StmtList = append(forstmt.StmtList, castStmt)
		}
	}

	return &forstmt, nil
}
