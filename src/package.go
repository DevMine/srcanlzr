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
	// The package documentation.
	// FIXME use slice for packages with multiple languages
	Doc []string `json:"doc,omitempty"`

	// The name of the pacakge (the folder name)
	Name string `json:"name"`

	// The full path of the package. The path must be relative to the root of
	// the project and never be an absolute path.
	Path string `json:"path"`

	// The list of all source files contained in the package.
	SrcFiles []*SrcFile `json:"source_files"`

	// The total number of lines of code.
	LoC int64 `json:"loc"`
}

func newPackage(m map[string]interface{}) (*Package, error) {
	var err error
	errPrefix := "src/package"
	pkg := Package{}

	if pkg.Doc, err = extractStringSliceValue("doc", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if pkg.Name, err = extractStringValue("name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if pkg.Path, err = extractStringValue("path", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if pkg.LoC, err = extractInt64Value("loc", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	srcsMap, ok := m["source_files"]
	if !ok {
		return nil, addDebugInfo(errors.New(errPrefix + ": field 'source_files' does not exist"))
	}

	var s reflect.Value
	if s = reflect.ValueOf(srcsMap); s.Kind() != reflect.Slice {
		return nil, addDebugInfo(errors.New(
			errPrefix + ": field 'source_files' is supposed to be a slice"))
	}

	srcs := make([]*SrcFile, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		src := s.Index(i).Interface()

		switch src.(type) {
		case map[string]interface{}:
			if srcs[i], err = newSrcFile(src.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(errors.New(
				errPrefix + ": 'source_file' must be a map[string]interface{}"))
		}
	}

	pkg.SrcFiles = srcs

	return &pkg, nil
}

func newPackagesSlice(key, errPrefix string, m map[string]interface{}) ([]*Package, error) {
	var err error
	var s reflect.Value

	pkgsMap, ok := m[key]
	if !ok {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	}

	if s = reflect.ValueOf(pkgsMap); s.Kind() != reflect.Slice {
		return nil, addDebugInfo(fmt.Errorf(
			"%s: field '%s' is supposed to be a slice", errPrefix, key))
	}

	pkgs := make([]*Package, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		pkg := s.Index(i).Interface()

		switch pkg.(type) {
		case map[string]interface{}:
			if pkgs[i], err = newPackage(pkg.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}", errPrefix, key))
		}
	}

	return pkgs, nil
}

func mergePackage(p1, p2 *Package) (*Package, error) {
	if p1 == nil {
		return nil, addDebugInfo(errors.New("p1 cannot be nil"))
	}

	if p2 == nil {
		return nil, addDebugInfo(errors.New("p2 cannot be nil"))
	}

	var err error

	newPkg := new(Package)
	newPkg.Name = p1.Name
	newPkg.Path = p1.Path
	newPkg.Doc = p1.Doc

	if newPkg.SrcFiles, err = mergeSrcFilesSlices(p1.SrcFiles, p2.SrcFiles); err != nil {
		return nil, addDebugInfo(err)
	}

	newPkg.LoC = p1.LoC + p2.LoC

	return newPkg, nil
}

func mergePackageSlices(ps1, ps2 []*Package) ([]*Package, error) {
	if ps1 == nil {
		return nil, addDebugInfo(errors.New("ps1 cannot be nil"))
	}

	if ps2 == nil {
		return nil, addDebugInfo(errors.New("ps2 cannot be nil"))
	}

	newPkgs := make([]*Package, 0)

	var err error
	for _, p2 := range ps2 {
		var pkg *Package
		for _, p1 := range ps1 {
			if p1.Path == p2.Path {
				pkg = p1
				break
			}
		}

		if pkg == nil && p2.LoC > 0 {
			newPkgs = append(newPkgs, p2)
		} else {
			if pkg, err = mergePackage(pkg, p2); err != nil {
				return nil, addDebugInfo(err)
			}

			if pkg.LoC > 0 {
				newPkgs = append(newPkgs, pkg)
			}
		}
	}

	return newPkgs, nil
}
