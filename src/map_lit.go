// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"fmt"
	"reflect"
)

type MapLit struct {
	Type *MapType        `json:"type"`
	Elts []*KeyValuePair `json:"elements"`
}

type KeyValuePair struct {
	Key   Expr `json:"key"`
	Value Expr `json:"value"`
}

func newMapLit(m map[string]interface{}) (*MapLit, error) {
	var err error
	errPrefix := "src/map_lit"
	maplit := MapLit{}

	typeMap, err := extractMapValue("type", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if maplit.Type, err = newMapType(typeMap); err != nil {
		return nil, addDebugInfo(err)
	}

	if maplit.Elts, err = newKeyValuePairsSlice("elements", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &maplit, nil
}

func newKeyValuePair(m map[string]interface{}) (*KeyValuePair, error) {
	var err error
	errPrefix := "src/key_value_pair"
	kvpair := KeyValuePair{}

	keyMap, err := extractMapValue("key", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if kvpair.Key, err = newArrayType(keyMap); err != nil {
		return nil, addDebugInfo(err)
	}

	valueMap, err := extractMapValue("value", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if kvpair.Value, err = newArrayType(valueMap); err != nil {
		return nil, addDebugInfo(err)
	}

	return &kvpair, nil
}

func newKeyValuePairsSlice(key, errPrefix string, m map[string]interface{}) ([]*KeyValuePair, error) {
	var err error
	var s reflect.Value

	kvpairsMap, ok := m[key]
	if !ok || kvpairsMap == nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	}

	if s = reflect.ValueOf(kvpairsMap); s.Kind() != reflect.Slice {
		return nil, addDebugInfo(fmt.Errorf(
			"%s: field '%s' is supposed to be a slice", errPrefix, key))
	}

	kvpairs := make([]*KeyValuePair, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		kvpair := s.Index(i).Interface()
		if kvpair == nil {
			continue
		}

		switch kvpair.(type) {
		case map[string]interface{}:
			if kvpairs[i], err = newKeyValuePair(kvpair.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}, found %v",
				errPrefix, key, reflect.TypeOf(kvpair)))
		}
	}

	return kvpairs, nil
}
