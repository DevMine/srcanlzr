// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

type Constant struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
	Doc   string `json:"doc"`
}
