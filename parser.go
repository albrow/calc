package calc

import (
	"fmt"
)

// For our parser we consider the following grammar:
//
// E -> E + E | E - E | (E) | Number

// E -> E' Op E | E'
// E' -> Number | (E)
// Op -> + | -

type node struct {
	children []node
	parent   *node
	token    token
}

func unexpectedTokenError(tokens []token, pos *int) error {
	return fmt.Errorf("Unexpected token: %s", tokens[*pos].value)
}

func parse(tokens []token) (*node, error) {
	pos := 0
	parent := &node{}
	if e(tokens, &pos) && pos == len(tokens) {
		return parent, nil
	}
	if pos > len(tokens)-1 {
		return nil, fmt.Errorf("Unexpected end of input")
	}
	return nil, unexpectedTokenError(tokens, &pos)
}

func term(class tokenClass) func(tokens []token, pos *int) bool {
	return func(tokens []token, pos *int) bool {
		if *pos > len(tokens)-1 {
			return false
		}
		if tokens[*pos].class == class {
			// fmt.Printf("At %d, term %s returns true\n", *pos, class)
			*pos = *pos + 1
			return true
		}
		return false
	}
}

var (
	termOpenParen  = term(tokenOpenParen)
	termCloseParen = term(tokenCloseParen)
	termNumber     = term(tokenNumber)
	termAdd        = term(tokenAdd)
	termSubtract   = term(tokenSubtract)
)

var op1 = termAdd
var op2 = termSubtract

// Op = + | -
func op(tokens []token, pos *int) bool {
	// fmt.Printf("op, pos: %d\n", *pos)
	if *pos > len(tokens)-1 {
		return false
	}
	start := *pos
	if *pos = start; op1(tokens, pos) {
		// fmt.Printf("op1 is true at %d\n", *pos-1)
		return true
	} else if *pos = start; op2(tokens, pos) {
		// fmt.Printf("op2 is true at %d\n", *pos-1)
		return true
	}
	return false
}

// E1 -> E' Op E
func e1(tokens []token, pos *int) bool {
	return ep(tokens, pos) &&
		op(tokens, pos) &&
		e(tokens, pos)
}

// E2 -> E'
func e2(tokens []token, pos *int) bool {
	return ep(tokens, pos)
}

// E -> E' Op E | E'
func e(tokens []token, pos *int) bool {
	// fmt.Printf("e, pos: %d\n", *pos)
	if *pos > len(tokens)-1 {
		return false
	}
	start := *pos
	if *pos = start; e1(tokens, pos) {
		// fmt.Printf("e1 is true at %d\n", *pos-1)
		return true
	} else if *pos = start; e2(tokens, pos) {
		// fmt.Printf("e2 is true at %d\n", *pos-1)
		return true
	}
	return false
}

// E'1 -> Number
var ep1 = termNumber

// E'2 -> (E)
func ep2(tokens []token, pos *int) bool {
	return termOpenParen(tokens, pos) &&
		e(tokens, pos) &&
		termCloseParen(tokens, pos)
}

// E' -> Number | (E)
func ep(tokens []token, pos *int) bool {
	// fmt.Printf("ep, pos: %d\n", *pos)
	if *pos > len(tokens)-1 {
		return false
	}
	start := *pos
	if *pos = start; ep1(tokens, pos) {
		// fmt.Printf("ep1 is true at %d\n", *pos-1)
		return true
	} else if *pos = start; ep2(tokens, pos) {
		// fmt.Printf("ep2 is true at %d\n", *pos-1)
		return true
	}
	return false
}
