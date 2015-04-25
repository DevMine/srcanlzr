// Copyright 2014-2015 The DevMine Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package anlzr

import (
	"strings"

	"github.com/DevMine/srcanlzr/src"
)

type CommentRatios struct {
	TypeComRatio   float32
	StructComRatio float32
	ConstComRatio  float32
	VarsComRatio   float32
	FuncComRatio   float32
	InterComRatio  float32
	ClassComRatio  float32
	MethComRatio   float32
	AttrComRatio   float32
	EnumComRatio   float32
}

func (dc CommentRatios) Analyze(p *src.Project, r *Result) error {

	var nbComType int
	var nbType int

	var nbComStruct int
	var nbStruct int

	var nbComConst int
	var nbConst int

	var nbComVars int
	var nbVars int

	var nbComFunc int
	var nbFunc int

	var nbComInter int
	var nbInter int

	var nbComProto int
	var nbProto int

	var nbComClas int
	var nbClas int

	var nbComAttr int
	var nbAttr int

	var nbComEnum int
	var nbEnum int

	var nbComFcts int
	var nbFcts int

	for _, pack := range p.Packages {
		for _, srcFile := range pack.SrcFiles {
			nbType += len(srcFile.TypeSpecs)
			for _, typeSpec := range srcFile.TypeSpecs {
				if hasComment(typeSpec.Doc) {
					nbComType++
				}
			}

			nbStruct += len(srcFile.Structs)
			for _, structDecl := range srcFile.Structs {
				if hasComment(structDecl.Doc) {
					nbComStruct++
				}
			}

			for _, constDecl := range srcFile.Constants {
				if isVisible(constDecl.Visibility) {
					nbConst++
					if hasComment(constDecl.Doc) {
						nbComConst++
					}
				}
			}

			for _, varDecl := range srcFile.Vars {
				if isVisible(varDecl.Visibility) {
					nbVars++
					if hasComment(varDecl.Doc) {
						nbComVars++
					}
				}
			}

			for _, funcDecl := range srcFile.Funcs {
				if isVisible(funcDecl.Visibility) {
					nbFunc++
					if hasComment(funcDecl.Doc) {
						nbComFunc++
					}
				}
			}

			for _, interfaceDecl := range srcFile.Interfaces {
				dc.interfaceCommentCovery(interfaceDecl, &nbComInter, &nbInter, &nbComProto, &nbProto)
			}

			for _, classDecl := range srcFile.Classes {
				dc.classesCommentCovery(classDecl, &nbComClas, &nbClas, &nbComAttr, &nbAttr, &nbComFcts, &nbFcts)
			}

			for _, enumDecl := range srcFile.Enums {
				dc.enumCommentCovery(enumDecl, &nbComEnum, &nbEnum, &nbComAttr, &nbAttr, &nbComFcts, &nbFcts)
			}
		}
	}

	if nbType != 0 {
		dc.TypeComRatio = float32(nbComType) / float32(nbType)
	}
	if nbStruct != 0 {
		dc.StructComRatio = float32(nbComStruct) / float32(nbStruct)
	}
	if nbConst != 0 {
		dc.ConstComRatio = float32(nbComConst) / float32(nbConst)
	}
	if nbVars != 0 {
		dc.VarsComRatio = float32(nbComVars) / float32(nbVars)
	}
	if nbFunc != 0 {
		dc.FuncComRatio = float32(nbComFunc) / float32(nbFunc)
	}
	if nbInter != 0 {
		dc.InterComRatio = float32(nbComInter) / float32(nbInter)
	}
	if nbProto != 0 {
		dc.MethComRatio = float32(nbComProto) / float32(nbProto)
	}
	if nbClas != 0 {
		dc.ClassComRatio = float32(nbComClas) / float32(nbClas)
	}
	if nbAttr != 0 {
		dc.AttrComRatio = float32(nbComAttr) / float32(nbAttr)
	}
	if nbEnum != 0 {
		dc.EnumComRatio = float32(nbComEnum) / float32(nbEnum)
	}
	if nbAttr != 0 {
		dc.AttrComRatio = float32(nbComAttr) / float32(nbAttr)
	}
	if nbFcts != 0 {
		dc.MethComRatio = float32(nbComFcts) / float32(nbFcts)
	}

	r.DocCovery = dc

	return nil
}

func (dc *CommentRatios) interfaceCommentCovery(interfaceDecl *src.Interface, nbComInter, nbInter, nbComProto, nbProto *int) {

	if isVisible(interfaceDecl.Visibility) {
		*nbInter++
		if hasComment(interfaceDecl.Doc) {
			*nbComInter++
		}
	}

	for _, proto := range interfaceDecl.Protos {
		if isVisible(proto.Visibility) {
			*nbProto++
			if hasComment(proto.Doc) {
				*nbComProto++
			}
		}
	}

}

func (dc *CommentRatios) classesCommentCovery(classDecl *src.ClassDecl, nbComClas, nbClas, nbComAttr, nbAttr, nbComFcts, nbFcts *int) {

	if isVisible(classDecl.Visibility) {
		*nbClas++
		if hasComment(classDecl.Doc) {
			*nbComClas++
		}
	}

	for _, attr := range classDecl.Attrs {
		if isVisible(attr.Visibility) {
			*nbAttr++
			if hasComment(attr.Doc) {
				*nbComAttr++
			}
		}
	}

	dc.functionsCommentCovery(classDecl.Constructors, classDecl.Destructors, classDecl.Methods, nbComFcts, nbFcts)

	for _, class := range classDecl.NestedClasses {
		dc.classesCommentCovery(class, nbComClas, nbClas, nbComAttr, nbAttr, nbComFcts, nbFcts)
	}

}

func (dc *CommentRatios) enumCommentCovery(enumDecl *src.EnumDecl, nbComEnum, nbEnum, nbComAttr, nbAttr, nbComFcts, nbFcts *int) {

	if isVisible(enumDecl.Visibility) {
		*nbEnum++
		if hasComment(enumDecl.Doc) {
			*nbComEnum++
		}
	}

	for _, attr := range enumDecl.Attrs {
		if isVisible(attr.Visibility) {
			*nbAttr++
			if hasComment(attr.Doc) {
				*nbComAttr++
			}
		}
	}

	dc.functionsCommentCovery(enumDecl.Constructors, enumDecl.Destructors, enumDecl.Methods, nbComFcts, nbFcts)
}

func (dc *CommentRatios) functionsCommentCovery(cstrs []*src.ConstructorDecl, dstrs []*src.DestructorDecl, mthds []*src.MethodDecl, nbComFcts, nbFcts *int) {

	for _, fct := range cstrs {
		if isVisible(fct.Visibility) {
			*nbFcts++
			if hasComment(fct.Doc) {
				*nbComFcts++
			}
		}
	}
	for _, fct := range dstrs {
		if isVisible(fct.Visibility) {
			*nbFcts++
			if hasComment(fct.Doc) {
				*nbComFcts++
			}
		}
	}
	for _, fct := range mthds {
		if isVisible(fct.Visibility) {
			*nbFcts++
			if hasComment(fct.Doc) {
				*nbComFcts++
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
