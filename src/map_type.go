// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

type MapType struct {
	KeyType   Expr `json:"key_type"`
	ValueType Expr `json:"value_type"`
}

func newMapType(m map[string]interface{}) (*MapType, error) {
	var err error
	errPrefix := "src/map_type"
	maptype := MapType{}

	exprMap, err := extractMapValue("key_type", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if maptype.KeyType, err = newExpr(exprMap); err != nil {
		return nil, addDebugInfo(err)
	}

	exprMap, err = extractMapValue("value_type", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if maptype.ValueType, err = newExpr(exprMap); err != nil {
		return nil, addDebugInfo(err)
	}

	return &maptype, nil
}
