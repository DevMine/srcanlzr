// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"fmt"
	"reflect"
)

type TraitRef struct {
	Namespace string `json:"namespace"`
	TraitName string `json:"trait_name"`
}

func newTraitRef(m map[string]interface{}) (*TraitRef, error) {
	var err error
	errPrefix := "src/trait_ref"
	traitref := TraitRef{}

	if traitref.Namespace, err = extractStringValue("namespace", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if traitref.TraitName, err = extractStringValue("trait_name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &traitref, nil
}

func newTraitRefsSlice(key, errPrefix string, m map[string]interface{}) ([]*TraitRef, error) {
	var err error
	var s *reflect.Value

	if s, err = reflectSliceValue(key, errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, err
	}

	traitrefs := make([]*TraitRef, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		traitref := s.Index(i).Interface()
		if traitref == nil {
			continue
		}

		switch traitref.(type) {
		case map[string]interface{}:
			if traitrefs[i], err = newTraitRef(traitref.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}, found %v",
				errPrefix, key, reflect.TypeOf(traitref)))
		}
	}

	return traitrefs, nil
}
