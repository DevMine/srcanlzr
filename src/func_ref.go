// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "errors"

type FuncRef struct {
	Namespace string `json:"namespace"`
	FuncName  string `json:"function_name"`
	External  bool   `json:"external"`
}

// NewFuncRef creates a new FuncRef from a generic map.
func NewFuncRef(m map[string]interface{}) (*FuncRef, error) {
	fctref := FuncRef{}

	var ok bool

	var namespace interface{}
	if namespace, ok = m["Namespace"]; !ok {
		return nil, errors.New("malformed FuncRef, no Namespace field")
	}

	var funcName interface{}
	if funcName, ok = m["FuncName"]; !ok {
		return nil, errors.New("malformed FuncRef, no FuncName field")
	}

	fctref.Namespace = namespace.(string)
	fctref.FuncName = funcName.(string)

	return &fctref, nil
}
