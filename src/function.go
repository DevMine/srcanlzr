// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

// TODO it would be nice to only have the raw function prototype instead of the
// whole function.
// TODO proposal: keep only positions instead of a raw string. This would make
// FIXME add the possibility to have multiple return statements
// the parsing faster and the generated JSON smaller.
type Function struct {
	Name     string
	Comments string
	Args     []Variable
	Return   []Variable
	StmtList []Statement
	LoC      int64  // Lines of Code
	Raw      string // Function raw source code.
}
