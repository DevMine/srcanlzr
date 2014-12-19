// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

type Package struct {
	Name        string        // Package name
	Path        string        // Package location
	Doc         string        // Package doc comments
	SourceFiles []*SourceFile // Source files
	LoC         int64         // Lines of Code
}
