// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

type ListType struct {
	Len int64 `json:"length,omitempty"`
	Max int64 `json:"capacity,omitempty"` // maximum capacity
	Elt Expr  `json:"element_type"`
}

func newListType(m map[string]interface{}) (*ListType, error) {
	var err error
	errPrefix := "src/list_type"
	listtype := ListType{}

	if listtype.Len, err = extractInt64Value("length", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if listtype.Max, err = extractInt64Value("Capacity", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	exprMap, err := extractMapValue("element_type", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if listtype.Elt, err = newExpr(exprMap); err != nil {
		return nil, addDebugInfo(err)
	}

	return &listtype, nil
}
