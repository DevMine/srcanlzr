// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

// ExprType can be either a structured type (Struct) or another type (PrimitiveType)
// such as an int, float, etc.
type ExprType interface{}

type PrimitiveType struct {
	Name string `json:"name"` // int, float, string, etc.
}

type StructuredType struct {
	Name   string  `json:"name"`
	Doc    string  `json:"doc"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Doc  string `json:"doc"`
}
