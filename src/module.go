// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

type Module struct {
	Name       string
	Attributes []*Attribute
	Methods    []*Method
	Modules    []*Module // For languages supporting mixins
}
