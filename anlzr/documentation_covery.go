// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package anlzr

import (
	"strings"

	"github.com/DevMine/srcanlzr/src"
)

type CommentRatios struct {
	TypeComRatio   float32 `json:"type_comment_ratio"`
	StructComRatio float32 `json:"structure_comment_ratio"`
	ConstComRatio  float32 `json:"constant_comment_ratio"`
	VarsComRatio   float32 `json:"variable_comment_ratio"`
	FuncComRatio   float32 `json:"function_comment_ratio"`
	InterComRatio  float32 `json:"interface_comment_ratio"`
	ClassComRatio  float32 `json:"class_comment_ratio"`
	MethComRatio   float32 `json:"method_comment"`
	AttrComRatio   float32 `json:"attribute_comment_ratio"`
	EnumComRatio   float32 `json:"enumeration_comment_ratio"`
}

type counters struct {
	nbComType   int
	nbType      int
	nbComStruct int
	nbStruct    int
	nbComConst  int
	nbConst     int
	nbComVars   int
	nbVars      int
	nbComFunc   int
	nbFunc      int
	nbComInter  int
	nbInter     int
	nbComClas   int
	nbClas      int
	nbComAttr   int
	nbAttr      int
	nbComEnum   int
	nbEnum      int
	nbComFcts   int
	nbFcts      int
}

func (dc CommentRatios) Analyze(p *src.Project, r *Result) error {

	cnt := counters{}

	for _, pack := range p.Packages {
		for _, srcFile := range pack.SrcFiles {
			cnt.nbType += len(srcFile.TypeSpecs)
			for _, typeSpec := range srcFile.TypeSpecs {
				if hasComment(typeSpec.Doc) {
					cnt.nbComType++
				}
			}

			cnt.nbStruct += len(srcFile.Structs)
			for _, structDecl := range srcFile.Structs {
				if hasComment(structDecl.Doc) {
					cnt.nbComStruct++
				}
			}

			for _, constDecl := range srcFile.Constants {
				if isVisible(constDecl.Visibility) {
					cnt.nbConst++
					if hasComment(constDecl.Doc) {
						cnt.nbComConst++
					}
				}
			}

			for _, varDecl := range srcFile.Vars {
				if isVisible(varDecl.Visibility) {
					cnt.nbVars++
					if hasComment(varDecl.Doc) {
						cnt.nbComVars++
					}
				}
			}

			for _, funcDecl := range srcFile.Funcs {
				if isVisible(funcDecl.Visibility) {
					cnt.nbFunc++
					if hasComment(funcDecl.Doc) {
						cnt.nbComFunc++
					}
				}
			}

			for _, interfaceDecl := range srcFile.Interfaces {
				cnt.interfaceCommentCoverage(interfaceDecl)
			}

			for _, classDecl := range srcFile.Classes {
				cnt.classesCommentCoverage(classDecl)
			}

			for _, enumDecl := range srcFile.Enums {
				cnt.enumCommentCoverage(enumDecl)
			}
		}
	}

	if cnt.nbType != 0 {
		dc.TypeComRatio = float32(cnt.nbComType) / float32(cnt.nbType)
	}
	if cnt.nbStruct != 0 {
		dc.StructComRatio = float32(cnt.nbComStruct) / float32(cnt.nbStruct)
	}
	if cnt.nbConst != 0 {
		dc.ConstComRatio = float32(cnt.nbComConst) / float32(cnt.nbConst)
	}
	if cnt.nbVars != 0 {
		dc.VarsComRatio = float32(cnt.nbComVars) / float32(cnt.nbVars)
	}
	if cnt.nbFunc != 0 {
		dc.FuncComRatio = float32(cnt.nbComFunc) / float32(cnt.nbFunc)
	}
	if cnt.nbInter != 0 {
		dc.InterComRatio = float32(cnt.nbComInter) / float32(cnt.nbInter)
	}
	if cnt.nbClas != 0 {
		dc.ClassComRatio = float32(cnt.nbComClas) / float32(cnt.nbClas)
	}
	if cnt.nbFcts != 0 {
		dc.MethComRatio = float32(cnt.nbComFcts) / float32(cnt.nbFcts)
	}
	if cnt.nbAttr != 0 {
		dc.AttrComRatio = float32(cnt.nbComAttr) / float32(cnt.nbAttr)
	}
	if cnt.nbEnum != 0 {
		dc.EnumComRatio = float32(cnt.nbComEnum) / float32(cnt.nbEnum)
	}

	r.DocCoverage = dc

	return nil
}

func (cnt *counters) interfaceCommentCoverage(interfaceDecl *src.Interface) {

	if isVisible(interfaceDecl.Visibility) {
		cnt.nbInter++
		if hasComment(interfaceDecl.Doc) {
			cnt.nbComInter++
		}
	}

	for _, proto := range interfaceDecl.Protos {
		if isVisible(proto.Visibility) {
			cnt.nbFcts++
			if hasComment(proto.Doc) {
				cnt.nbComFcts++
			}
		}
	}

}

func (cnt *counters) classesCommentCoverage(classDecl *src.ClassDecl) {

	if isVisible(classDecl.Visibility) {
		cnt.nbClas++
		if hasComment(classDecl.Doc) {
			cnt.nbComClas++
		}
	}

	for _, attr := range classDecl.Attrs {
		if isVisible(attr.Visibility) {
			cnt.nbAttr++
			if hasComment(attr.Doc) {
				cnt.nbComAttr++
			}
		}
	}

	cnt.functionsCommentCoverage(classDecl.Constructors, classDecl.Destructors, classDecl.Methods)

	for _, class := range classDecl.NestedClasses {
		cnt.classesCommentCoverage(class)
	}

}

func (cnt *counters) enumCommentCoverage(enumDecl *src.EnumDecl) {

	if isVisible(enumDecl.Visibility) {
		cnt.nbEnum++
		if hasComment(enumDecl.Doc) {
			cnt.nbComEnum++
		}
	}

	for _, attr := range enumDecl.Attrs {
		if isVisible(attr.Visibility) {
			cnt.nbAttr++
			if hasComment(attr.Doc) {
				cnt.nbComAttr++
			}
		}
	}

	cnt.functionsCommentCoverage(enumDecl.Constructors, enumDecl.Destructors, enumDecl.Methods)
}

func (cnt *counters) functionsCommentCoverage(cstrs []*src.ConstructorDecl, dstrs []*src.DestructorDecl, mthds []*src.MethodDecl) {

	for _, fct := range cstrs {
		if isVisible(fct.Visibility) {
			cnt.nbFcts++
			if hasComment(fct.Doc) {
				cnt.nbComFcts++
			}
		}
	}
	for _, fct := range dstrs {
		if isVisible(fct.Visibility) {
			cnt.nbFcts++
			if hasComment(fct.Doc) {
				cnt.nbComFcts++
			}
		}
	}
	for _, fct := range mthds {
		if isVisible(fct.Visibility) {
			cnt.nbFcts++
			if hasComment(fct.Doc) {
				cnt.nbComFcts++
			}
		}
	}

}

func isVisible(v string) bool {
	return v == src.PublicVisibility ||
		v == src.ProtectedVisibility ||
		v == src.PackageVisibility
}

func hasComment(doc []string) bool {
	return doc != nil || len(strings.Trim(strings.Join(doc, ""), " ")) != 0
}
