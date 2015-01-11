// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"fmt"
	"reflect"
)

// Method/Function prototype declaration
type ProtoDecl struct {
	Doc        []string  `json:"doc"`
	Ident      *Ident    `json:"ident"`
	Type       *FuncType `json:"type"`
	Visibility string    `json:"visibility"`
}

func newProtoDecl(m map[string]interface{}) (*ProtoDecl, error) {
	var err error
	errPrefix := "src/proto_decl"
	proto := ProtoDecl{}

	if proto.Doc, err = extractStringSliceValue("doc", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	identMap, err := extractMapValue("ident", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if proto.Ident, err = newIdent(identMap); err != nil {
		return nil, addDebugInfo(err)
	}

	fctTypeMap, err := extractMapValue("type", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if proto.Type, err = newFuncType(fctTypeMap); err != nil {
		return nil, addDebugInfo(err)
	}

	if proto.Visibility, err = extractStringValue("visibility", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &proto, nil
}

func newProtoDeclsSlice(key, errPrefix string, m map[string]interface{}) ([]*ProtoDecl, error) {
	var err error
	var s *reflect.Value

	if s, err = reflectSliceValue(key, errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	protos := make([]*ProtoDecl, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		proto := s.Index(i).Interface()

		switch proto.(type) {
		case map[string]interface{}:
			if protos[i], err = newProtoDecl(proto.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}", errPrefix, key))
		}
	}

	return protos, nil
}
