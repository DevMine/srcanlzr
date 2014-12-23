// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

// TODO add support of structs and types
type SourceFile struct {
	Path       string           `json:"path"`       // Path is the source file location
	ProgLang   *Language        `json:"language"`   // Programming language
	Encoding   string           `json:"encoding"`   // Encoding of the source file
	MIME       string           `json:"mime"`       // MIME type as defined in RFC 2046 / TODO remove?
	Imports    []string         `json:"imports"`    // Imports
	TypeDefs   []TypeDef        `json:"type_defs"`  // TypeDefs
	Structs    []StructuredType `json:"structs"`    // Struct
	Constants  []Constant       `json:"constants"`  // Constants
	Variables  []Variable       `json:"variables"`  // Variables
	Functions  []Function       `json:"functions"`  // Functions
	Interfaces []Interface      `json:"interfaces"` // Interfaces
	Classes    []Class          `json:"classes"`    // Classes
	Modules    []Module         `json:"modules"`    // Modules (sometimes called "trait")
	LoC        int64            `json:"loc"`        // Lines of Code
}
