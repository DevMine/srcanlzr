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
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
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
	Name      string     // Name of the project
	Repo      *repo.Repo // Repository in which the project is hosted
	RepoURL   *url.URL   // Repo URL
	ProgLangs []Language // List of all programming languages used
	Packages  []*Package // Project's packages
	LoC       int64      // Lines of Code
}

// UnmarshalProject unmarshals a JSON representation of a Project into a real
//  Project structure.
//
// It is required to use this function instead of json.Unmarshal because we use
// an interface to abstract a Statement, thus json.Unmarshal is unable to
// unmarshal the statements correctly.
//
// TODO Find a more elegant way for solving this problem (eg. write a custom
// JSON parser).
func UnmarshalProject(bs []byte) (*Project, error) {
	p := &Project{}

	if err := json.Unmarshal(bs, p); err != nil {
		return nil, err
	}

	for _, pkgs := range p.Packages {
		for _, sfs := range pkgs.SourceFiles {
			for _, fct := range sfs.Functions {
				castStmts := make([]Statement, 0)

				for _, stmt := range fct.StmtList {
					castStmt, err := castToStatement(stmt.(map[string]interface{}))
					if err != nil {
						return nil, err
					}

					castStmts = append(castStmts, castStmt)
				}

				fct.StmtList = castStmts
			}

			for _, cls := range sfs.Classes {
				for _, mds := range cls.Methods {
					castStmts := make([]Statement, 0)

					for _, stmt := range mds.StmtList {
						castStmt, err := castToStatement(stmt.(map[string]interface{}))
						if err != nil {
							return nil, err
						}

						castStmts = append(castStmts, castStmt)
					}
				}
			}

			for _, mods := range sfs.Modules {
				for _, mds := range mods.Methods {
					castStmts := make([]Statement, 0)

					for _, stmt := range mds.StmtList {
						castStmt, err := castToStatement(stmt.(map[string]interface{}))
						if err != nil {
							return nil, err
						}

						castStmts = append(castStmts, castStmt)
					}
				}
			}
		}
	}

	return p, nil
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
// FIXME add the possibility to have multiple return statements
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

// CastToIfStatement "cast" a generic map into a IfStatement.
func CastToIfStatement(m map[string]interface{}) (*IfStatement, error) {
	ifstmt := IfStatement{}

	if typ, ok := m["Type"]; !ok || typ != "IF" {
		return nil, errors.New("the generic map supplied is not a IfStatement")
	}

	ifstmt.Type = m["Type"].(string)

	if line, ok := m["Line"]; ok {
		ifstmt.Line = int(line.(float64))
	}

	if stmts, ok := m["StmtList"]; ok && stmts != nil {
		ifstmt.StmtList = make([]Statement, 0)

		for _, stmt := range m["StmtList"].([]interface{}) {
			castStmt, err := castToStatement(stmt.(map[string]interface{}))
			if err != nil {
				return nil, err
			}

			ifstmt.StmtList = append(ifstmt.StmtList, castStmt)
		}
	}

	return &ifstmt, nil
}

// TODO rename this type into LoopStatement and the loop stop condition
type ForStatement struct {
	Type     string
	StmtList []Statement
	Line     int // Line number of the statement relatively to the function.
}

// CastToForStatement "cast" a generic map into a ForStatement.
func CastToForStatement(m map[string]interface{}) (*ForStatement, error) {
	forstmt := ForStatement{}

	// TODO remove FOR
	if typ, ok := m["Type"]; !ok || (typ != "FOR" && typ != "LOOP") {
		return nil, errors.New("the generic map supplied is not a ForStatement")
	}

	forstmt.Type = m["Type"].(string)

	if line, ok := m["Line"]; ok {
		forstmt.Line = int(line.(float64))
	}

	if stmts, ok := m["StmtList"]; ok && stmts != nil {
		forstmt.StmtList = make([]Statement, 0)

		for _, stmt := range m["StmtList"].([]interface{}) {
			castStmt, err := castToStatement(stmt.(map[string]interface{}))
			if err != nil {
				return nil, err
			}

			forstmt.StmtList = append(forstmt.StmtList, castStmt)
		}
	}

	return &forstmt, nil
}

type CallStatement struct {
	Type string
	Ref  FuncRef // Reference to the function
	Line int     // Line number of the statement relatively to the function.
}

// CastToCallStatement "cast" a generic map into a CallStatement.
func CastToCallStatement(m map[string]interface{}) (*CallStatement, error) {
	callstmt := CallStatement{}

	if typ, ok := m["Type"]; !ok || typ != "CALL" {
		return nil, errors.New("the generic map supplied is not a CallStatement")
	}

	callstmt.Type = m["Type"].(string)

	if line, ok := m["Line"]; ok {
		callstmt.Line = int(line.(float64))
	}

	if ref, ok := m["Ref"]; ok {
		ref, err := CastToFuncRef(ref.(map[string]interface{}))
		if err != nil {
			return nil, err
		}

		callstmt.Ref = *ref
	}

	return &callstmt, nil
}

type AssignStatement struct {
	Type     string
	VarName  string
	VarValue string // TODO handle case where value is a literal, function call, etc.
	Line     int    // Line number of the statement relatively to the function.
}

// CastToAssignStatement "cast" a generic map into a AssignStatement.
func CastToAssignStatement(m map[string]interface{}) (*AssignStatement, error) {
	assignstmt := AssignStatement{}

	if typ, ok := m["Type"]; !ok || typ != "ASSIGN" {
		return nil, errors.New("the generic map supplied is not a AssignStatement")
	}

	assignstmt.Type = m["Type"].(string)

	if line, ok := m["Line"]; ok {
		// XXX unsafe cast
		assignstmt.Line = int(line.(float64))
	}

	if varName, ok := m["VarName"]; ok {
		assignstmt.VarName = varName.(string)
	}

	if varValue, ok := m["VarValue"]; ok {
		assignstmt.VarValue = varValue.(string)
	}

	return &assignstmt, nil
}

type OtherStatement struct {
	Type     string
	StmtList []Statement
	Line     int // Line number of the statement relatively to the function.
}

// CastToOtherStatement "cast" a generic map into a OtherStatement.
func CastToOtherStatement(m map[string]interface{}) (*OtherStatement, error) {
	otherstmt := OtherStatement{}

	if typ, ok := m["Type"]; !ok || typ != "OTHER" {
		return nil, errors.New("the generic map supplied is not a OtherStatement")
	}

	otherstmt.Type = m["Type"].(string)

	if line, ok := m["Line"]; ok {
		otherstmt.Line = int(line.(float64))
	}

	if stmts, ok := m["StmtList"]; ok && stmts != nil {
		otherstmt.StmtList = make([]Statement, 0)

		for _, stmt := range m["StmtList"].([]interface{}) {
			castStmt, err := castToStatement(stmt.(map[string]interface{}))
			if err != nil {
				return nil, err
			}

			otherstmt.StmtList = append(otherstmt.StmtList, castStmt)
		}
	}

	return &otherstmt, nil
}

type FuncRef struct {
	Namespace string
	FuncName  string
}

// CastToFuncRef "cast" a generic map into a FuncRef.
func CastToFuncRef(m map[string]interface{}) (*FuncRef, error) {
	fctref := FuncRef{}

	var ok bool

	var namespace interface{}
	if namespace, ok = m["Namespace"]; !ok {
		return nil, errors.New("malformed FuncRef, no Namespace field")
	}

	var funcName interface{}
	if funcName, ok = m["FuncName"]; !ok {
		return nil, errors.New("malformed FuncRef, no FuncName field")
	}

	fctref.Namespace = namespace.(string)
	fctref.FuncName = funcName.(string)

	return &fctref, nil
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
