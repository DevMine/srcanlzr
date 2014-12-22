// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

type Attribute struct {
	Variable
	Visibility string
	Constant   bool
	Static     bool
}
