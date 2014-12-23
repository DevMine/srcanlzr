// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

// Repo contains information relative to a repository.
type Repo struct {
	// The Version Control System used.
	VCS string `json:"vcs"`

	// The clone URL
	cloneURL string `json:"url"`
}
