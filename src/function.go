// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

// TODO it would be nice to only have the raw function prototype instead of the
// whole function.
// TODO proposal: keep only positions instead of a raw string. This would make
// FIXME add the possibility to have multiple return statements
// the parsing faster and the generated JSON smaller.
type Function struct {
	Name     string      `json:"name"`
	Doc      string      `json:"doc"` // TODO rename into doc?
	Args     []Variable  `json:"args"`
	Return   []Variable  `json:"return"`
	StmtList []Statement `json:"statements_list"`
	LoC      int64       `json:"loc"` // Lines of Code
	Raw      string      `json:"raw"` // Function raw source code.
}
