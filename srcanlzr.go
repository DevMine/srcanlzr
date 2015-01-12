// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime/pprof"

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
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to file")
	memprofile := flag.String("memprofile", "", "write memory profile to this file")
	vflag := flag.Bool("v", false, "Print version.")
	flag.Parse()

	if *vflag {
		fmt.Printf("%s - v%s\n", filepath.Base(os.Args[0]), version)
		return
	}

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}

	var bs []byte
	var err error

	if l := len(flag.Args()); l > 1 {
		// more that one file are passed as arguments
		fmt.Fprint(os.Stderr, "too many arguments\n\n")
		flag.Usage()
	} else if l == 0 {
		// no argument, we read from stdin
		buf := new(bytes.Buffer)
		r := bufio.NewReader(os.Stdin)
		if _, err = io.Copy(buf, r); err != nil {
			fatal(err)
		}
		bs = buf.Bytes()
	} else {
		var f *os.File
		if f, err = os.Open(flag.Arg(0)); err != nil {
			fatal(err)
		}

		buf := new(bytes.Buffer)
		r := bufio.NewReader(f)

		// there is only one file passed as argument, so we read it
		if _, err = io.Copy(buf, r); err != nil {
			fatal(err)
		}

		bs = buf.Bytes()
	}

	out := os.Stdout
	if len(*outputFileName) > 0 {
		var err error
		out, err = os.Open(*outputFileName)
		if err != nil {
			fatal(err)
		}
	}

	p, err := src.Unmarshal(bs)
	if err != nil {
		fatal(err)
	}

	res, err := anlzr.RunAnalyzers(p, anlzr.LoC{}, anlzr.Complexity{})
	if err != nil {
		fatal(err)
	}

	bs, err = formatOutput(res)
	if err != nil {
		fatal(err)
	}

	fmt.Fprintln(out, string(bs))
}
