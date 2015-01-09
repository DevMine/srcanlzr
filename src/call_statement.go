// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"errors"
	"fmt"
)

type CallStatement struct {
	Type string   `json:"type"`
	Ref  *FuncRef `json:"reference"` // Reference to the function
	Line int64    `json:"line"`      // Line number of the statement relatively to the function.
}

// newCallStatement creates a new CallStatement from a generic map.
func newCallStatement(m map[string]interface{}) (*CallStatement, error) {
	var err error
	errPrefix := "src/call_statement"
	callstmt := CallStatement{}

	// should never happen
	if typ, ok := m["type"]; !ok || typ != CallStmtName {
		return nil, addDebugInfo(errors.New(fmt.Sprintf(
			"%s: the generic map supplied is not a AssignStatement", errPrefix)))
	}

	refMap, err := extractMapValue("reference", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if callstmt.Ref, err = newFuncRef(refMap); err != nil {
		return nil, addDebugInfo(err)
	}

	if callstmt.Line, err = extractInt64Value("line", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &callstmt, nil
}
