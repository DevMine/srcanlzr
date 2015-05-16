// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	Package src provides a set of structures for representing a source code
	independently of the language. In other words, it provides a generic
	representation (abstraction) of a source code.

	The number of real lines of code must be precomputed by the language
	parsers. This is the only "feature" that must be precomputed because it is
	used by srctool (https://github.com/DevMine/srctool) to eliminate empty
	packages. Since this information is already calculated, the source analyzers
	won't re-count. They will just use the total as it is. Therefore, it must be
	accurate.

	We only count statements and declarations as a line of code. Comments,
	package declaration, imports, expression, etc. must not be taken into
	account. Since an exemple is worth more than a thousand words, let's
	condider the following snippet:
	   // Package doc (does not count as a line of code)
	   package main // does not count as a line of code

	   import "fmt" // does not count as a line of code

	   func main() { // count as 1 line of code
	     fmt.Println(
	        "Hello, World!
	     ) // count as 1 line of code
	   }

	The expected number of lines of code is 2: The main function declaration
	and the call to the fmt.Println function.
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

// Type names
const (
	TypeMapName         = "MAP"
	TypeStructName      = "STRUCT"
	TypeArrayName       = "ARRAY"
	TypeFuncName        = "FUNC"
	TypeInterfaceName   = "INTERFACE"
	TypeUnsupportedName = "UNSUPPORTED"
)

// Unmarshal parses a JSON-encoded src.Project and returns it.
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

// Marshal returns the JSON encoding of src.Project p.
// It is just a wrapper for json.Marshal. For more details, see:
//    http://golang.org/pkg/encoding/json/#Marshal
//
// This function only serves a semantic purpose. Since the src package must wrap
// json.Unmarshal function, it makes sense to also provides a Marshal function.
func Marshal(p *Project) ([]byte, error) {
	bs, err := json.Marshal(p.Repo)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	p.RepoRaw = json.RawMessage(bs)

	if bs, err = json.Marshal(p); err != nil {
		return nil, addDebugInfo(err)
	}
	return bs, nil
}

// spaces to use when marshalling
const indentSpaces = "    "

// MarshalIndent is like Marshal but applies Indent to format the output.
// It uses 4 spaces for indentation.
func MarshalIndent(p *Project) ([]byte, error) {
	bs, err := json.Marshal(p.Repo)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	p.RepoRaw = json.RawMessage(bs)

	if bs, err = json.MarshalIndent(p, "", indentSpaces); err != nil {
		return nil, addDebugInfo(err)
	}
	return bs, nil
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

	if ps[0] == nil {
		return nil, addDebugInfo(errors.New("p[0] cannot be nil"))
	}

	newPrj := &Project{
		Name:     ps[0].Name,
		Langs:    ps[0].Langs,
		Packages: ps[0].Packages,
		LoC:      ps[0].LoC,
	}

	if len(ps) == 1 {
		return newPrj, nil
	}

	var err error
	for i := 1; i < len(ps); i++ {
		curr := ps[i]
		if curr == nil {
			return nil, addDebugInfo(fmt.Errorf("p[%d] is nil", i))
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
