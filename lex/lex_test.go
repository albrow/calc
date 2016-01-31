package lex

import (
	"errors"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/albrow/calc/token"
)

type testCase struct {
	input          string
	expectedOutput []token.Token
	expectedError  error
}

func testLexerCases(t *testing.T, testCases []testCase) {
	for _, testCase := range testCases {
		output, err := Lex([]byte(testCase.input))
		if err != nil {
			if testCase.expectedError == nil {
				t.Error(err)
				continue
			} else if !reflect.DeepEqual(testCase.expectedError, err) {
				t.Errorf(
					"Expected error: %v\nBut got: %v",
					testCase.expectedError,
					err,
				)
			}
		}
		if !reflect.DeepEqual(output, testCase.expectedOutput) {
			t.Errorf(
				"For input: %s\nExpected: %s\n  but got: %s",
				testCase.input,
				spew.Sdump(testCase.expectedOutput),
				spew.Sdump(output),
			)
		}
	}
}

func TestLexParen(t *testing.T) {
	testLexerCases(t, []testCase{
		{
			input: "(",
			expectedOutput: []token.Token{
				openParen,
			},
		},
		{
			input: ")",
			expectedOutput: []token.Token{
				closeParen,
			},
		},
		{
			input: "()(()())",
			expectedOutput: []token.Token{
				openParen,
				closeParen,
				openParen,
				openParen,
				closeParen,
				openParen,
				closeParen,
				closeParen,
			},
		},
		{
			input: " \n(\t)\t ",
			expectedOutput: []token.Token{
				openParen,
				closeParen,
			},
		},
	})
}

func TestLexNumber(t *testing.T) {
	testLexerCases(t, []testCase{
		{
			input: "1",
			expectedOutput: []token.Token{
				newNumberToken("1"),
			},
		},
		{
			input: "123456",
			expectedOutput: []token.Token{
				newNumberToken("123456"),
			},
		},
		{
			input: " \n\t123456\t\n",
			expectedOutput: []token.Token{
				newNumberToken("123456"),
			},
		},
	})
}

func TestLexOperator(t *testing.T) {
	testLexerCases(t, []testCase{
		{
			input: "+",
			expectedOutput: []token.Token{
				opAdd,
			},
		},
		{
			input: "-",
			expectedOutput: []token.Token{
				opSubtract,
			},
		},
		{
			input: "++--+-",
			expectedOutput: []token.Token{
				opAdd,
				opAdd,
				opSubtract,
				opSubtract,
				opAdd,
				opSubtract,
			},
		},
		{
			input: "\t+ -\t \n",
			expectedOutput: []token.Token{
				opAdd,
				opSubtract,
			},
		},
	})
}

func TestLexCombos(t *testing.T) {
	testLexerCases(t, []testCase{
		{
			input: "2 + 2",
			expectedOutput: []token.Token{
				newNumberToken("2"),
				opAdd,
				newNumberToken("2"),
			},
		},
		{
			input: "\t(2 + 2) -\n (31 + 42)\n\n",
			expectedOutput: []token.Token{
				openParen,
				newNumberToken("2"),
				opAdd,
				newNumberToken("2"),
				closeParen,
				opSubtract,
				openParen,
				newNumberToken("31"),
				opAdd,
				newNumberToken("42"),
				closeParen,
			},
		},
	})
}

func TestLexUnexpectedChar(t *testing.T) {
	testLexerCases(t, []testCase{
		{
			input:         "f2.0 + 2",
			expectedError: errors.New("Unexpected character at 0: 'f'"),
		},
		{
			input:         "2.0 + 2",
			expectedError: errors.New("Unexpected character at 1: '.'"),
		},
		{
			input:         "(2 + 2) - foo",
			expectedError: errors.New("Unexpected character at 10: 'f'"),
		},
	})
}
