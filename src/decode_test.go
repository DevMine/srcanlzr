// Copyright 2014-2015 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package src

import (
	"bytes"
	"testing"
)

func TestDecodeStringsList(t *testing.T) {
	expected := []string{"foo", "bar"}
	buf := bytes.NewBufferString(`["foo","bar"]`)
	dec := newDecoder(buf)
	sl := dec.decodeStringsList()
	if dec.err != nil {
		t.Fatal(dec.err)
	}
	if !stringsSlicesEquals(sl, expected) {
		t.Errorf("decodeStringsList: found %v, expected %v", sl, expected)
	}
}

func stringsSlicesEquals(sl1, sl2 []string) bool {
	if len(sl1) != len(sl2) {
		return false
	}
	for i := 0; i < len(sl1); i++ {
		if sl1[i] != sl2[i] {
			return false
		}
	}
	return true
}
