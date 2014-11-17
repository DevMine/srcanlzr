// Copyright 2014 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	Package src provides a set of structures for representing a source code
	indepently of the language. In other words, it provides a generic
	representation of source code.
*/
package src

import (
	"encoding/xml"
	"net/url"

	"github.com/DevMine/srcanlzr/repo"
)

// Source file type
// XXX is it really useful? Is it possible to abstract that ? Because this
// strongly related to C and it does not really fit with generic representation
// of a source code...
const (
	Source = iota
	Header
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

type Project struct {
	Name      string // Name of the project
	Repo      *repo.Repo
	RepoURL   *url.URL   // Repo URL
	ProgLangs []Language // List of all programming languages used
	Packages  []*Package // Project's packages
	LoC       int64      // Lines of Code
}

type Language struct {
	XMLName   xml.Name `json:"-" xml:"languages"`
	Lang      int      `json:"language" xml:"language"`
	Paradigms []int    `json:"paradigms" xml:"paradigm"`
}

type Package struct {
	Name        string        // Package name
	Path        string        // Package location
	Doc         string        // Package doc comments
	SourceFiles []*SourceFile // Source files
	LoC         int64         // Lines of Code
}

// TODO add support of structs and types
type SourceFile struct {
	Path       string      // Path is the source file location
	Type       int         // Type of source file (source, header, ...)
	ProgLang   *Language   // Programming language
	Encoding   string      // Encoding of the source file
	MIME       string      // MIME type as defined in RFC 2046
	Imports    []string    // Imports
	TypeDefs   []TypeDef   // TypeDefs
	Structs    []Struct    // Struct
	Constants  []Constant  // Constants
	Variables  []Variable  // Variables
	Functions  []Function  // Functions
	Interfaces []Interface // Interfaces
	Classes    []Class     // Classes
	Modules    []Module    // Modules (sometimes called "trait")
	LoC        int64       // Lines of Code
}

// TODO it would be nice to only have the raw function prototype instead of the
// whole function.
// TODO proposal: keep only positions instead of a raw string. This would make
// the parsing faster and the generated JSON smaller.
type Function struct {
	Name     string
	Comments string
	Args     []Variable
	Return   []Variable
	StmtList []Statement
	LoC      int64  // Lines of Code
	Raw      string // Function raw source code.
}

type Statement interface{}

type IfStatement struct {
	Type     string
	StmtList []Statement
	Line     int // Line number of the statement relatively to the function.
}

type ForStatement struct {
	Type     string
	StmtList []Statement
	Line     int // Line number of the statement relatively to the function.
}

type CallStatement struct {
	Type     string
	Ref      FuncRef // Reference to the function
	StmtList []Statement
	Line     int // Line number of the statement relatively to the function.
}

type AssignStatement struct {
	Type     string
	VarName  string
	VarValue string // TODO handle case where value is a literal, function call, etc.
	Line     int    // Line number of the statement relatively to the function.
}

type OtherStatement struct {
	Type     string
	StmtList []Statement
	Line     int // Line number of the statement relatively to the function.
}

type FuncRef struct {
	Namespace string
	FuncName  string
}

type Variable struct {
	Name  string
	Type  string
	Value string
	Doc   string
}

type Constant struct {
	Name  string
	Type  string
	Value string
	Doc   string
}

type TypeDef struct {
	Name string
	Doc  string
	Type ExprType
}

// ExprType can be either a structured type (Struct) or another type (OtherType)
// such as an int, float, etc.
type ExprType interface{}

type OtherType struct {
	Name string // int, float, string, etc.
}

type Struct struct {
	Name   string
	Doc    string
	Fields []Field
}

type Field struct {
	Name string
	Type string
	Doc  string
}

type Class struct {
	Name                  string
	Visiblity             int
	ExtendedClasses       []*Class
	ImplementedInterfaces []*Interface
	Attributes            []*Attribute
	Methods               []*Method
	Modules               []*Module // For languages supporting mixins
}

type Interface struct {
	Name       string
	Visibility int
	// TODO
}

type Module struct {
	Name       string
	Attributes []*Attribute
	Methods    []*Method
	Modules    []*Module // For languages supporting mixins
}

type Attribute struct {
	Variable
	Visibility int
	Constant   bool
	Static     bool
}

type Method struct {
	Function
	Visibility int
}
