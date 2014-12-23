// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

type Module struct {
	Name       string       `json:"name"`
	Attributes []*Attribute `json:"attributes"`
	Methods    []*Method    `json:"methods"`
	Modules    []*Module    `json:"modules"` // For languages supporting mixins
}
