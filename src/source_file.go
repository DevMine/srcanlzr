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

	// List of the imports used by the srouce file.
	Imports []string `json:"imports,omitempty"`

	// Types definition
	// TODO rename JSON key into type_definition
	TypeDefs []TypeDef `json:"type_defs,omitempty"`

	// Structures definition
	// TODO rename JSON key into structures
	Structs []StructuredType `json:"structs,omitempty"`

	// List of constants defined at the file level (e.g. global constants)
	Constants []Constant `json:"constants,omitempty"`

	// List of variables defined at the file level (e.g. global variables)
	Variables []Variable `json:"variables,omitempty"`

	// List of functions
	Functions []Function `json:"functions,omitempty"`

	// List of interfaces
	Interfaces []Interface `json:"interfaces,omitempty"`

	// List of classes
	Classes []Class `json:"classes,omitempty"`

	// List of traits
	// See http://en.wikipedia.org/wiki/Trait_%28computer_programming%29
	Traits []Trait `json:"traits,omitempty"`

	// The total number of lines of code.
	LoC int64 `json:"loc"`
}
