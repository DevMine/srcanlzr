// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

type ArrayType struct {
	// Dimensions
	Dims []int64 `json:"dimensions"`

	Elt Expr `json:"element_type,omitempty"` // element type
}

func newArrayType(m map[string]interface{}) (*ArrayType, error) {
	var err error
	errPrefix := "src/array_type"
	arytype := ArrayType{}

	if arytype.Dims, err = extractInt64SliceValue("dimensions", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	exprMap, err := extractMapValue("element_type", errPrefix, m)
	if err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	} else if err == nil {
		if arytype.Elt, err = newExpr(exprMap); err != nil {
			return nil, addDebugInfo(err)
		}
	}

	return &arytype, nil
}
