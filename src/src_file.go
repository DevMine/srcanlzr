// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "errors"

// SourceFile holds information about a source file.
type SrcFile struct {
	// The path of the source file, relative to the root of the project.
	Path string `json:"path"`

	// Programming language used.
	Lang *Language `json:"language"`

	// List of the imports used by the srouce file.
	Imports []string `json:"imports,omitempty"`

	// Types definition
	TypeSpecs []*TypeSpec `json:"type_specifiers,omitempty"`

	// Structures definition
	// TODO rename JSON key into structures
	Structs []*StructType `json:"structs,omitempty"`

	// List of constants defined at the file level (e.g. global constants)
	Constants []*Constant `json:"constants,omitempty"`

	// List of variables defined at the file level (e.g. global variables)
	Vars []*Var `json:"variables,omitempty"`

	// List of functions
	Funcs []*FuncDecl `json:"functions,omitempty"`

	// List of interfaces
	Interfaces []*Interface `json:"interfaces,omitempty"`

	// List of classes
	Classes []*ClassDecl `json:"classes,omitempty"`

	// List of enums
	Enums []*EnumDecl `json:"enums,omitempty"`

	// List of traits
	// See http://en.wikipedia.org/wiki/Trait_%28computer_programming%29
	Traits []*Trait `json:"traits,omitempty"`

	// The total number of lines of code.
	LoC int64 `json:"loc"`
}

func newSrcFile(m map[string]interface{}) (*SrcFile, error) {
	var err error
	errPrefix := "src/src"
	src := SrcFile{}

	if src.Path, err = extractStringValue("path", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	progLangMap, err := extractMapValue("language", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if src.Lang, err = newLanguage(progLangMap); err != nil {
		return nil, addDebugInfo(err)
	}

	if src.LoC, err = extractInt64Value("loc", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if src.Imports, err = extractStringSliceValue("imports", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if src.TypeSpecs, err = newTypeSpecsSlice("type_specifiers", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if src.Structs, err = newStructTypesSlice("structs", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if src.Constants, err = newConstantsSlice("constants", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if src.Vars, err = newVarsSlice("variables", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if src.Funcs, err = newFuncDeclsSlice("functions", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if src.Interfaces, err = newInterfacesSlice("interfaces", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if src.Classes, err = newClassDeclsSlice("classes", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if src.Traits, err = newTraitsSlice("traits", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	return &src, nil
}

func mergeSrcFilesSlices(sfs1, sfs2 []*SrcFile) ([]*SrcFile, error) {
	if sfs1 == nil {
		return nil, addDebugInfo(errors.New("ps1 cannot be nil"))
	}

	if sfs2 == nil {
		return nil, addDebugInfo(errors.New("ps2 cannot be nil"))
	}

	newSfs := make([]*SrcFile, 0)
	newSfs = append(newSfs, sfs1...)
	newSfs = append(newSfs, sfs2...)

	return newSfs, nil
}
