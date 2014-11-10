// Copyright 2014 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	Package repo provides function for accessing VCS information.

	The currently supported VCS are: Git, Mercurial.
*/
package repo

import (
	"errors"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"regexp"
)

// Repo types
const (
	None = iota
	Git
	Mercurial
	Subversion
	Bazaar
	CVS // unfortunately, it is still in use on some projects :(
)

type Repo struct {
	Type int
	URL  url.URL
}

func New(path string) (*Repo, error) {
	// check for git repo
	if _, err := os.Stat(path + "/.git"); err == nil {
		return &Repo{
			Type: Git,
			URL:  *extractGitURL(path),
		}, nil
	}

	// check for hg repo
	if _, err := os.Stat(path + "/.hg"); err == nil {
		return &Repo{
			Type: Mercurial,
			URL:  *extractMercurialURL(path),
		}, nil
	}

	return nil, errors.New("no supported repo found")
}

func extractGitURL(path string) *url.URL {
	f, err := os.Open(path + "/.git/config")
	if err != nil {
		log.Fatal(err)
	}

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile("url ?= ?(.+)")
	m := re.FindStringSubmatch(string(bs))
	if len(m) != 2 {
		log.Fatal("invalid git config file")
	}

	repoURL, err := url.Parse(m[1])
	if err != nil {
		log.Fatal(err)
	}

	return repoURL
}

func extractMercurialURL(path string) *url.URL {
	f, err := os.Open(path + "/.hg/hgrc")
	if err != nil {
		log.Fatal(err)
	}

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	// XXX naive ? check if always true
	re := regexp.MustCompile("default ?= ?(.+)")
	m := re.FindStringSubmatch(string(bs))
	if len(m) != 2 {
		log.Fatal("invalid hg config file")
	}

	repoURL, err := url.Parse(m[1])
	if err != nil {
		log.Fatal(err)
	}

	return repoURL
}
