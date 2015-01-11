// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

type ArrayLit struct {
	Type *ArrayType `json:"type"`
	Elts []Expr     `json:"elements"`
}

func newArrayLit(m map[string]interface{}) (*ArrayLit, error) {
	var err error
	errPrefix := "src/array_lit"
	arylit := ArrayLit{}

	typeMap, err := extractMapValue("type", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if arylit.Type, err = newArrayType(typeMap); err != nil {
		return nil, addDebugInfo(err)
	}

	if arylit.Elts, err = newExprsSlice("elements", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &arylit, nil
}
