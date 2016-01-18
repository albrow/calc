package calc

import (
	"bytes"
	"fmt"
	"io"
)

type tokenClass uint

const (
	tokenParen tokenClass = iota
	tokenNumber
	tokenOperator
)

func (tc tokenClass) String() string {
	switch tc {
	case tokenParen:
		return "tokenParen"
	case tokenNumber:
		return "tokenNumber"
	case tokenOperator:
		return "tokenOperator"
	default:
		panic(fmt.Sprintf("Unknown token class: %v", uint(tc)))
	}
}

type token struct {
	class tokenClass
	value string
}

func newNumberToken(value string) token {
	return token{
		class: tokenNumber,
		value: value,
	}
}

var (
	openParen = token{
		class: tokenParen,
		value: "(",
	}
	closeParen = token{
		class: tokenParen,
		value: ")",
	}
	opAdd = token{
		class: tokenOperator,
		value: "+",
	}
	opSubtract = token{
		class: tokenOperator,
		value: "-",
	}
)

func lex(input []byte) ([]token, error) {
	buf := bytes.NewBuffer(input)
	tokens := []token{}
	for {
		b, err := buf.ReadByte()
		if err != nil {
			if err == io.EOF {
				return tokens, nil
			}
			return nil, err
		}
		switch b {
		case '(':
			tokens = append(tokens, openParen)
		case ')':
			tokens = append(tokens, closeParen)
		case '+':
			tokens = append(tokens, opAdd)
		case '-':
			tokens = append(tokens, opSubtract)
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			buf.UnreadByte()
			token, err := readNumber(buf)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
		case ' ', '\t', '\n':
			continue
		default:
			pos := len(input) - buf.Len() - 1
			return nil, fmt.Errorf(
				"Unexpected character at %d: '%s'", pos, []byte{b},
			)
		}
	}
}

func readNumber(buf *bytes.Buffer) (token, error) {
	value := []byte{}
	for {
		b, err := buf.ReadByte()
		if err != nil {
			if err == io.EOF {
				return newNumberToken(string(value)), nil
			}
			return token{}, err
		}
		switch b {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			value = append(value, b)
		default:
			buf.UnreadByte()
			return newNumberToken(string(value)), nil
		}
	}
}
