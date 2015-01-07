// Copyright 2014 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	errNotExist = errors.New("the field does not exist")
)

func isNotExist(err error) bool {
	return err == errNotExist
}

func isExist(err error) bool {
	return err != errNotExist
}

func extractStringValue(key, errPrefix string, m map[string]interface{}) (string, error) {
	val, ok := m[key]
	if !ok {
		return "", errNotExist
	}

	switch val.(type) {
	case string:
		return val.(string), nil
	}

	return "", errors.New(fmt.Sprintf(
		"%s: '%s' field is expected to be a string, found %v",
		errPrefix, key, reflect.TypeOf(key)))
}

func extractBoolValue(key, errPrefix string, m map[string]interface{}) (bool, error) {
	val, ok := m[key]
	if !ok {
		return false, errNotExist
	}

	switch val.(type) {
	case bool:
		return val.(bool), nil
	}

	return false, errors.New(fmt.Sprintf(
		"%s: '%s' field is expected to be a bool, found %v",
		errPrefix, key, reflect.TypeOf(key)))
}

func extractFloat64Value(key, errPrefix string, m map[string]interface{}) (float64, error) {
	val, ok := m[key]
	if !ok {
		return 0.0, errNotExist
	}

	switch val.(type) {
	case float64:
		return val.(float64), nil
	}

	return 0.0, errors.New(fmt.Sprintf(
		"%s: '%s' field is expected to be a float64, found %v",
		errPrefix, key, reflect.TypeOf(key)))
}

func extractInt64Value(key, errPrefix string, m map[string]interface{}) (int64, error) {
	val, ok := m[key]
	if !ok {
		return 0, errNotExist
	}

	switch val.(type) {
	case int64:
		return val.(int64), nil
	case float64:
		fl := val.(float64)
		// XXX make a safe cast...
		return int64(fl), nil
	}

	return 0, errors.New(fmt.Sprintf(
		"%s: '%s' field is expected to be a int64, found %v",
		errPrefix, key, reflect.TypeOf(key)))
}

func extractMapValue(key, errPrefix string, m map[string]interface{}) (map[string]interface{}, error) {
	val, ok := m[key]
	if !ok {
		return nil, errNotExist
	}

	switch val.(type) {
	case map[string]interface{}:
		return val.(map[string]interface{}), nil
	}

	return nil, errors.New(fmt.Sprintf(
		"%s: '%s' field is expected to be a generic map, found %v",
		errPrefix, key, reflect.TypeOf(key)))
}

func extractStringSliceValue(key, errPrefix string, m map[string]interface{}) ([]string, error) {
	s, err := reflectSliceValue(key, errPrefix, m)
	if err != nil {
		return nil, err
	}

	ss := make([]string, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		val := s.Index(i).Interface()

		switch val.(type) {
		case string:
			ss = append(ss, val.(string))
		default:
			return nil, errors.New(fmt.Sprintf("%s: '%s' must be a []string", errPrefix, key))
		}
	}

	return ss, nil
	/*return nil, errors.New(fmt.Sprintf(
	"%s: '%s' field is expected to be a slice of string, found %v",
	errPrefix, key, reflect.TypeOf(key)))*/
}

func reflectSliceValue(key, errPrefix string, m map[string]interface{}) (*reflect.Value, error) {
	val, ok := m[key]
	if !ok {
		return nil, errNotExist
	}

	var s reflect.Value
	if s = reflect.ValueOf(val); s.Kind() != reflect.Slice {
		return nil, errors.New(fmt.Sprintf("%s: field '%s' is supposed to be a slice", errPrefix, key))
	}

	return &s, nil
}

func castToMap(key, errPrefix string, val interface{}) (map[string]interface{}, error) {
	switch val.(type) {
	case map[string]interface{}:
		return val.(map[string]interface{}), nil
	}

	return nil, errors.New(
		fmt.Sprintf("%s: '%s' must be a map[string]interface{}", errPrefix, key))
}
