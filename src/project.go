// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import "errors"

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
	Langs []*Language `json:"languages"`

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

	if prj.Langs, err = newLanguagesSlice("languages", errPrefix, m); err != nil {
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

// mergeProjects merges projects p1 and p2 into p1
func mergeProjects(p1, p2 *Project) (*Project, error) {
	if p1 == nil {
		return nil, addDebugInfo(errors.New("p1 cannot be nil"))
	}

	if p2 == nil {
		return nil, addDebugInfo(errors.New("p2 cannot be nil"))
	}

	newPrj := new(Project)
	newPrj.Name = p1.Name

	var err error

	if newPrj.Langs, err = mergeLanguageSlices(p1.Langs, p2.Langs); err != nil {
		return nil, addDebugInfo(err)
	}

	if newPrj.Packages, err = mergePackageSlices(p1.Packages, p2.Packages); err != nil {
		return nil, addDebugInfo(err)
	}

	newPrj.LoC += p1.LoC + p2.LoC

	return newPrj, nil
}
