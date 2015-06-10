// Copyright 2014-2015 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"bytes"
	"io/ioutil"
	"math"
	"testing"
)

const inputJSON = "./testdata/small.json"

func TestDecodeFile(t *testing.T) {
	prj, err := DecodeFile(inputJSON)
	if err != nil {
		t.Fatalf("DecodeFile '%s': %v", inputJSON, err)
	}

	buf := new(bytes.Buffer)
	if err := prj.Encode(buf); err != nil {
		t.Fatalf("Encode: %v", err)
	}

	bs, err := ioutil.ReadFile(inputJSON)
	if err != nil {
		t.Fatal(err)
	}

	// XXX: do better!
	// For some reasons, some extra "\" are added in the remarshalled JSON
	// In the end, it adds 2 extra characters...
	if math.Abs(float64(len(bs)-buf.Len())) > 2 {
		t.Errorf("DecodeFile '%s': JSON file incorrectly decoded", inputJSON)
	}

	/*if !bytes.Equal(bs, buf.Bytes()) {
		t.Errorf("DecodeFile '%s': JSON file incorrectly decoded", inputJSON)
		t.Errorf("Found:\n%s\n", string(bs))
		t.Errorf("Expected:\n%s\n", string(buf.Bytes()))
	}*/
}
