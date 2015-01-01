// Copyright 2014-2015 The DevMine Authors. All rights reserved.
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

// Supported VCS (Version Control System)
const (
	Git = "git"
	Hg  = "mercurial"
	SVN = "subversion"
	Bzr = "bazaar"
	CVS = "cvs"
)

var suppVCS = []string{
	Git,
	Hg,
	SVN,
	Bzr,
	CVS,
}

// Supported programming languages
const (
	Go     = "go"
	Ruby   = "ruby"
	Python = "python"
	C      = "c"
	Java   = "java"
	Scala  = "scala"
)

var suppLang = []string{
	Go,
	Ruby,
	Python,
	C,
	Java,
	Scala,
}

// Supported visiblities
const (
	PublicVisibility    = "public"
	PackageVisibility   = "package"
	ProtectedVisibility = "protected"
	PrivateVisibility   = "private"
)

var suppVisibility = []string{
	PublicVisibility,
	PackageVisibility,
	ProtectedVisibility,
	PrivateVisibility,
}

// Supported paradigms
const (
	Structured     = "structured"
	Imperative     = "imperative"
	Procedural     = "procedural"
	Compiled       = "compiled"
	Concurrent     = "concurrent"
	Functional     = "functional"
	ObjectOriented = "object oriented"
	Generic        = "generic"
	Reflective     = "reflective"
)

var suppParadigms = []string{
	Structured,
	Imperative,
	Procedural,
	Compiled,
	Concurrent,
	Functional,
	ObjectOriented,
	Generic,
	Reflective,
}

// castToStatement cast appropriately a given general map into a Statement.
//func castToStatement(m map[string]interface{}) (Statement, error) {
func castToStatement(m map[string]interface{}) (Statement, error) {
	if _, ok := m["Type"]; !ok {
		return nil, errors.New("statements list contains an element that is not a Statement")
	}

	switch m["Type"] {
	case "IF":
		return NewIfStatement(m)
	case "LOOP", "FOR": // TODO remove FOR
		return NewLoopStatement(m)
	case "ASSIGN":
		return NewAssignStatement(m)
	case "CALL":
		return NewCallStatement(m)
	case "OTHER":
		return NewOtherStatement(m)
	}

	fmt.Println(m["Type"])

	return nil, errors.New("unknown statement")
}
