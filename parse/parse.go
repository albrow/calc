package calc

import (
	"github.com/albrow/calc/ast"
	"github.com/albrow/calc/token"
)

// For our parser we consider the following grammar:
//
// E -> E + E | E - E | (E) | Number

// Rewritten to avoid left recursion:
//
// E -> E' Op E | E'
// E' -> Number | (E)
// Op -> + | -

func Parse(tokens []token.Token) (ast.Node, error) {
	return nil, nil
}
