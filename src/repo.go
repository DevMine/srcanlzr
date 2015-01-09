// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

// Repo contains information relative to a repository.
type Repo struct {
	// The Version Control System used.
	VCS string `json:"vcs"`

	// The clone URL
	CloneURL string `json:"url"`
}

// newRepo creates a new Repo from a generic map.
func newRepo(m map[string]interface{}) (*Repo, error) {
	var err error
	errPrefix := "src/repo"
	repo := Repo{}

	repo.VCS, err = extractStringValue("vcs", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	repo.CloneURL, err = extractStringValue("url", errPrefix, m)
	if err != nil {
		return nil, addDebugInfo(err)
	}

	return &repo, nil
}
