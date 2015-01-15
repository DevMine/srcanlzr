// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

type AttrRef struct {
	ExprName string `json:"expression_name"`
	Name     *Ident `json:"name"`
}

func newAttrRef(m map[string]interface{}) (*AttrRef, error) {
	var err error
	errPrefix := "src/attr_ref"
	attrref := AttrRef{}

	identMap, err := extractMapValue("name", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if attrref.Name, err = newIdent(identMap); err != nil {
		return nil, addDebugInfo(err)
	}

	return &attrref, nil
}
