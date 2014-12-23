// Copyright 2014 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "errors"

type CallStatement struct {
	Type string  `json:"type"`
	Ref  FuncRef `json:"reference"` // Reference to the function
	Line int     `json:"line"`      // Line number of the statement relatively to the function.
}

// CastToCallStatement "cast" a generic map into a CallStatement.
func CastToCallStatement(m map[string]interface{}) (*CallStatement, error) {
	callstmt := CallStatement{}

	if typ, ok := m["Type"]; !ok || typ != "CALL" {
		return nil, errors.New("the generic map supplied is not a CallStatement")
	}

	callstmt.Type = m["Type"].(string)

	if line, ok := m["Line"]; ok {
		callstmt.Line = int(line.(float64))
	}

	if ref, ok := m["Ref"]; ok {
		ref, err := CastToFuncRef(ref.(map[string]interface{}))
		if err != nil {
			return nil, err
		}

		callstmt.Ref = *ref
	}

	return &callstmt, nil
}
