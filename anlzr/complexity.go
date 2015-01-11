// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package anlzr

import (
	"fmt"

	"github.com/DevMine/srcanlzr/src"
)

type Complexity struct{}

func (c Complexity) Analyze(p *src.Project, r *Result) error {
	cm := ComplexityMetrics{}

	var totalFuncs int64
	var totalFiles int64

	var totalComplexityPerFunc int64
	var totalComplexityPerFile float32

	for _, pkg := range p.Packages {
		for _, sf := range pkg.SrcFiles {
			var fileComplexity int64
			var numFuncs int64

			for _, f := range sf.Funcs {
				numFuncs++
				fileComplexity += functionCyclomaticComplexity(f)

				for _, stmt := range f.Body {
					switch stmt.(type) {
					case src.IfStmt, src.LoopStmt:
						fileComplexity++
					}
				}

				fileComplexity += int64(len(f.Type.Results))
			}

			for _, cls := range sf.Classes {
				for _, m := range cls.Methods {
					numFuncs++
					fileComplexity += methodCyclomaticComplexity(m)

					for _, stmt := range m.Body {
						switch stmt.(type) {
						case src.IfStmt, src.LoopStmt:
							fileComplexity++
						}
					}

					fileComplexity += int64(len(m.Type.Results))

				}
			}

			for _, mod := range sf.Traits {
				for _, m := range mod.Methods {
					numFuncs++
					fileComplexity += methodCyclomaticComplexity(m)

					for _, stmt := range m.Body {
						switch stmt.(type) {
						case src.IfStmt, src.LoopStmt:
							fileComplexity++
						}
					}

					fileComplexity += int64(len(m.Type.Results))

				}
			}

			if numFuncs > 0 {
				totalFiles++
				totalFuncs += numFuncs
				totalComplexityPerFunc += fileComplexity
				totalComplexityPerFile += float32(fileComplexity) / float32(numFuncs)
			}
		}
	}

	cm.AveragePerFunc = float32(totalComplexityPerFunc) / float32(totalFuncs)
	cm.AveragePerFile = totalComplexityPerFile / float32(totalFiles)

	r.Complexity = cm

	return nil
}

func functionCyclomaticComplexity(f *src.FuncDecl) int64 {
	cc := int64(1) // cyclomatic complexity

	for _, s := range f.Body {
		cc += statementComplexity(&s)
	}

	return cc
}

func methodCyclomaticComplexity(m *src.Method) int64 {
	cc := int64(1) // cyclomatic complexity

	for _, s := range m.Body {
		cc += statementComplexity(&s)
	}

	return cc
}

func statementComplexity(s src.Stmt) int64 {
	var c int64

	switch s.(type) {
	case src.IfStmt:
		fmt.Println("foo")
		c++

		stmt := s.(src.IfStmt)

		for _, s := range stmt.StmtsList {
			c += statementComplexity(s)
		}
	case src.LoopStmt:
		fmt.Println("bar")
		c++

		stmt := s.(src.LoopStmt)

		for _, s := range stmt.StmtsList {
			c += statementComplexity(s)
		}
	}

	return c
}
