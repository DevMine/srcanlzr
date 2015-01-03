// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"errors"
	"fmt"
	"reflect"
)

type Method struct {
	Function
	Visibility string `json:"visibility"`
}

func newMethod(m map[string]interface{}) (*Method, error) {
	var err error
	errPrefix := "src/method"
	mthd := Method{}

	var fct *Function
	if fct, err = newFunction(m); err != nil {
		return nil, err
	}
	mthd.Function = *fct

	if mthd.Visibility, err = extractStringValue("visibility", errPrefix, m); err != nil {
		return nil, err
	}

	return &mthd, nil
}

func newMethodsSlice(key, errPrefix string, m map[string]interface{}) ([]*Method, error) {
	var err error
	var s reflect.Value

	mthdsMap, ok := m[key]
	if !ok {
		return nil, errors.New(fmt.Sprintf("%s: field '%s' does not exist", errPrefix, key))
	}

	if s = reflect.ValueOf(mthdsMap); s.Kind() != reflect.Slice {
		return nil, errors.New(fmt.Sprintf("%s: field '%s' is supposed to be a slice",
			errPrefix, key))
	}

	mthds := make([]*Method, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		mthd := s.Index(i).Interface()

		switch mthd.(type) {
		case map[string]interface{}:
			if mthds[i], err = newMethod(mthd.(map[string]interface{})); err != nil {
				return nil, err
			}
		default:
			return nil, errors.New(fmt.Sprintf("%s: '%s' must be a map[string]interface{}",
				errPrefix, key))
		}
	}

	return mthds, nil
}
