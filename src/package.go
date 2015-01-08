// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"errors"
	"fmt"
	"reflect"
)

// A package is a folder contaning at least one source file.
type Package struct {
	// The name of the pacakge (the folder name)
	Name string `json:"name"`

	// The full path of the package. The path must be relative to the root of
	// the project and never be an absolute path.
	Path string `json:"path"`

	// The package documentation.
	// FIXME use slice for packages with multiple languages
	Doc string `json:"doc,omitempty"`

	// The list of all source files contained in the package.
	SourceFiles []*SourceFile `json:"source_files"`

	// The total number of lines of code.
	LoC int64 `json:"loc"`
}

func newPackage(m map[string]interface{}) (*Package, error) {
	var err error
	errPrefix := "src/package"
	pkg := Package{}

	if pkg.Name, err = extractStringValue("name", errPrefix, m); err != nil {
		return nil, err
	}

	if pkg.Path, err = extractStringValue("path", errPrefix, m); err != nil {
		return nil, err
	}

	if pkg.Doc, err = extractStringValue("doc", errPrefix, m); err != nil {
		return nil, err
	}

	if pkg.LoC, err = extractInt64Value("loc", errPrefix, m); err != nil {
		return nil, err
	}

	srcsMap, ok := m["source_files"]
	if !ok {
		return nil, errors.New(errPrefix + ": field 'source_files' does not exist")
	}

	var s reflect.Value
	if s = reflect.ValueOf(srcsMap); s.Kind() != reflect.Slice {
		return nil, errors.New(errPrefix + ": field 'source_files' is supposed to be a slice")
	}

	srcs := make([]*SourceFile, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		src := s.Index(i).Interface()

		switch src.(type) {
		case map[string]interface{}:
			if srcs[i], err = newSourceFile(src.(map[string]interface{})); err != nil {
				return nil, err
			}
		default:
			return nil, errors.New(errPrefix + ": 'source_file' must be a map[string]interface{}")
		}
	}

	pkg.SourceFiles = srcs

	return &pkg, nil
}

func newPackagesSlice(key, errPrefix string, m map[string]interface{}) ([]*Package, error) {
	var err error
	var s reflect.Value

	pkgsMap, ok := m[key]
	if !ok {
		return nil, errNotExist
	}

	if s = reflect.ValueOf(pkgsMap); s.Kind() != reflect.Slice {
		return nil, errors.New(fmt.Sprintf("%s: field '%s' is supposed to be a slice",
			errPrefix, key))
	}

	pkgs := make([]*Package, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		pkg := s.Index(i).Interface()

		switch pkg.(type) {
		case map[string]interface{}:
			if pkgs[i], err = newPackage(pkg.(map[string]interface{})); err != nil {
				return nil, err
			}
		default:
			return nil, errors.New(fmt.Sprintf("%s: '%s' must be a map[string]interface{}",
				errPrefix, key))
		}
	}

	return pkgs, nil
}
