// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"fmt"
	"reflect"
)

type ClassDecl struct {
	Doc                   []string           `json:"doc,omitempty"`
	Name                  string             `json:"name"`
	Visibility            string             `json:"visibility"`
	ExtendedClasses       []*ClassRef        `json:"extended_classes,omitempty"`
	ImplementedInterfaces []*InterfaceRef    `json:"implemented_interfaces,omitempty"`
	Attrs                 []*Attr            `json:"attributes,omitempty"`
	Constructors          []*ConstructorDecl `json:"constructors,omitempty"`
	Destructors           []*DestructorDecl  `json:"destructors,omitempty"`
	Methods               []*MethodDecl      `json:"methods,omitempty"`
	NestedClasses         []*ClassDecl       `json:"nested_classes,omitempty"`
	Mixins                []*TraitRef        `json:"mixins,omitempty"`
}

func newClassDecl(m map[string]interface{}) (*ClassDecl, error) {
	var err error
	errPrefix := "src/class_decl"
	cls := ClassDecl{}

	if cls.Doc, err = extractStringSliceValue("doc", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if cls.Name, err = extractStringValue("name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if cls.Visibility, err = extractStringValue("visibility", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if cls.ExtendedClasses, err = newClassRefsSlice("extended_classes", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if cls.ImplementedInterfaces, err = newInterfaceRefsSlice("implemented_interfaces", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if cls.Attrs, err = newAttrsSlice("attributes", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if cls.Constructors, err = newConstructorDeclsSlice("constructors", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if cls.Destructors, err = newDestructorDeclsSlice("destructors", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if cls.Methods, err = newMethodDeclsSlice("methods", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if cls.NestedClasses, err = newClassDeclsSlice("nested_classes", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if cls.Mixins, err = newTraitRefsSlice("mixins", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	return &cls, nil
}

func newClassDeclsSlice(key, errPrefix string, m map[string]interface{}) ([]*ClassDecl, error) {
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

	clss := make([]*ClassDecl, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		cls := s.Index(i).Interface()

		switch cls.(type) {
		case map[string]interface{}:
			if clss[i], err = newClassDecl(cls.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}", errPrefix, key))
		}
	}

	return clss, nil
}
