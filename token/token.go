package token

import "fmt"

type Class uint

const (
	Number Class = iota
	OpenParen
	CloseParen
	Add
	Subtract
)

func (c Class) String() string {
	switch c {
	case OpenParen:
		return "token.OpenParen"
	case CloseParen:
		return "token.CloseParen"
	case Number:
		return "token.Number"
	case Add:
		return "token.Add"
	case Subtract:
		return "token.Subtract"
	default:
		panic(fmt.Sprintf("Unknown token class: %v", uint(c)))
	}
}

type Token struct {
	Class Class
	Value string
}
