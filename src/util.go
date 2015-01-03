// Copyright 2014 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"errors"
	"fmt"
	"reflect"
)

func extractStringValue(key, errPrefix string, m map[string]interface{}) (string, error) {
	val, ok := m[key]
	if !ok {
		return "", errors.New(fmt.Sprintf("%s: field '%s' does not exist", errPrefix, key))
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
		return false, errors.New(fmt.Sprintf("%s: field '%s' does not exist", errPrefix, key))
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
		return 0.0, errors.New(fmt.Sprintf("%s: field '%s' does not exist", errPrefix, key))
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
		return 0, errors.New(fmt.Sprintf("%s: field '%s' does not exist", errPrefix, key))
	}

	switch val.(type) {
	case int64:
		return val.(int64), nil
	}

	return 0, errors.New(fmt.Sprintf(
		"%s: '%s' field is expected to be a int64, found %v",
		errPrefix, key, reflect.TypeOf(key)))
}

func extractMapValue(key, errPrefix string, m map[string]interface{}) (map[string]interface{}, error) {
	val, ok := m[key]
	if !ok {
		return nil, errors.New(fmt.Sprintf("%s: field '%s' does not exist", errPrefix, key))
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
	val, ok := m[key]
	if !ok {
		return []string{}, errors.New(fmt.Sprintf("%s: field '%s' does not exist", errPrefix, key))
	}

	switch val.(type) {
	case []string:
		return val.([]string), nil
	}

	return nil, errors.New(fmt.Sprintf(
		"%s: '%s' field is expected to be a slice of string, found %v",
		errPrefix, key, reflect.TypeOf(key)))
}

func reflectSliceValue(key, errPrefix string, m map[string]interface{}) (*reflect.Value, error) {
	val, ok := m[key]
	if !ok {
		return nil, errors.New(fmt.Sprintf("%s: field '%s' does not exist", errPrefix, key))
	}

	var s reflect.Value
	if s = reflect.ValueOf(val); s.Kind() != reflect.Slice {
		return nil, errors.New(errPrefix + ": field 'languages' is supposed to be a slice")
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
