// Copyright 2014-2015 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	Package src provides a set of structures for representing a source code
	indepently of the language. In other words, it provides a generic
	representation (abstraction) of a source code.

	The number of real lines of code must be precomputed by the language
	parsers. This is the only "feature" that must be precomputed because it is
	used by srctool (https://github.com/DevMine/srctool) to eliminate empty
	packages. Since this information is already calculated, the source analyzers
	won't re-count. They will just use the total as it is. Therefore, it must be
	accurate.

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
	and the call to the fmt.Println function.
*/
package src
