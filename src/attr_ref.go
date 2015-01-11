// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

type AttrRef struct {
	ExprName string `json:"expression_name"`
	Ident    *Ident `json:"ident"`
}

func newAttrRef(m map[string]interface{}) (*AttrRef, error) {
	var err error
	errPrefix := "src/attr_ref"
	attrref := AttrRef{}

	identMap, err := extractMapValue("ident", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if attrref.Ident, err = newIdent(identMap); err != nil {
		return nil, addDebugInfo(err)
	}

	return &attrref, nil
}
