// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

type FuncRef struct {
	Namespace string `json:"namespace"`
	FuncName  string `json:"function_name"`
	External  bool   `json:"external"`
}

// newFuncRef creates a new FuncRef from a generic map.
func newFuncRef(m map[string]interface{}) (*FuncRef, error) {
	var err error
	errPrefix := "src/func_ref"
	fctref := FuncRef{}

	if fctref.Namespace, err = extractStringValue("namespace", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if fctref.FuncName, err = extractStringValue("function_name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if fctref.External, err = extractBoolValue("external", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &fctref, nil
}
