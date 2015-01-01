// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

// A package is a folder contaning at least one source file.
type Package struct {
	// The name of the pacakge (the folder name)
	Name string `json:"name"`

	// The full path of the package. The path must be relative to the root of
	// the project and never be an absolute path.
	Path string `json:"path"`

	// The package documentation.
	Doc string `json:"doc"`

	// The list of all source files contained in the package.
	SourceFiles []*SourceFile `json:"source_files"`

	// The total number of lines of code.
	LoC int64 `json:"loc"`
}
