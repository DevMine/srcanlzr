//usr/bin/env go run $0 $@; exit
// Copyright 2014 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"path/filepath"

	"github.com/DevMine/srcanlzr/anlzr"
	"github.com/DevMine/srcanlzr/src"
)

const (
	version = "0.0.0"
)

var format *string

func formatOutput(r *anlzr.Result) ([]byte, error) {
	var bs []byte
	var err error

	switch *format {
	case "JSON":
		bs, err = json.Marshal(r)
	case "XML":
		bs, err = xml.Marshal(r)
	case "protobuf":
		err = errors.New("protobuf is not yet implemented")
	default:
		err = errors.New("unsupported output format")
	}

	return bs, err
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func main() {
	flag.Usage = func() {
		fmt.Printf("usage: %s [PROJECT PATH]\n\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}

	format = flag.String("f", "JSON", "Output format. Possible values are: JSON, XML, protobuf")
	outputFileName := flag.String("o", "", "Output file name. By default, the output is set to stdout")
	vflag := flag.Bool("v", false, "Print version.")
	flag.Parse()

	if *vflag {
		fmt.Printf("%s - v%s\n", filepath.Base(os.Args[0]), version)
		return
	}

	if len(flag.Args()) != 1 {
		fmt.Fprint(os.Stderr, "invalid # of arguments\n\n")
		flag.Usage()
	}

	out := os.Stdout
	if len(*outputFileName) > 0 {
		var err error
		out, err = os.Open(*outputFileName)
		if err != nil {
			fatal(err)
		}
	}

	bs, _ := ioutil.ReadFile(flag.Arg(0))

	p := new(src.Project)
	if err := json.Unmarshal(bs, p); err != nil {
		fatal(err)
	}

	res, err := anlzr.RunAnalyzers(p, anlzr.LoC{})
	if err != nil {
		fatal(err)
	}

	bs, err = formatOutput(res)
	if err != nil {
		fatal(err)
	}

	fmt.Fprintln(out, string(bs))
}
