// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"encoding/json"
	"errors"

	"github.com/DevMine/repotool/model"
)

// Project is the root of the src API and therefore it must be at the root of
// the JSON.
//
// It contains the metadata of a project and the list of all packages.
type Project struct {
	// The name of the project. Since it may be something really difficult to
	// guess, it should generally be the name of the folder containing the
	// project.
	Name string `json:"name"`

	// The repository in which the project is hosted, or nil. This field is not
	// meant to be filled by one of the language parsers. Only repotool should
	// take care of it. For more details, see:
	//    https://github.com/DevMine/repotool
	//
	// Since this field uses an external type, it is not unmarshalled by
	// src.Unmarshal itself but by the standard json.Unmarshal function.
	// To do so, its unmarshalling is defered using json.RawMessage.
	// See the RepoRaw field.
	Repo *model.Repository `json:"-"`

	// RepoRaw is only used to defer the unmarshalling of a model.Repository.
	RepoRaw json.RawMessage `json:"repository,omitempty"`

	// The list of all programming languages used by the project. Each language
	// must be added by the corresponding language parsers if and only if the
	// project contains at least one line of code written in this language.
	Langs []*Language `json:"languages"`

	// List of all packages of the project. We call "package" every folder that
	// contains at least one source file.
	Packages []*Package `json:"packages"`

	// The total number of lines of code in the whole project, independently of
	// the language.
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

	if repoGen, ok := m["repository"]; ok && repoGen != nil {
		rawRepo := repoGen.(json.RawMessage)
		if err := json.Unmarshal(rawRepo, &prj.Repo); err != nil {
			return nil, addDebugInfo(err)
		}
	}

	return &prj, nil
}

// mergeProjects merges projects p1 and p2 into a new Project.
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
