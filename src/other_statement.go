// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"errors"
	"fmt"
)

type OtherStatement struct {
	Type     string      `json:"type"`
	StmtList []Statement `json:"statements_list"`
	Line     int64       `json:"line"` // Line number of the statement relatively to the function.
}

// newOtherStatement creates a new OtherStatement from a generic map.
func newOtherStatement(m map[string]interface{}) (*OtherStatement, error) {
	var err error
	errPrefix := "src/other_statement"
	otherstmt := OtherStatement{}

	// should never happen
	if typ, ok := m["type"]; !ok || typ != OtherStmtName {
		return nil, errors.New(fmt.Sprintf("%s: the generic map supplied is not a OtherStatement",
			errPrefix))
	}

	if otherstmt.Type, err = extractStringValue("type", errPrefix, m); err != nil {
		return nil, err
	}

	if otherstmt.Line, err = extractInt64Value("line", errPrefix, m); err != nil {
		return nil, err
	}

	if otherstmt.StmtList, err = newStatementsSlice("statements_list", errPrefix, m); err != nil {
		return nil, err
	}

	return &otherstmt, nil
}
