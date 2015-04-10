// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	Package src provides a set of structures for representing a source code
	indepently of the language. In other words, it provides a generic
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

	"github.com/DevMine/repotool/model"
	"github.com/DevMine/srcanlzr/src/ast"
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

// A Language represents a programming language.
type Language struct {
	// The programming language name (e.g. go, ruby, java, etc.)
	//
	// The name must match one of the supported programming languages defined in
	// the constants.
	Lang string `json:"language"` // TODO rename into name

	// The paradigms of the programming language (e.g. structured, imperative,
	// object oriented, etc.)
	//
	// The name must match one of the supported paradigms defined in the
	// constants.
	Paradigms []string `json:"paradigms"`
}

// Project is the root of the src API and therefore it must be at the root of
// the JSON.
//
// It contains the metadata of a project and the list of all packages.
type Project struct {
	// The name of the project. Since it may be something really difficult to
	// guess, it should generally be the name of the folder containing the
	// project.
	Name string `json:"name"`

	// The repository in which the project is hosted, or nil. This field is not
	// meant to be filled by one of the language parsers. Only repotool should
	// take care of it. For more details, see:
	//    https://github.com/DevMine/repotool
	//
	// Since this field uses an external type, it is not unmarshalled by
	// src.Unmarshal itself but by the standard json.Unmarshal function.
	// To do so, its unmarshalling is defered using json.RawMessage.
	// See the RepoRaw field.
	Repo *model.Repository `json:"-"`

	// RepoRaw is only used to defer the unmarshalling of a model.Repository.
	RepoRaw json.RawMessage `json:"repository,omitempty"`

	// The list of all programming languages used by the project. Each language
	// must be added by the corresponding language parsers if and only if the
	// project contains at least one line of code written in this language.
	Langs []*Language `json:"languages"`

	// List of all packages of the project. We call "package" every folder that
	// contains at least one source file.
	Packages []*Package `json:"packages"`

	// The total number of lines of code in the whole project, independently of
	// the language.
	LoC int64 `json:"loc"`
}

// Package holds information about a package, which is, basically, just a
// folder that contains at least one source file.
type Package struct {
	// The package documentation, or nil.
	// TODO support docucmentation for multiple languages.
	Doc []string `json:"doc,omitempty"`

	// The package name. This should be the name of the parent folder.
	Name string `json:"name"`

	// The full path of the package. The path must be relative to the root of
	// the project and never be an absolute path.
	Path string `json:"path"`

	// The list of all source files contained in the package.
	SrcFiles []*SrcFile `json:"source_files"`

	// The total number of lines of code of the package.
	LoC int64 `json:"loc"`
}

// SrcFile holds information about a source file.
type SrcFile struct {
	// The path of the source file, relative to the root of the project.
	Path string `json:"path"`

	// Programming language used.
	Lang *Language `json:"language"`

	// List of the imports used by the srouce file.
	Imports []string `json:"imports,omitempty"`

	// Types definition
	TypeSpecs []*ast.TypeSpec `json:"type_specifiers,omitempty"`

	// Structures definition
	// TODO rename JSON key into structures
	Structs []*ast.StructType `json:"structs,omitempty"`

	// List of constants defined at the file level (e.g. global constants)
	Constants []*ast.GlobalDecl `json:"constants,omitempty"`

	// List of variables defined at the file level (e.g. global variables)
	Vars []*ast.GlobalDecl `json:"variables,omitempty"`

	// List of functions
	Funcs []*ast.FuncDecl `json:"functions,omitempty"`

	// List of interfaces
	Interfaces []*ast.Interface `json:"interfaces,omitempty"`

	// List of classes
	Classes []*ast.ClassDecl `json:"classes,omitempty"`

	// List of enums
	Enums []*ast.EnumDecl `json:"enums,omitempty"`

	// List of traits
	// See http://en.wikipedia.org/wiki/Trait_%28computer_programming%29
	Traits []*ast.Trait `json:"traits,omitempty"`

	// The total number of lines of code.
	LoC int64 `json:"loc"`
}

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
