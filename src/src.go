// Copyright 2014 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	Package src provides a set of structures for representing a source code
	indepently of the language. In other words, it provides a generic
	representation (abstraction) of a source code.
*/
package src

import (
	"errors"
	"fmt"
)

// Programming languages
const (
	Go = iota
	Ruby
	Python
	C
	Java
	Scala
	// ... and so on
)

// Visibilities
const (
	PublicVisibility = iota
	PackageVisibility
	ProtectedVisibility
	PrivateVisibility
)

// Paradigms
const (
	Strucured = iota
	Imperative
	Procedural
	Compiled
	Concurrent
	Functional
	ObjectOriented
	Generic
	Reflective
)

// castToStatement cast appropriately a given general map into a Statement.
//func castToStatement(m map[string]interface{}) (Statement, error) {
func castToStatement(m map[string]interface{}) (Statement, error) {
	if _, ok := m["Type"]; !ok {
		return nil, errors.New("statements list contains an element that is not a Statement")
	}

	switch m["Type"] {
	case "IF":
		return CastToIfStatement(m)
	case "LOOP", "FOR": // TODO remove FOR
		return CastToForStatement(m)
	case "ASSIGN":
		return CastToAssignStatement(m)
	case "CALL":
		return CastToCallStatement(m)
	case "OTHER":
		return CastToOtherStatement(m)
	}

	fmt.Println(m["Type"])

	return nil, errors.New("unknown statement")
}
