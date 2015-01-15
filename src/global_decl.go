// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"fmt"
	"reflect"
)

type GlobalDecl struct {
	DeclStmt
	Visibility string `json:"visibility"`
}

func newGlobalDecl(m map[string]interface{}) (*GlobalDecl, error) {
	var err error
	errPrefix := "src/global_decl"
	globaldecl := GlobalDecl{}

	if globaldecl.Lhs, err = newExprsSlice("left_hand_side", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if globaldecl.Rhs, err = newExprsSlice("right_hand_side", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if globaldecl.Kind, err = extractStringValue("kind", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if globaldecl.Visibility, err = extractStringValue("visibility", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &globaldecl, nil
}

func newGlobalDeclsSlice(key, errPrefix string, m map[string]interface{}) ([]*GlobalDecl, error) {
	var err error
	var s *reflect.Value

	if s, err = reflectSliceValue(key, errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	globaldecls := make([]*GlobalDecl, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		globaldecl := s.Index(i).Interface()
		if globaldecl == nil {
			continue
		}

		switch globaldecl.(type) {
		case map[string]interface{}:
			if globaldecls[i], err = newGlobalDecl(globaldecl.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}, found %v",
				errPrefix, key, reflect.TypeOf(globaldecl)))
		}
	}

	return globaldecls, nil
}
