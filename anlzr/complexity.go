// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package anlzr

import "github.com/DevMine/srcanlzr/src"

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

				/*for _, stmt := range f.Body {
					switch stmt.(type) {
					case src.IfStmt, src.LoopStmt:
						fileComplexity++
					}
				}

				fileComplexity += int64(len(f.Type.Results))*/
			}

			for _, cls := range sf.Classes {
				for _, m := range cls.Methods {
					numFuncs++
					fileComplexity += methodCyclomaticComplexity(m)

					/*for _, stmt := range m.Body {
						switch stmt.(type) {
						case src.IfStmt, src.LoopStmt:
							fileComplexity++
						}
					}

					fileComplexity += int64(len(m.Type.Results))*/
				}
			}

			for _, mod := range sf.Traits {
				for _, m := range mod.Methods {
					numFuncs++
					fileComplexity += methodCyclomaticComplexity(m)

					/*for _, stmt := range m.Body {
						switch stmt.(type) {
						case src.IfStmt, src.LoopStmt:
							fileComplexity++
						}
					}

					fileComplexity += int64(len(m.Type.Results))*/
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
		cc += statementComplexity(s)
	}

	return cc
}

func methodCyclomaticComplexity(m *src.MethodDecl) int64 {
	cc := int64(1) // cyclomatic complexity

	for _, s := range m.Body {
		cc += statementComplexity(s)
	}

	return cc
}

func exprComplexity(e src.Expr) int64 {
	var c int64

	switch e.(type) {
	case *src.BinaryExpr:
		be := e.(*src.BinaryExpr)
		// TODO check the operator
		c += exprComplexity(be.LeftExpr)
		c += exprComplexity(be.RightExpr)
	}

	return c
}

func statementComplexity(s src.Stmt) int64 {
	var c int64

	switch s.(type) {
	case *src.IfStmt:
		c++

		is := s.(*src.IfStmt)

		if is.Cond != nil {
			c += exprComplexity(is.Cond)
		}

		for _, s := range is.Body {
			c += statementComplexity(s)
		}

		if is.Else != nil && len(is.Else) > 0 {
			for _, s := range is.Else {
				c += statementComplexity(s)
			}
		}
	case *src.LoopStmt:
		c++

		ls := s.(*src.LoopStmt)

		if ls.Cond != nil {
			c += exprComplexity(ls.Cond)
		}

		for _, s := range ls.Body {
			c += statementComplexity(s)
		}

		if ls.Else != nil && len(ls.Else) > 0 {
			for _, s := range ls.Else {
				c += statementComplexity(s)
			}
		}
	case *src.RangeLoopStmt:
		c++

		rls := s.(*src.RangeLoopStmt)

		for _, s := range rls.Body {
			c += statementComplexity(s)
		}
	case *src.SwitchStmt:
		c++

		ss := s.(*src.SwitchStmt)

		if ss.Cond != nil {
			c += exprComplexity(ss.Cond)
		}

		if ss.CaseClauses != nil {
			for _, cc := range ss.CaseClauses {
				c++

				for _, cond := range cc.Conds {
					if cond != nil {
						c += exprComplexity(cond)
					}
				}

				for _, s := range cc.Body {
					c += statementComplexity(s)
				}
			}
		}

		if ss.Default != nil && len(ss.Default) > 0 {
			for _, s := range ss.Default {
				c += statementComplexity(s)
			}
		}

	}

	return c
}
