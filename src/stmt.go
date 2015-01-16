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
	IfStmtName        = "IF"
	SwitchStmtName    = "SWITCH"
	LoopStmtName      = "LOOP"
	RangeLoopStmtName = "RANGE_LOOP"
	AssignStmtName    = "ASSIGN"
	DeclStmtName      = "DECL"
	ReturnStmtName    = "RETURN"
	ExprStmtName      = "EXPR"
	TryStmtName       = "TRY"
	ThrowStmtName     = "THROW"
	OtherStmtName     = "OTHER"
)

type Stmt interface{}

func newStmt(m map[string]interface{}) (Stmt, error) {
	errPrefix := "src/stmt"

	typ, err := extractStringValue("statement_name", errPrefix, m)
	if err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	}

	switch typ {
	case IfStmtName:
		return newIfStmt(m)
	case SwitchStmtName:
		return newSwitchStmt(m)
	case LoopStmtName:
		return newLoopStmt(m)
	case RangeLoopStmtName:
		return newRangeLoopStmt(m)
	case AssignStmtName:
		return newAssignStmt(m)
	case DeclStmtName:
		return newDeclStmt(m)
	case ReturnStmtName:
		return newReturnStmt(m)
	case ExprStmtName:
		return newExprStmt(m)
	case TryStmtName:
		return newTryStmt(m)
	case ThrowStmtName:
		return newThrowStmt(m)
	case OtherStmtName:
		return newOtherStmt(m)
	}

	return nil, addDebugInfo(errors.New("unknown statement type"))
}

func newStmtsSlice(key, errPrefix string, m map[string]interface{}) ([]Stmt, error) {
	var err error
	var s *reflect.Value

	if s, err = reflectSliceValue(key, errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, err
	}

	stmts := make([]Stmt, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		stmt := s.Index(i).Interface()
		if stmt == nil {
			continue
		}

		switch stmt.(type) {
		case map[string]interface{}:
			if stmts[i], err = newStmt(stmt.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		case []interface{}:
			// XXX investigate why this case happen?
			continue
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}, found %v",
				errPrefix, key, reflect.TypeOf(stmt)))
		}
	}

	return stmts, nil
}
