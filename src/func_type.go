// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

type FuncType struct {
	Params  []*Field `json:"parameters,omitempty"`
	Results []*Field `json:"results,omitempty"`
}

func newFuncType(m map[string]interface{}) (*FuncType, error) {
	var err error
	errPrefix := "src/func_type"
	fct := FuncType{}

	if fct.Params, err = newFieldsSlice("parameters", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if fct.Results, err = newFieldsSlice("results", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	return &fct, nil
}
