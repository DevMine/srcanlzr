// Copyright 2014-2015 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package token

const (
	IntLit    = "INT"
	FloatLit  = "FLOAT"
	ImagLit   = "IMAG"
	CharLit   = "CHAR"
	StringLit = "STRING"
	BoolLit   = "BOOl"
	NilLit    = "NIL"
)

// Binary operators
const (
	// numerical operators
	ADD = "ADD"
	SUB = "SUB"
	MUL = "MUL"
	QUO = "QUO"
	MOD = "MOD"

	// logical operators
	AND     = "AND"         // binary and (&)
	OR      = "OR"          // binary or (|)
	XOR     = "XOR"         // binary xor (^)
	SHL     = "SHIFT_LEFT"  // binary left shift <<
	SHR     = "SHIFT_RIGHT" // binary right shift >>
	AND_NOT = "AND_NOT"     // binary and not (&^)

	// comparison operators
	NEQ  = "NEQ"  // not equal
	LEQ  = "LEQ"  // less or equal
	GEQ  = "GEQ"  // greater or equal
	EQ   = "EQ"   // equal
	LSS  = "LSS"  // less
	GTR  = "GTR"  // greater
	LAND = "LAND" // and (&&)
	LOR  = "LOR"  // or (||)
)

// Kind of declarations
const (
	ConstDecl = "CONSTANT" // constant
	VarDecl   = "VAR"      // variable
)

const (
	UnaryExprName           = "UNARY"
	BinaryExprName          = "BINARY"
	TernaryExprName         = "TERNARY"
	IncDecExprName          = "INC_DEC"
	CallExprName            = "CALL"
	ConstructorCallExprName = "CONSTRUCTOR_CALL"
	ArrayExprName           = "ARRAY"
	IndexExprName           = "INDEX"

	BasicLitName = "BASIC_LIT"
	FuncLitName  = "FUNC_LIT"
	ClassLitName = "CLASS_LIT"
	ArrayLitName = "ARRAY_LIT"

	StructTypeName = "STRUCT_TYPE"

	AttrRefName = "ATTR_REF"

	ValueSpecName = "VALUE_SPEC"

	IdentName = "IDENT"
)

// Increment/Decrement operators
const (
	INC = "INC"
	DEC = "DEC"
)

// Unary operators
const (
	NOT  = "NOT"
	ADDR = "ADDR" // memory address (&foo)
	STAR = "STAR" // dereference operator (*foo)

	NEG = "NEG"
	POS = "POS"
)

const (
	IfStmtName        = "IF"
	SwitchStmtName    = "SWITCH"
	LoopStmtName      = "LOOP"
	RangeLoopStmtName = "RANGE_LOOP"
	AssignStmtName    = "ASSIGN"
	DeclStmtName      = "DECL"
	ReturnStmtName    = "RETURN"
	ExprStmtName      = "EXPR"
	TryStmtName       = "TRY"
	ThrowStmtName     = "THROW"
	OtherStmtName     = "OTHER"
)
