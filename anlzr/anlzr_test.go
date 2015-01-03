// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package anlzr_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/DevMine/srcanlzr/anlzr"
	"github.com/DevMine/srcanlzr/src"
)

//var testdata = os.Getenv("GOPATH") + "/src/github.com/DevMine/srcanlzr/testdata/go.json"
var testdata = "../testdata/go.json"

func TestRunAnalyzers(t *testing.T) {
	bs, err := ioutil.ReadFile(testdata)
	if err != nil {
		t.Fatal(err)
	}

	p := new(src.Project)
	if err := json.Unmarshal(bs, p); err != nil {
		t.Fatal(err)
	}

	res, err := anlzr.RunAnalyzers(p, anlzr.LoC{})
	if err != nil {
		t.Fatal(err)
	}

	if fmt.Sprintf("%.3f", res.AverageFuncLen) != "7.256" {
		t.Errorf("average_function_length: expected 7.256, found %.3f",
			res.AverageFuncLen)
	}

	if res.MinFuncLen != 1 {
		t.Errorf("min_function_length: expected 1, found %d",
			res.MinFuncLen)
	}

	if res.MaxFuncLen != 413 {
		t.Errorf("max_function_length: expected 413, found %d",
			res.MaxFuncLen)
	}

	if res.MedianFuncLen != 4 {
		t.Errorf("max_function_length: expected 4, found %d",
			res.MedianFuncLen)
	}

	if res.TotalLoC != 128042 {
		t.Errorf("total_loc: expected 128042, found %d",
			res.TotalLoC)
	}
}
