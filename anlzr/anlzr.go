// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package anlzr provides the necessary API for making source code analyzis.
package anlzr

import "github.com/DevMine/srcanlzr/src"

type Analyzer interface {
	Analyze(p *src.Project, r *Result) error
}

type Result struct {
	ProgLangs      []Language        `json:"programming_languages" xml:"programming-languages"`
	AverageFuncLen float32           `json:"average_function_length" xml:"average-function-length"`
	MaxFuncLen     int64             `json:"max_function_length" xml:"max-function-length"`
	MinFuncLen     int64             `json:"min_function_length" xml:"min-function-length"`
	MedianFuncLen  int64             `json:"median_function_length" xml:"median-function-length"`
	TotalLoC       int64             `json:"total_loc" xml:"total-loc"`
	Complexity     ComplexityMetrics `json:"complexity" xml:"complexity"`
	DocCovery      CommentRatios     `json:"documentation_coverage" xml:"documentation-coverage"`
}

// CyclomaticComplexity metrics, also known as McCabe metric.
type ComplexityMetrics struct {
	AveragePerFunc float32 `json:"average_per_func" xml:"average-per-func"` // Average complexity per function.
	AveragePerFile float32 `json:"average_per_file" xml:"average-per-file"` // Average complexity per file.
}

type Language struct {
	src.Language
	Lines int64 `json:"lines" xml:"lines"`
}

func RunAnalyzers(p *src.Project, a ...Analyzer) (*Result, error) {
	r := &Result{
		ProgLangs:      []Language{},
		AverageFuncLen: -1.0,
		MaxFuncLen:     -1,
		MinFuncLen:     -1,
		MedianFuncLen:  -1,
		TotalLoC:       -1,
		Complexity:     ComplexityMetrics{AveragePerFunc: -1, AveragePerFile: -1},
		DocCovery:      CommentRatios{},
	}

	for _, anlzr := range a {
		err := anlzr.Analyze(p, r)
		if err != nil {
			return nil, err
		}
	}

	return r, nil
}
