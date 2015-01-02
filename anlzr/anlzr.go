// Copyright 2014 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	Package anlz provides the necessary API for making source code analyzis.
*/
package anlzr

import (
	"encoding/xml"

	"github.com/DevMine/srcanlzr/src"
)

type Analyzer interface {
	Analyze(p *src.Project, r *Result) error
}

type Result struct {
	XMLName        xml.Name          `json:"-" xml:"result"`
	ProgLangs      []src.Language    `json:"programming_languages" xml:"programming-languages"`
	AverageFuncLen float32           `json:"average_function_length" xml:"average-function-length"`
	MaxFuncLen     int64             `json:"max_function_length" xml:"max-function-length"`
	MinFuncLen     int64             `json:"min_function_length" xml:"min-function-length"`
	MedianFuncLen  int64             `json:"median_function_length" xml:"median-function-length"`
	TotalLoC       int64             `json:"total_loc" xml:"total-loc"`
	Complexity     ComplexityMetrics `json:"complexity" xml:"complexity"`
}

// Cyclomatic complexity metrics, also known as McCabe metric.
type ComplexityMetrics struct {
	AveragePerFunc float32 `json:"average_per_func"` // Average complexity per function.
	AveragePerFile float32 `json:"average_per_file"` // Average complexity per file.
}

func RunAnalyzers(p *src.Project, a ...Analyzer) (*Result, error) {
	r := &Result{
		ProgLangs:      p.ProgLangs,
		AverageFuncLen: -1.0,
		MaxFuncLen:     -1,
		MinFuncLen:     -1,
		MedianFuncLen:  -1,
		TotalLoC:       -1,
		Complexity:     ComplexityMetrics{AveragePerFunc: -1, AveragePerFile: -1},
	}

	for _, anlzr := range a {
		err := anlzr.Analyze(p, r)
		if err != nil {
			return nil, err
		}
	}

	return r, nil
}
