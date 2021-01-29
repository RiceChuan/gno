package gno

import (
	"fmt"
	"strings"
	"unicode"
	"unsafe"
)

//----------------------------------------
// Misc.

// Returns the associated machine operation for binary AST operations.  TODO:
// to make this faster and inlineable, remove the switch statement and create a
// mathematical mapping between them.
func word2BinaryOp(w Word) Op {
	switch w {
	case ADD:
		return OpAdd
	case SUB:
		return OpSub
	case MUL:
		return OpMul
	case QUO:
		return OpQuo
	case REM:
		return OpRem
	case BAND:
		return OpBand
	case BOR:
		return OpBor
	case XOR:
		return OpXor
	case SHL:
		return OpShl
	case SHR:
		return OpShr
	case BAND_NOT:
		return OpBandn
	case LAND:
		return OpLand
	case LOR:
		return OpLor
	case EQL:
		return OpEql
	case LSS:
		return OpLss
	case GTR:
		return OpGtr
	case NEQ:
		return OpNeq
	case LEQ:
		return OpLeq
	case GEQ:
		return OpGeq
	default:
		panic(fmt.Sprintf("unexpected binary operation word %v", w.String()))
	}
}

func word2UnaryOp(w Word) Op {
	switch w {
	case ADD:
		return OpUpos
	case SUB:
		return OpUneg
	case NOT:
		return OpUnot
	case XOR:
		return OpUxor
	case MUL:
		panic("unexpected unary operation * - use StarExpr instead")
	case BAND:
		panic("unexpected unary operation & - use RefExpr instead")
	case ARROW:
		return OpUrecv
	default:
		panic("unexpected unary operation")
	}
}

func toString(n Node) string {
	if n == nil {
		return "<nil>"
	}
	return n.String()
}

// true if the first rune is uppercase.
func isUpper(s string) bool {
	var first rune
	for _, c := range s {
		first = c
		break
	}
	return unicode.IsUpper(first)
}

//----------------------------------------
// converting uintptr to bytes.

const sizeOfUintPtr = unsafe.Sizeof(uintptr(0))

func uintptrToBytes(u *uintptr) []byte {
	return (*[sizeOfUintPtr]byte)(unsafe.Pointer(u))[:]
}

func defaultPkgName(gopkgPath string) Name {
	parts := strings.Split(gopkgPath, "/")
	last := parts[len(parts)-1]
	parts = strings.Split(last, "-")
	name := parts[len(parts)-1]
	name = strings.ToLower(name)
	return Name(name)
}
