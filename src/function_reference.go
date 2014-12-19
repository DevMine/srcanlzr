// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "errors"

type FuncRef struct {
	Namespace string
	FuncName  string
	External  bool
}

// CastToFuncRef "cast" a generic map into a FuncRef.
func CastToFuncRef(m map[string]interface{}) (*FuncRef, error) {
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
