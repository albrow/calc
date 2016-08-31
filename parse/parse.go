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

func newUnexpectedTokenErrorNext(buf *token.Buffer) error {
	t, err := buf.Read()
	if err != nil {
		return err
	}
	return newUnexpectedTokenError(t)
}

// nullTerm returns a function which will check if the next token is the given
// token class. If it is, that function will return nil, nil. If it is not,
// that function will return an error. nullTerm is used to check for terminal
// tokens which should not be added to the AST, such as "(" and ")".
func nullTerm(class token.Class) func(*token.Buffer) error {
	return func(buf *token.Buffer) (err error) {
		origPos := buf.Pos()
		defer func() {
			if err != nil {
				buf.MustSeek(origPos)
			}
		}()
		if t, err := buf.Read(); err != nil {
			return err
		} else if t.Class == class {
			return nil
		} else {
			return newUnexpectedTokenError(t)
		}
	}
}

var termOpenParen = nullTerm(token.OpenParen)
var termCloseParen = nullTerm(token.CloseParen)

func termOp(buf *token.Buffer) (node ast.Node, err error) {
	origPos := buf.Pos()
	defer func() {
		if err != nil {
			buf.MustSeek(origPos)
		}
	}()
	t, err := buf.Read()
	if err != nil {
		return nil, err
	}
	switch t.Class {
	case token.Add:
		return &ast.Operator{
			Class: ast.OpAdd,
		}, nil
	case token.Subtract:
		return &ast.Operator{
			Class: ast.OpSubtract,
		}, nil
	}
	return nil, newUnexpectedTokenError(t)
}

func termNumber(buf *token.Buffer) (node ast.Node, err error) {
	origPos := buf.Pos()
	defer func() {
		if err != nil {
			buf.MustSeek(origPos)
		}
	}()
	if t, err := buf.Read(); err != nil {
		return nil, err
	} else if t.Class == token.Number {
		return &ast.Number{
			Value: t.Value,
		}, nil
	} else {
		return nil, newUnexpectedTokenError(t)
	}
}

// E -> E' Op E | E'
func e(buf *token.Buffer, tree ast.Node) (newTree ast.Node, err error) {
	origPos := buf.Pos()
	defer func() {
		if err != nil {
			buf.MustSeek(origPos)
		}
	}()
	if newTree, err := e1(buf, tree); err == nil {
		return newTree, nil
	} else if newTree, err := e2(buf, tree); err == nil {
		return newTree, nil
	}
	buf.MustSeek(origPos)
	return nil, newUnexpectedTokenErrorNext(buf)
}

// E1 -> E' Op E
func e1(buf *token.Buffer, tree ast.Node) (newTree ast.Node, err error) {
	origPos := buf.Pos()
	defer func() {
		if err != nil {
			buf.MustSeek(origPos)
		}
	}()
	newTree = tree.Copy()
	epTree, err := ep(buf, tree.Copy())
	if err != nil {
		return nil, err
	}
	newTree.AddChildren(epTree.Children())
	if buf.Pos() >= buf.Len() {
		return nil, io.EOF
	}
	opNode, err := termOp(buf)
	if err != nil {
		return nil, err
	}
	newTree.AddChild(opNode)
	if buf.Pos() >= buf.Len() {
		return nil, io.EOF
	}
	eTree, err := e(buf, tree.Copy())
	if err != nil {
		return nil, err
	}
	newTree.AddChildren(eTree.Children())
	return newTree, nil
}

// E2 -> E'
func e2(buf *token.Buffer, tree ast.Node) (ast.Node, error) {
	return ep(buf, tree)
}

// E' -> Number | "(" E ")"
func ep(buf *token.Buffer, tree ast.Node) (newTree ast.Node, err error) {
	origPos := buf.Pos()
	defer func() {
		if err != nil {
			buf.MustSeek(origPos)
		}
	}()
	if newTree, err := ep1(buf, tree); err == nil {
		return newTree, nil
	} else if newTree, err := ep2(buf, tree); err == nil {
		return newTree, nil
	}
	buf.MustSeek(origPos)
	return nil, newUnexpectedTokenErrorNext(buf)
}

// E'1 -> Number
func ep1(buf *token.Buffer, tree ast.Node) (newTree ast.Node, err error) {
	origPos := buf.Pos()
	defer func() {
		if err != nil {
			buf.MustSeek(origPos)
		}
	}()
	node, err := termNumber(buf)
	if err != nil {
		return nil, err
	}
	newTree = tree.Copy()
	newTree.AddChild(node)
	return newTree, nil
}

// E'2 -> "(" E ")"
func ep2(buf *token.Buffer, tree ast.Node) (newTree ast.Node, err error) {
	origPos := buf.Pos()
	defer func() {
		if err != nil {
			buf.MustSeek(origPos)
		}
	}()
	newTree = tree.Copy()
	expr := ast.New()
	newTree.AddChild(expr)
	if err := termOpenParen(buf); err != nil {
		return nil, err
	}
	if buf.Pos() >= buf.Len() {
		return nil, io.EOF
	}
	eTree, err := e(buf, tree.Copy())
	if err != nil {
		return nil, err
	}
	expr.AddChildren(eTree.Children())
	if buf.Pos() >= buf.Len() {
		return nil, io.EOF
	}
	if err := termCloseParen(buf); err != nil {
		return nil, err
	}
	return newTree, nil
}
