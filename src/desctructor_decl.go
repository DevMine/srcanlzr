// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"fmt"
	"reflect"
)

type DestructorDecl struct {
	ConstructorDecl
}

func newDestructorDecl(m map[string]interface{}) (*DestructorDecl, error) {
	var err error
	errPrefix := "src/destructor_decl"
	destruc := DestructorDecl{}

	if destruc.Doc, err = extractStringSliceValue("doc", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if destruc.Name, err = extractStringValue("name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if destruc.Params, err = newFieldsSlice("parameters", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if destruc.Body, err = newStmtsSlice("body", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if destruc.Visibility, err = extractStringValue("visibility", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if destruc.LoC, err = extractInt64Value("loc", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &destruc, nil
}

func newDestructorDeclsSlice(key, errPrefix string, m map[string]interface{}) ([]*DestructorDecl, error) {
	var err error
	var s *reflect.Value

	if s, err = reflectSliceValue(key, errPrefix, m); err != nil {
		return nil, err
	}

	destrucs := make([]*DestructorDecl, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		destruc := s.Index(i).Interface()

		switch destruc.(type) {
		case map[string]interface{}:
			if destrucs[i], err = newDestructorDecl(destruc.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}", errPrefix, key))
		}
	}

	return destrucs, nil
}
