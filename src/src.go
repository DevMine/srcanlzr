// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	Package src provides a set of structures for representing a source code
	indepently of the language. In other words, it provides a generic
	representation (abstraction) of a source code.
*/
package src

import "encoding/json"

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

// Unmarshal unmarshals a JSON representation of a Project into a real
//  Project structure.
//
// It is required to use this function instead of json.Unmarshal because we use
// an interface to abstract a Statement, thus json.Unmarshal is unable to
// unmarshal the statements correctly.
func Unmarshal(bs []byte) (*Project, error) {
	genMap := map[string]interface{}{}
	if err := json.Unmarshal(bs, &genMap); err != nil {
		return nil, err
	}

	prj, err := newProject(genMap)
	if err != nil {
		return nil, err
	}

	return prj, nil
}
