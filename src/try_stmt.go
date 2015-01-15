// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"fmt"
	"reflect"
)

type TryStmt struct {
	StmtName     string         `json:"statement_name"`
	Body         []Stmt         `json:"body"`
	CatchClauses []*CatchClause `json:"catch_clauses,omitempty"`
	Finally      []Stmt         `json:"finally,omitempty"`
}

type CatchClause struct {
	Params []*Field `json:"params"`
	Body   []Stmt   `json:"body"`
}

func newTryStmt(m map[string]interface{}) (*TryStmt, error) {
	var err error
	errPrefix := "src/try_stmt"
	trystmt := TryStmt{}

	if typ, err := extractStringValue("statement_name", errPrefix, m); err != nil {
		// XXX It is not possible to add debug info on this error because it is
		// required that this error be en "errNotExist".
		return nil, errNotExist
	} else if typ != TryStmtName {
		return nil, fmt.Errorf("invalid type: expected 'TryStmt', found '%s'", typ)
	}

	trystmt.StmtName = TryStmtName

	if trystmt.Body, err = newStmtsSlice("body", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if trystmt.CatchClauses, err = newCatchClausesSlice("catch_clauses", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	if trystmt.Finally, err = newStmtsSlice("finally", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	}

	return &trystmt, nil
}

func newCatchClause(m map[string]interface{}) (*CatchClause, error) {
	var err error
	errPrefix := "src/catch_clause"
	catchclause := CatchClause{}

	if catchclause.Params, err = newFieldsSlice("parameters", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if catchclause.Body, err = newStmtsSlice("body", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	return &catchclause, nil
}

func newCatchClausesSlice(key, errPrefix string, m map[string]interface{}) ([]*CatchClause, error) {
	var err error
	var s *reflect.Value

	if s, err = reflectSliceValue("catch_clauses", errPrefix, m); err != nil {
		return nil, err
	}

	catchclauses := make([]*CatchClause, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		catchclause := s.Index(i).Interface()
		if catchclause == nil {
			continue
		}

		switch catchclause.(type) {
		case map[string]interface{}:
			if catchclauses[i], err = newCatchClause(catchclause.(map[string]interface{})); err != nil {
				return nil, addDebugInfo(err)
			}
		default:
			return nil, addDebugInfo(fmt.Errorf(
				"%s: '%s' must be a map[string]interface{}, found %v",
				errPrefix, key, reflect.TypeOf(catchclause)))
		}
	}

	return catchclauses, nil
}
