package lex

import (
	"bytes"
	"fmt"
	"io"

	"github.com/albrow/calc/token"
)

func newNumberToken(value string) token.Token {
	return token.Token{
		Class: token.Number,
		Value: value,
	}
}

var (
	openParen = token.Token{
		Class: token.OpenParen,
		Value: "(",
	}
	closeParen = token.Token{
		Class: token.CloseParen,
		Value: ")",
	}
	opAdd = token.Token{
		Class: token.Add,
		Value: "+",
	}
	opSubtract = token.Token{
		Class: token.Subtract,
		Value: "-",
	}
)

func Lex(input []byte) ([]token.Token, error) {
	buf := bytes.NewBuffer(input)
	tokens := []token.Token{}
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

func readNumber(buf *bytes.Buffer) (token.Token, error) {
	value := []byte{}
	for {
		b, err := buf.ReadByte()
		if err != nil {
			if err == io.EOF {
				return newNumberToken(string(value)), nil
			}
			return token.Token{}, err
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
