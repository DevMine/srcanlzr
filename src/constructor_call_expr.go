// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

type ConstructorCallExpr struct {
	CallExpr
}

func newConstructorCallExpr(m map[string]interface{}) (*ConstructorCallExpr, error) {
	var err error
	errPrefix := "src/constructor_call_expr"
	ccallexpr := ConstructorCallExpr{}

	if typ, err := extractStringValue("expression_name", errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	} else if typ != ConstructorCallExprName {
		return nil, fmt.Errorf("invalid type: expected 'ConstructorCallExpr', found '%s'", typ)
	}

	ccallexpr.ExprName = CallExprName

	refMap, err := extractMapValue("function", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if ccallexpr.Fun, err = newFuncRef(refMap); err != nil {
		return nil, addDebugInfo(err)
	}

	if ccallexpr.Args, err = newExprsSlice("arguments", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if ccallexpr.Line, err = extractInt64Value("line", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &ccallexpr, nil
}
