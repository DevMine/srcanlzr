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
	"encoding/json"
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

// Unmarshal unmarshals a JSON representation of a Project into a real
//  Project structure.
//
// It is required to use this function instead of json.Unmarshal because we use
// an interface to abstract a Statement, thus json.Unmarshal is unable to
// unmarshal the statements correctly.
func Unmarshal(bs []byte) (*Project, error) {
	genMap := map[string]interface{}{}
	if err := json.Unmarshal(bs, &genMap); err != nil {
		return nil, addDebugInfo(err)
	}

	prj, err := newProject(genMap)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	return prj, nil
}

// MergeAll merges a list of projects.
//
// There must be at least one project. In this case, it just returns a copy of
// the project.
//
// The merge only performs shallow copies, which means that if the field value
// is a pointer it copies the memory address and not the value pointed.
func MergeAll(ps ...*Project) (*Project, error) {
	if len(ps) == 0 {
		return nil, addDebugInfo(errors.New("p cannot be nil"))
	}

	newPrj := &Project{
		Name:      ps[0].Name,
		ProgLangs: ps[0].ProgLangs,
		Packages:  ps[0].Packages,
		LoC:       ps[0].LoC,
	}

	if len(ps) == 1 {
		return newPrj, nil
	}

	var err error
	for i := 1; i < len(ps); i++ {
		curr := ps[i]
		if curr == nil {
			return nil, addDebugInfo(errors.New(fmt.Sprintf("p[%d] is nil", i)))
		}

		if newPrj, err = Merge(newPrj, curr); err != nil {
			return nil, addDebugInfo(err)
		}
	}

	return newPrj, nil
}

// Merge merges two project. See MergeAll for more details.
func Merge(p1, p2 *Project) (*Project, error) {
	return mergeProjects(p1, p2)
}
