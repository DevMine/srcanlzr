// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"fmt"
	"reflect"
)

type Class struct {
	Name                  string             `json:"name"`
	Visibility            string             `json:"visibility"`
	ExtendedClasses       []*Class           `json:"extended_classes"`
	ImplementedInterfaces []*Interface       `json:"implemented_interfaces"`
	Attrs                 []*Attr            `json:"attributes"`
	Constructors          []*ConstructorDecl `json:"constructors"`
	Methods               []*Method          `json:"methods"`
	Traits                []*Trait           `json:"traits"`
}

func newClass(m map[string]interface{}) (*Class, error) {
	var err error
	errPrefix := "src/class"
	cls := Class{}

	if cls.Name, err = extractStringValue("name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if cls.Visibility, err = extractStringValue("visibility", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if cls.ExtendedClasses, err = newClassesSlice("classes", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if cls.ImplementedInterfaces, err = newInterfacesSlice("interfaces", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if cls.Attrs, err = newAttrsSlice("attributes", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if cls.Constructors, err = newConstructorDeclsSlice("constructors", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if cls.Methods, err = newMethodsSlice("methods", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &cls, nil
}

func newClassesSlice(key, errPrefix string, m map[string]interface{}) ([]*Class, error) {
	var err error
	var s reflect.Value

	clssMap, ok := m[key]
	if !ok {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	}

	if s = reflect.ValueOf(clssMap); s.Kind() != reflect.Slice {
		return nil, addDebugInfo(fmt.Errorf(
			"%s: field '%s' is supposed to be a slice", errPrefix, key))
	}

	clss := make([]*Class, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		cls := s.Index(i).Interface()

		switch cls.(type) {
		case map[string]interface{}:
			if clss[i], err = newClass(cls.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}", errPrefix, key))
		}
	}

	return clss, nil
}
