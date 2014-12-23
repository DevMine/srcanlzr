// Copyright 2014 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

// SourceFile holds information about a source file.
type SourceFile struct {
	// The path of the source file, relative to the root of the project.
	Path string `json:"path"`

	// Programming language used.
	ProgLang *Language `json:"language"`

	// File encoding.
	Encoding string `json:"encoding"`

	// TODO remove
	MIME string `json:"mime"`

	// List of the imports used by the srouce file.
	Imports []string `json:"imports"`

	// Types definition
	// TODO rename JSON key into type_definition
	TypeDefs []TypeDef `json:"type_defs"`

	// Structures definition
	// TODO rename JSON key into structures
	Structs []StructuredType `json:"structs"`

	// List of constants defined at the file level (e.g. global constants)
	Constants []Constant `json:"constants"`

	// List of variables defined at the file level (e.g. global variables)
	Variables []Variable `json:"variables"`

	// List of functions
	Functions []Function `json:"functions"`

	// List of interfaces
	Interfaces []Interface `json:"interfaces"`

	// List of classes
	Classes []Class `json:"classes"`

	// List of modules
	// TODO rename into "traits"
	// See http://en.wikipedia.org/wiki/Trait_%28computer_programming%29
	Modules []Module `json:"modules"`

	// The total number of lines of code.
	LoC int64 `json:"loc"`
}
