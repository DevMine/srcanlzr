// Copyright 2014-2015 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"io"
	"os"
)

// Decode a JSON encoded src.Project read from r.
func Decode(r io.Reader) (*Project, error) {
	dec := newDecoder(r)
	return dec.decode()
}

// Decode a JSON encoded src.Project read from a given file.
func DecodeFile(path string) (*Project, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Decode(f)
}

// MergeAll merges a list of projects.
//
// There must be at least one project. In this case, it just returns a copy of
// the project.
//
// The merge only performs shallow copies, which means that if the field value
// is a pointer it copies the memory address and not the value pointed.
/*func MergeAll(ps ...*Project) (*Project, error) {
	if len(ps) == 0 {
		return nil, addDebugInfo(errors.New("p cannot be nil"))
	}

	newPrj := &Project{
		Name:     ps[0].Name,
		Langs:    ps[0].Langs,
		Packages: ps[0].Packages,
		LoC:      ps[0].LoC,
	}

	if len(ps) == 1 {
		return newPrj, nil
	}

	var err error
	for i := 1; i < len(ps); i++ {
		curr := ps[i]
		if curr == nil {
			return nil, addDebugInfo(fmt.Errorf("p[%d] is nil", i))
		}

		if newPrj, err = Merge(newPrj, curr); err != nil {
			return nil, addDebugInfo(err)
		}
	}

	return newPrj, nil
}

// Merge merges two project. See MergeAll for more details.
func Merge(p1, p2 *Project) (*Project, error) {
	return mergeProjects(p1, p2)
}*/
