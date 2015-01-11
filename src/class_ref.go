// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"fmt"
	"reflect"
)

type ClassRef struct {
	Namespace string `json:"namespace"`
	ClassName string `json:"class_name"`
}

func newClassRef(m map[string]interface{}) (*ClassRef, error) {
	var err error
	errPrefix := "src/class_ref"
	clsref := ClassRef{}

	if clsref.Namespace, err = extractStringValue("namespace", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if clsref.ClassName, err = extractStringValue("class_name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &clsref, nil
}

func newClassRefsSlice(key, errPrefix string, m map[string]interface{}) ([]*ClassRef, error) {
	var err error
	var s *reflect.Value

	if s, err = reflectSliceValue(key, errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	clsrefs := make([]*ClassRef, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		clsref := s.Index(i).Interface()
		if clsref == nil {
			continue
		}

		switch clsref.(type) {
		case map[string]interface{}:
			if clsrefs[i], err = newClassRef(clsref.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}, found %v",
				errPrefix, key, reflect.TypeOf(clsref)))
		}
	}

	return clsrefs, nil
}
