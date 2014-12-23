// Copyright 2014 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

type Package struct {
	Name        string        `json:"name"`         // Package name
	Path        string        `json:"path"`         // Package location
	Doc         string        `json:"doc"`          // Package doc comments
	SourceFiles []*SourceFile `json:"source_files"` // Source files
	LoC         int64         `json:"loc"`          // Lines of Code
}
