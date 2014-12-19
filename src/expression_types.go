// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

// ExprType can be either a structured type (Struct) or another type (PrimitiveType)
// such as an int, float, etc.
type ExprType interface{}

type PrimitiveType struct {
	Name string // int, float, string, etc.
}

type StructuredType struct {
	Name   string
	Doc    string
	Fields []Field
}

type Field struct {
	Name string
	Type string
	Doc  string
}
