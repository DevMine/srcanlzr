// Copyright 2014-2015 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"regexp"
	"strings"
	"text/template"
)

// kind of AST node
const (
	Expression = iota
	Statement
)

// go source file header
const header = `// Copyright 2014-2015 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// DO NOT EDIT: This source file has been generated by gen/gen_ast_decoder.go

package src

import (
	"github.com/DevMine/srcanlzr/src/ast"
	"github.com/DevMine/srcanlzr/src/token"
)

`

// template for expression
const tmplExpr = `
func (dec *decoder) decode{{ .Name }}() *ast.{{ .Name }} {
	if !dec.assertNewObject() {
		return nil
	}
	if dec.isEmptyObject() {
		dec.err = errors.New("{{ .Name }} object cannot be empty")
		return nil
	}
	if dec.err != nil {
		return nil
	}
	return dec.decode{{ .Name }}Attrs()
}

func (dec *decoder) decode{{ .Name }}Attrs() *ast.{{ .Name }} {
	expr := ast.{{ .Name }}{ExprName: token.{{ .Name }}Name}
	for {
		key, err := dec.scan.nextKey()
		if err != nil {
			if err == io.EOF {
				break
			}
			dec.err = err
			return nil
		}
		if key == "" {
			dec.err = errors.New("empty key")
			return nil
		}

		{{ if .HasBasicType }}
		val, tok, err := dec.scan.nextValue()
		{{ else }}
		_, _, err = dec.scan.nextValue()
		{{ end }}
		if err != nil {
			dec.err = err
			return nil
		}

		switch key {
		{{ range $index, $field := .Fields }}
		case "{{ $field.JSONName }}":
			{{ if $field.BasicType }}
				if tok != scan{{ $field.Type }}Lit {
					dec.err = fmt.Errorf("expected '{{ $field.Type }} literal', found '%v'", tok)
					return nil
				}
				expr.{{ $field.Name }}, dec.err = dec.unmarshal{{ $field.Type }}(val)
			{{ else }}
				{{ if $field.Array }}
					dec.scan.back()
					expr.{{ $field.Name }} = dec.decode{{ $field.Type }}s()
				{{ else }}
					dec.scan.back()
					expr.{{ $field.Name }} = dec.decode{{ $field.Type }}s()
				{{ end }}
			{{ end }}
		{{ end }}
		default:
			dec.err = fmt.Errorf("unexpected value for the key '%s' of a {{.Name}} object", key)
		}

		if dec.err != nil {
			return nil
		}

		if dec.isEndObject() {
			break
		}
		if err != nil {
			return nil
		}
	}
	return &expr
}

`

// template for statements
const tmplStmt = ``

// template for "normal" types
const tmplOther = ``

const outputPath = "decode_gen.go"

type DecoderTmpl struct {
	Name   string
	Fields []Field
}

func (dt DecoderTmpl) HasBasicType() bool {
	for _, field := range dt.Fields {
		if field.BasicType {
			return true
		}
	}
	return false
}

type Field struct {
	Name      string
	JSONName  string
	Type      string
	BasicType bool
	Array     bool
}

func genExprs(w io.Writer, exprs []DecoderTmpl) error {
	if len(exprs) == 0 {
		return nil
	}

	// TODO generate decodeExpr

	t := template.Must(template.New("expressions").Parse(tmplExpr))

	for _, expr := range exprs {
		if err := t.Execute(w, expr); err != nil {
			return err
		}
	}
	return nil
}

func genStmts(w io.Writer, stmts []DecoderTmpl) error {
	if len(stmts) == 0 {
		return nil
	}

	return nil
}

func genOthers(w io.Writer, others []DecoderTmpl) error {
	if len(others) == 0 {
		return nil
	}

	return nil
}

func extractType(field *ast.Field) (typ string, array bool, basicType bool) {
	switch s := field.Type.(type) {
	case *ast.StarExpr:
		if ident, ok := s.X.(*ast.Ident); ok {
			typ = ident.String()
		}
	case *ast.ArrayType:
		switch elt := s.Elt.(type) {
		case *ast.StarExpr:
			if ident, ok := elt.X.(*ast.Ident); ok {
				typ = ident.String()
				array = true
			}
		case *ast.Ident:
			typ = elt.String()
			array = true
		}
	case *ast.Ident:
		typ = s.String()
		typ = strings.ToUpper(string(typ[0])) + typ[1:]

		switch typ {
		case "String", "Integer", "Float":
			basicType = true
		}
	}
	return
}

func extractTag(tag *ast.BasicLit) string {
	// extract JSON tag
	re := regexp.MustCompile("`json:\"([a-zA-Z0-9_]+)(,omitempty)?\"`")
	m := re.FindStringSubmatch(tag.Value)
	if len(m) < 2 {
		return ""
	}
	return m[1]
}

func warn(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func warnf(format string, a ...interface{}) {
	warn(fmt.Sprintf(format, a...) + "\n")
}

func fatal(a ...interface{}) {
	warn(a...)
	os.Exit(1)
}

func fatalf(format string, a ...interface{}) {
	warnf(format, a...)
	os.Exit(1)
}

func main() {
	flag.Parse()

	out, err := os.Create(outputPath)
	if err != nil {
		fatal(err)
	}
	defer out.Close()

	// write file header
	out.WriteString(header)

	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "./ast/ast.go", nil, 0)
	if err != nil {
		fatal(err)
	}

	exprs := []DecoderTmpl{}
	stmts := []DecoderTmpl{}
	others := []DecoderTmpl{}

	for _, decl := range f.Decls {
		var genDecl *ast.GenDecl
		var ok bool
		if genDecl, ok = decl.(*ast.GenDecl); !ok {
			continue
		}

		for _, spec := range genDecl.Specs {
			var typeSpec *ast.TypeSpec
			var ok bool
			if typeSpec, ok = spec.(*ast.TypeSpec); !ok {
				continue
			}

			dec := DecoderTmpl{Name: typeSpec.Name.String(), Fields: []Field{}}

			var structType *ast.StructType
			if structType, ok = typeSpec.Type.(*ast.StructType); !ok {
				continue
			}
			var kind int
			for _, field := range structType.Fields.List {
				fieldTmpl := Field{}
				// XXX: handle compositions
				if len(field.Names) == 0 {
					continue
				}
				fieldTmpl.Name = field.Names[0].String()
				fieldTmpl.JSONName = extractTag(field.Tag)
				switch fieldTmpl.JSONName {
				case "expression_name":
					kind = Expression
				case "statement_name":
					kind = Statement
				}
				if fieldTmpl.Type, fieldTmpl.Array, fieldTmpl.BasicType = extractType(field); fieldTmpl.Type == "" {
					warnf("invalid type for %s.%s", dec.Name, fieldTmpl.Name)
					continue
				}
				dec.Fields = append(dec.Fields, fieldTmpl)
			}

			switch kind {
			case Expression:
				exprs = append(exprs, dec)
			case Statement:
				stmts = append(stmts, dec)
			default:
				others = append(others, dec)
			}
		}
	}

	// generate code
	if err := genExprs(out, exprs); err != nil {
		fatal(err)
	}
	if err := genStmts(out, stmts); err != nil {
		fatal(err)
	}
	if err := genStmts(out, stmts); err != nil {
		fatal(err)
	}
}
