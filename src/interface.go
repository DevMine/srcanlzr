// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"fmt"
	"reflect"
)

type Interface struct {
	Doc                   []string        `json:"doc,omitempty"`
	Name                  string          `json:"name"`
	ImplementedInterfaces []*InterfaceRef `json:"implemented_interfaces,omitempty"`
	Protos                []*ProtoDecl    `json:"prototypes"`
	Visibility            string          `json:"visibility"`
}

func newInterface(m map[string]interface{}) (*Interface, error) {
	var err error
	errPrefix := "src/interface"
	i := Interface{}

	if i.Doc, err = extractStringSliceValue("doc", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if i.Name, err = extractStringValue("name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if i.ImplementedInterfaces, err = newInterfaceRefsSlice("implemented_interfaces", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if i.Protos, err = newProtoDeclsSlice("prototypes", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if i.Visibility, err = extractStringValue("visibility", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &i, nil
}

func newInterfacesSlice(key, errPrefix string, m map[string]interface{}) ([]*Interface, error) {
	var err error
	var s reflect.Value

	interfsMap, ok := m[key]
	if !ok {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	}

	if s = reflect.ValueOf(interfsMap); s.Kind() != reflect.Slice {
		return nil, addDebugInfo(fmt.Errorf(
			"%s: field '%s' is supposed to be a slice", errPrefix, key))
	}

	interfs := make([]*Interface, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		interf := s.Index(i).Interface()

		switch interf.(type) {
		case map[string]interface{}:
			if interfs[i], err = newInterface(interf.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}", errPrefix, key))
		}
	}

	return interfs, nil
}
