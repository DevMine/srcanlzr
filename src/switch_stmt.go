// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"fmt"
	"reflect"
)

type SwitchStmt struct {
	StmtName    string        `json:"statement_name"`
	Init        Expr          `json:"initialization,omitempty"`
	Cond        Expr          `json:"condition"`
	CaseClauses []*CaseClause `json:"case_clauses,omitempty"`
	Default     []Stmt        `json:"default,omitempty"`
}

type CaseClause struct {
	Conds []Expr `json:"conditions,omitempty"`
	Body  []Stmt `json:"body,omitempty"`
}

func newSwitchStmt(m map[string]interface{}) (*SwitchStmt, error) {
	var err error
	errPrefix := "src/switch_stmt"
	switchstmt := SwitchStmt{}

	if typ, err := extractStringValue("statement_name", errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	} else if typ != SwitchStmtName {
		return nil, fmt.Errorf("invalid type: expected 'SwitchStmt', found '%s'", typ)
	}

	switchstmt.StmtName = SwitchStmtName

	initMap, err := extractMapValue("initialization", errPrefix, m)
	if err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	} else if err == nil {
		if switchstmt.Cond, err = newStmt(initMap); err != nil {
			return nil, addDebugInfo(err)
		}
	}

	condMap, err := extractMapValue("condition", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	if switchstmt.Cond, err = newExpr(condMap); err != nil {
		return nil, addDebugInfo(err)
	}

	if switchstmt.CaseClauses, err = newCaseClausesSlice("case_clauses", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if switchstmt.Default, err = newStmtsSlice("default", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	return &switchstmt, nil
}

func newCaseClause(m map[string]interface{}) (*CaseClause, error) {
	var err error
	errPrefix := "src/case_clause"
	caseclause := CaseClause{}

	if caseclause.Conds, err = newExprsSlice("conditions", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if caseclause.Body, err = newStmtsSlice("body", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	return &caseclause, nil
}

func newCaseClausesSlice(key, errPrefix string, m map[string]interface{}) ([]*CaseClause, error) {
	var err error
	var s *reflect.Value

	if s, err = reflectSliceValue(key, errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, err
	}

	ccs := make([]*CaseClause, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		cc := s.Index(i).Interface()
		if cc == nil {
			continue
		}

		switch cc.(type) {
		case map[string]interface{}:
			if ccs[i], err = newCaseClause(cc.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}, found %v",
				errPrefix, key, reflect.TypeOf(cc)))
		}
	}

	return ccs, nil
}
