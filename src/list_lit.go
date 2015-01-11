// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

type ListLit struct {
	Type *ListType `json:"type"`
	Elts []Expr    `json:"elements"`
}

func newListLit(m map[string]interface{}) (*ListLit, error) {
	var err error
	errPrefix := "src/list_lit"
	listlit := ListLit{}

	typeMap, err := extractMapValue("type", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if listlit.Type, err = newListType(typeMap); err != nil {
		return nil, addDebugInfo(err)
	}

	if listlit.Elts, err = newExprsSlice("elements", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &listlit, nil
}
