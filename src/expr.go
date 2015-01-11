// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"errors"
	"fmt"
	"reflect"
)

const (
	BinaryExprName = "BINARY"
	CallExprName   = "CALL"
	UnaryExprName  = "UNARY"

	BasicLitName = "BASIC_LIT"
	FuncLitName  = "FUNC_LIT"
	ClassLitName = "CLASS_LIT"

	StructTypeName = "STRUCT_TYPE"

	AttrRefName = "ATTR_REF"

	ValueSpecName = "VALUE_SPEC"

	IdentName = "IDENT"
)

type Expr interface{}

func newExpr(m map[string]interface{}) (Expr, error) {
	errPrefix := "src/expr"

	typ, err := extractStringValue("expression_name", errPrefix, m)
	if err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	}

	switch typ {
	case BinaryExprName:
		return newBinaryExpr(m)
	case CallExprName:
		return newCallExpr(m)
	case UnaryExprName:
		return newUnaryExpr(m)
	case BasicLitName:
		return newBasicLit(m)
	case FuncLitName:
		return newFuncLit(m)
	case ClassLitName:
		return newClassLit(m)
	case ValueSpecName:
		return newValueSpec(m)
	case StructTypeName:
		return newStructType(m)
	}

	return nil, addDebugInfo(errors.New("unknown statement type"))
}

func newExprsSlice(key, errPrefix string, m map[string]interface{}) ([]Expr, error) {
	var err error
	var s *reflect.Value

	if s, err = reflectSliceValue(key, errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	exprs := make([]Expr, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		expr := s.Index(i).Interface()
		if expr == nil {
			continue
		}

		switch expr.(type) {
		case map[string]interface{}:
			if exprs[i], err = newExpr(expr.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}, found %v",
				errPrefix, key, reflect.TypeOf(expr)))
		}
	}

	return exprs, nil
}
