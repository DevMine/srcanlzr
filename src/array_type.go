// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

type ArrayType struct {
	Len int64 `json:"length,omitempty"`
	Elt Expr  `json:"element_type"` // element type
}

func newArrayType(m map[string]interface{}) (*ArrayType, error) {
	var err error
	errPrefix := "src/array_type"
	arytype := ArrayType{}

	if arytype.Len, err = extractInt64Value("length", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	exprMap, err := extractMapValue("element_type", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if arytype.Elt, err = newExpr(exprMap); err != nil {
		return nil, addDebugInfo(err)
	}

	return &arytype, nil
}
