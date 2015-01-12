// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"fmt"
	"reflect"
)

type Ident struct {
	ExprName string `json:"expression_name"`
	Name     string `json:"name"`
}

func newIdent(m map[string]interface{}) (*Ident, error) {
	var err error
	errPrefix := "src/ident"
	ident := Ident{}

	if typ, err := extractStringValue("expression_name", errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	} else if typ != IdentName {
		return nil, fmt.Errorf("invalid type: expected 'Ident', found '%s'", typ)
	}

	ident.ExprName = IdentName

	if ident.Name, err = extractStringValue("name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &ident, nil
}

func newIdentsSlice(key, errPrefix string, m map[string]interface{}) ([]*Ident, error) {
	var err error
	var s *reflect.Value

	if s, err = reflectSliceValue(key, errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	idents := make([]*Ident, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		ident := s.Index(i).Interface()
		if ident == nil {
			continue
		}

		switch ident.(type) {
		case map[string]interface{}:
			if idents[i], err = newIdent(ident.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}, found %v",
				errPrefix, key, reflect.TypeOf(ident)))
		}
	}

	return idents, nil
}
