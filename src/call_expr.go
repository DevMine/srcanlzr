// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

type CallExpr struct {
	ExprName string   `json:"expression_name"`
	Fun      *FuncRef `json:"function"`  // Reference to the function
	Args     []Expr   `json:"arguments"` // function arguments
	Line     int64    `json:"line"`      // Line number of the statement relatively to the function.
}

// newCallExpr creates a new CallExpr from a generic map.
func newCallExpr(m map[string]interface{}) (*CallExpr, error) {
	var err error
	errPrefix := "src/call_expr"
	callexpr := CallExpr{}

	// should never happen
	if typ, ok := m["expression_name"]; !ok || typ != CallExprName {
		return nil, addDebugInfo(fmt.Errorf(
			"%s: the generic map supplied is not a CallExpr", errPrefix))
	}

	callexpr.ExprName = CallExprName

	refMap, err := extractMapValue("function", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if callexpr.Fun, err = newFuncRef(refMap); err != nil {
		return nil, addDebugInfo(err)
	}

	if callexpr.Args, err = newExprsSlice("arguments", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if callexpr.Line, err = extractInt64Value("line", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &callexpr, nil
}
