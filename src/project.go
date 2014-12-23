// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "encoding/json"

// A project is the root of the src API and must be at the root of the JSON
// generated string.
//
// It contains the metadata of a project and the list of all packages (basically
// folders).
type Project struct {
	// The name of the project. Since it may be something really difficult to
	// guess, it should generally be the name of the folder containing the
	// project.
	Name string `json:"name"`

	// The repository in which the project is hosted.
	Repo *Repo `json:"repository"`

	// Programming languages used in the project.
	ProgLangs []Language `json:"languages"`

	// List of all packages of the project. A packages is just a folder
	// containing at least one source file.
	Packages []*Package `json:"packages"`

	// The total number of lines of code in the whole project.
	LoC int64 `json:"loc"`
}

// UnmarshalProject unmarshals a JSON representation of a Project into a real
//  Project structure.
//
// It is required to use this function instead of json.Unmarshal because we use
// an interface to abstract a Statement, thus json.Unmarshal is unable to
// unmarshal the statements correctly.
//
// TODO Find a more elegant way for solving this problem (eg. write a custom
// JSON parser).
func UnmarshalProject(bs []byte) (*Project, error) {
	p := &Project{}

	if err := json.Unmarshal(bs, p); err != nil {
		return nil, err
	}

	for _, pkgs := range p.Packages {
		for _, sfs := range pkgs.SourceFiles {
			for _, fct := range sfs.Functions {
				castStmts := make([]Statement, 0)

				for _, stmt := range fct.StmtList {
					castStmt, err := castToStatement(stmt.(map[string]interface{}))
					if err != nil {
						return nil, err
					}

					castStmts = append(castStmts, castStmt)
				}

				fct.StmtList = castStmts
			}

			for _, cls := range sfs.Classes {
				for _, mds := range cls.Methods {
					castStmts := make([]Statement, 0)

					for _, stmt := range mds.StmtList {
						castStmt, err := castToStatement(stmt.(map[string]interface{}))
						if err != nil {
							return nil, err
						}

						castStmts = append(castStmts, castStmt)
					}
				}
			}

			for _, mods := range sfs.Modules {
				for _, mds := range mods.Methods {
					castStmts := make([]Statement, 0)

					for _, stmt := range mds.StmtList {
						castStmt, err := castToStatement(stmt.(map[string]interface{}))
						if err != nil {
							return nil, err
						}

						castStmts = append(castStmts, castStmt)
					}
				}
			}
		}
	}

	return p, nil
}
