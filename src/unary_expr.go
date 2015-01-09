// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

type UnaryExpr struct {
}

func newUnaryExpr(m map[string]interface{}) (*UnaryExpr, error) {
	//var err error
	//errPrefix := "src/decl_stmt"
	unaryexpr := UnaryExpr{}

	return &unaryexpr, nil
}
