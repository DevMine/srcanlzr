// Copyright 2014 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "encoding/xml"

type Language struct {
	XMLName   xml.Name `json:"-" xml:"languages"`
	Lang      string   `json:"language" xml:"language"`
	Paradigms []string `json:"paradigms" xml:"paradigm"`
}
