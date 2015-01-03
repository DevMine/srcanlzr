// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"errors"
	"fmt"
	"reflect"
)

type Interface struct {
	Name       string `json:"name"`
	Visibility string `json:"visibility"`
	// TODO
}

func newInterface(m map[string]interface{}) (*Interface, error) {
	var err error
	errPrefix := "src/constant"
	i := Interface{}

	if i.Name, err = extractStringValue("name", errPrefix, m); err != nil {
		return nil, err
	}

	if i.Visibility, err = extractStringValue("visibility", errPrefix, m); err != nil {
		return nil, err
	}

	return &i, nil
}

func newInterfacesSlice(key, errPrefix string, m map[string]interface{}) ([]*Interface, error) {
	var err error
	var s reflect.Value

	interfsMap, ok := m[key]
	if !ok {
		return nil, errors.New(fmt.Sprintf("%s: field '%s' does not exist", errPrefix, key))
	}

	if s = reflect.ValueOf(interfsMap); s.Kind() != reflect.Slice {
		return nil, errors.New(fmt.Sprintf("%s: field '%s' is supposed to be a slice",
			errPrefix, key))
	}

	interfs := make([]*Interface, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		interf := s.Index(i).Interface()

		switch interf.(type) {
		case map[string]interface{}:
			if interfs[i], err = newInterface(interf.(map[string]interface{})); err != nil {
				return nil, err
			}
		default:
			return nil, errors.New(fmt.Sprintf("%s: '%s' must be a map[string]interface{}",
				errPrefix, key))
		}
	}

	return interfs, nil
}
