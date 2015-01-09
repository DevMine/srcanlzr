// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

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
	Repo *Repo `json:"repository,omitempty"`

	// Programming languages used in the project.
	ProgLangs []*Language `json:"languages"`

	// List of all packages of the project. A packages is just a folder
	// containing at least one source file.
	Packages []*Package `json:"packages"`

	// The total number of lines of code in the whole project.
	LoC int64 `json:"loc"`
}

// newProject creates a new Project from a generic map.
func newProject(m map[string]interface{}) (*Project, error) {
	var err error
	errPrefix := "src/project"
	prj := Project{}

	if prj.Name, err = extractStringValue("name", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if prj.LoC, err = extractInt64Value("loc", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if prj.Packages, err = newPackagesSlice("packages", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if prj.ProgLangs, err = newLanguagesSlice("languages", errPrefix, m); err != nil {
		return nil, addDebugInfo(err)
	}

	if repoMap, err := extractMapValue("repository", errPrefix, m); err != nil && isExist(err) {
		return nil, addDebugInfo(err)
	} else if err == nil {
		if prj.Repo, err = newRepo(repoMap); err != nil && isExist(err) {
			return nil, addDebugInfo(err)
		}
	}

	return &prj, nil
}
