// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "fmt"

const (
	IntLit    = "INT"
	FloatLit  = "FLOAT"
	ImagLit   = "IMAG"
	CharLit   = "CHAR"
	StringLit = "STRING"
	BoolLit   = "BOOl"
	NilLit    = "NIL"
)

type BasicLit struct {
	ExprName string `json:"expression_name"`
	Kind     string `json:"kind"`
	Value    string `json:"value"`
}

func newBasicLit(m map[string]interface{}) (*BasicLit, error) {
	var err error
	errPrefix := "src/basic_lit"
	basiclit := BasicLit{}

	if typ, err := extractStringValue("expression_name", errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	} else if typ != BasicLitName {
		return nil, fmt.Errorf("invalid type: expected 'BasicLit', found '%s'", typ)
	}

	basiclit.ExprName = BasicLitName

	if basiclit.Kind, err = extractStringValue("kind", errPrefix, m); err != nil {
		return nil, err
	}

	if basiclit.Value, err = extractStringValue("value", errPrefix, m); err != nil {
		return nil, err
	}

	return &basiclit, nil
}
