// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "encoding/xml"

type Language struct {
	XMLName   xml.Name `json:"-" xml:"languages"`
	Lang      int      `json:"language" xml:"language"`
	Paradigms []int    `json:"paradigms" xml:"paradigm"`
}
