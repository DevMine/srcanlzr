// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

type IndexExpr struct {
	ExprName string `json:"expression_name"`
	X        Expr   `json:"expression,omitempty"` // expression
	Index    Expr   `json:"index,omitempty"`      // index expression
}

func newIndexExpr(m map[string]interface{}) (*IndexExpr, error) {
	var err error
	errPrefix := "src/index_expr"
	indexpr := IndexExpr{}

	if typ, err := extractStringValue("expression_name", errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	} else if typ != IndexExprName {
		return nil, fmt.Errorf("invalid type: expected 'IndexExprName', found '%s'", typ)
	}

	indexpr.ExprName = IndexExprName

	exprMap, err := extractMapValue("expression", errPrefix, m)
	if err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	} else if err == nil {
		if indexpr.X, err = newExpr(exprMap); err != nil {
			return nil, addDebugInfo(err)
		}
	}

	if exprMap, err = extractMapValue("expression", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	} else if err == nil {
		if indexpr.Index, err = newExpr(exprMap); err != nil {
			return nil, addDebugInfo(err)
		}
	}

	return &indexpr, nil
}
