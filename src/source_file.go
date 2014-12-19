// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

// TODO add support of structs and types
type SourceFile struct {
	Path       string           // Path is the source file location
	Type       int              // Type of source file (source, header, ...)
	ProgLang   *Language        // Programming language
	Encoding   string           // Encoding of the source file
	MIME       string           // MIME type as defined in RFC 2046
	Imports    []string         // Imports
	TypeDefs   []TypeDef        // TypeDefs
	Structs    []StructuredType // Struct
	Constants  []Constant       // Constants
	Variables  []Variable       // Variables
	Functions  []Function       // Functions
	Interfaces []Interface      // Interfaces
	Classes    []Class          // Classes
	Modules    []Module         // Modules (sometimes called "trait")
	LoC        int64            // Lines of Code
}
