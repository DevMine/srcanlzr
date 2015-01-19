// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src_test

import (
	"fmt"
	"log"

	"github.com/DevMine/srcanlzr/src"
)

func ExampleUnmarshal() {
	bs := []byte(`{"name": "foo", "languages": [], "packages": [], "loc": 0}`)
	p, _ := src.Unmarshal(bs)

	fmt.Printf("%#v\n", p)
	// Output:
	// src.Project{Name:"foo", Repo:(*model.Repository)(nil), RepoRaw:json.RawMessage(nil), Langs:[]*src.Language{}, Packages:[]*src.Package{}, LoC:0}
}

func ExampleMarshal() {
	p := &src.Project{
		Name:     "foo",
		Repo:     nil,
		RepoRaw:  nil,
		Langs:    []*src.Language{},
		Packages: []*src.Package{},
		LoC:      0,
	}

	bs, err := p.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))
	// Output:
	// {"name": "foo", "languages": [], "packages": [], "loc": 0}
}
