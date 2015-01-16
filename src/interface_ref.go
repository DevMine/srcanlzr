// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"fmt"
	"reflect"
)

type InterfaceRef struct {
	Namespace     string `json:"namespace"`
	InterfaceName string `json:"interface_name"`
}

func newInterfaceRef(m map[string]interface{}) (*InterfaceRef, error) {
	var err error
	errPrefix := "src/interface_ref"
	iref := InterfaceRef{}

	if iref.Namespace, err = extractStringValue("namespace", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if iref.InterfaceName, err = extractStringValue("interface_name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &iref, nil
}

func newInterfaceRefsSlice(key, errPrefix string, m map[string]interface{}) ([]*InterfaceRef, error) {
	var err error
	var s *reflect.Value

	if s, err = reflectSliceValue(key, errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, err
	}

	irefs := make([]*InterfaceRef, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		iref := s.Index(i).Interface()
		if iref == nil {
			continue
		}

		switch iref.(type) {
		case map[string]interface{}:
			if irefs[i], err = newInterfaceRef(iref.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}, found %v",
				errPrefix, key, reflect.TypeOf(iref)))
		}
	}

	return irefs, nil
}
