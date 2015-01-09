// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"encoding/xml"
	"errors"
	"fmt"
	"reflect"
)

type Language struct {
	XMLName xml.Name `json:"-" xml:"languages"`

	// TODO rename into name
	Lang string `json:"language" xml:"language"`

	Paradigms []string `json:"paradigms" xml:"paradigm"`
}

func newLanguage(m map[string]interface{}) (*Language, error) {
	var err error
	errPrefix := "src/language"
	lang := Language{}

	if lang.Lang, err = extractStringValue("language", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if lang.Paradigms, err = extractStringSliceValue("paradigms", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &lang, nil
}

func newLanguagesSlice(key, errPrefix string, m map[string]interface{}) ([]*Language, error) {
	var err error
	var s reflect.Value

	langsMap, ok := m[key]
	if !ok {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	}

	if s = reflect.ValueOf(langsMap); s.Kind() != reflect.Slice {
		return nil, addDebugInfo(errors.New(fmt.Sprintf(
			"%s: field '%s' is supposed to be a slice", errPrefix, key)))
	}

	langs := make([]*Language, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		lang := s.Index(i).Interface()

		switch lang.(type) {
		case map[string]interface{}:
			if langs[i], err = newLanguage(lang.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(errors.New(fmt.Sprintf(
				"%s: '%s' must be a map[string]interface{}", errPrefix, key)))
		}
	}

	return langs, nil
}

func mergeLanguageSlices(ls1, ls2 []*Language) ([]*Language, error) {
	if ls1 == nil {
		return nil, addDebugInfo(errors.New("ls1 cannot be nil"))
	}

	if ls2 == nil {
		return nil, addDebugInfo(errors.New("ls2 cannot be nil"))
	}

	newLangs := make([]*Language, 0)
	newLangs = append(newLangs, ls1...)

	for _, l2 := range ls2 {
		var found bool
		for _, l1 := range ls1 {
			if l1.Lang == l2.Lang {
				found = true
				break
			}
		}

		if !found {
			newLangs = append(newLangs, l2)
		}
	}

	return newLangs, nil
}
