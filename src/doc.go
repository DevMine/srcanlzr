// Copyright 2014-2015 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	Package src provides a set of structures for representing a project with its
	related source code independently of the language. In other words, it
	provides a generic representation (abstraction) of a source code.


	Goal

	The goal of this package is to provide a generic representation of a project
	that can be analyzed by the anlzr package as well as an API for
	encoding/decoding it to/from JSON.


	Usage

	There are two kinds of program that interact with a src.Project: language
	parsers and VCS support tools. The former visits all source files inside
	the project folder and parse every source file in order to fill the
	src.Project.Packages field (and few others). The latter read the VCS folder
	that contains VCS data and fill the src.Project.Repo structure. The next
	two chapters treat about them more in details.


	Language parsers

	TODO


	VCS support tools

	TODO


	Example

	TODO


	Lines of Code counting

	The number of real lines of code must be precomputed by the language
	parsers. This is the only "feature" that must be precomputed because it may
	have multiple usages:

	1. Eliminate empty projects

	2. Evalutate project size

	3. Verify that the decoding is correct

	4. Normalize various counts

	5. ...

	Therefore, this count must be accurate and strictly follow the following
	rules:

	We only count statements and declarations as a line of code. Comments,
	package declaration, imports, expression, etc. must not be taken into
	account. Since an exemple is worth more than a thousand words, let's
	condider the following snippet:

	   // Package doc (does not count as a line of code)
	   package main // does not count as a line of code

	   import "fmt" // does not count as a line of code

	   func main() { // count as 1 line of code
	     fmt.Println(
	        "Hello, World!
	     ) // count as 1 line of code
	   }

	The expected number of lines of code is 2: The main function declaration
	and the call to fmt.Println function.


	Performance

	DevMine project is dealing with Terabytes of source code, therefore the
	JSON decoding must be efficient. That is why we implemented our own JSON
	decoder that focuses on performance. To do so, we had to make some
	choices and add some constraints for language parsers in order to make this
	process as fast as possible.

	JSON is usually unpredicatable which forces JSON parsers to be generic to
	deal with every possible kind of input. In DevMine, we have a well defined
	structure, thus instead of writting a generic JSON decoder we wrote one that
	decodes only src.Project objects. This really improves the performances
	since we don't need to use reflextion, generic types (interface{}) and type
	assertion. The drawback of this choice is that we have to update the decoder
	everytime we modify our structures.

	Most JSON parsers assume that the JSON input is potentially invalid
	(ie. malformed). We don't. Unlike json.Unmarshal, we don't Check for
	well-formedness.

	We also force the language parsers to put the "expression_name" and
	"statement_name" fields at the beginning of the JSON object. We use that
	convention to decode generic ast.Expr and ast.Stmt without reading the whole
	JSON object.


	Further improvements

	The code became quite repetitive. Since most of the logic has been
	encapsulated into helper methods, it would be really nice to generate the
	decoding methods using "go generate".
*/
package src
