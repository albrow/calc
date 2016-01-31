package calc

import (
	"errors"
	"fmt"
	"io"

	"github.com/albrow/calc/ast"
	"github.com/albrow/calc/token"
)

// For our parser we consider the following grammar:
//
// E -> E + E | E - E | (E) | Number

// Rewritten to avoid left recursion:
//
// E -> E' Op E | E'
// E' -> Number | "(" E ")"
// Op -> "+" | "-"

func Parse(tokens []token.Token) (ast.Node, error) {
	buf := token.NewBuffer(tokens)
	root := ast.New()
	tree, err := e(buf, root)
	if err != nil {
		if err == io.EOF {
			return nil, errors.New("Unexpected end of input")
		}
		return nil, err
	}
	return tree, nil
}

func newUnexpectedTokenError(t token.Token) error {
	return fmt.Errorf("Unexpected token: %s", t.Value)
}

// nullTerm returns a function which will check if the next token is the given
// token class. If it is, that function will return the tree without making any
// changes. If it is not, that function will return an error. nullTerm is used
// to check for terminal tokens which should not be added to the AST, such as
// "(" and ")".
func nullTerm(class token.Class) func(*token.Buffer, ast.Node) (ast.Node, error) {
	return func(buf *token.Buffer, tree ast.Node) (ast.Node, error) {
		if t, err := buf.Read(); err != nil {
			return nil, err
		} else if t.Class == class {
			return tree, nil
		} else {
			return nil, newUnexpectedTokenError(t)
		}
	}
}

func termAdd(buf *token.Buffer, tree ast.Node) (ast.Node, error) {
	if t, err := buf.Read(); err != nil {
		return nil, err
	} else if t.Class == token.Add {
		// Add an operator node to the AST. The newly added node will become
		// the new root of the tree.
		newTree := tree.Copy()
		opNode := &ast.Operator{
			Class: ast.OpAdd,
		}
		newTree.AddChild(opNode)
		return opNode, nil
	} else {
		return nil, newUnexpectedTokenError(t)
	}
}

func termSubtract(buf *token.Buffer, tree ast.Node) (ast.Node, error) {
	if t, err := buf.Read(); err != nil {
		return nil, err
	} else if t.Class == token.Subtract {
		// Add an operator node to the AST. The newly added node will become
		// the new root of the tree.
		newTree := tree.Copy()
		opNode := &ast.Operator{
			Class: ast.OpSubtract,
		}
		newTree.AddChild(opNode)
		return opNode, nil
	} else {
		return nil, newUnexpectedTokenError(t)
	}
}

func termNumber(buf *token.Buffer, tree ast.Node) (ast.Node, error) {
	if t, err := buf.Read(); err != nil {
		return nil, err
	} else if t.Class == token.Number {
		// Create a copy of the current root and add a number node to it.
		newTree := tree.Copy()
		numNode := &ast.Number{
			Value: t.Value,
		}
		newTree.AddChild(numNode)
		return newTree, nil
	} else {
		return nil, newUnexpectedTokenError(t)
	}
}

// E -> E' Op E | E'
func e(buf *token.Buffer, tree ast.Node) (ast.Node, error) {
	return nil, nil
}

// E1 -> E' Op E
func e1(buf *token.Buffer, tree ast.Node) (ast.Node, error) {
	return nil, nil
}

// E2 -> E'
func e2(buf *token.Buffer, tree ast.Node) (ast.Node, error) {
	return nil, nil
}

// E' -> Number | "(" E ")"
func ep(buf *token.Buffer, tree ast.Node) (ast.Node, error) {
	return nil, nil
}

// E'1 -> Number
func ep1(buf *token.Buffer, tree ast.Node) (ast.Node, error) {
	return nil, nil
}

// E'2 -> "(" E ")"
func ep2(buf *token.Buffer, tree ast.Node) (ast.Node, error) {
	return nil, nil
}

// Op -> "+" | "-"
func op(buf *token.Buffer, tree ast.Node) (ast.Node, error) {
	return nil, nil
}

// Op1 -> "+"
func op1(buf *token.Buffer, tree ast.Node) (ast.Node, error) {
	return nil, nil
}

// Op2 -> "-"
func op2(buf *token.Buffer, tree ast.Node) (ast.Node, error) {
	return nil, nil
}
