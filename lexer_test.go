package calc

import (
	"errors"
	"reflect"
	"testing"
)

type lexTestCase struct {
	input          string
	expectedOutput []token
	expectedError  error
}

func testLexerCases(t *testing.T, lexTestCases []lexTestCase) {
	for _, testCase := range lexTestCases {
		output, err := lex([]byte(testCase.input))
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
				"For input: %s\nExpected: %v\n  but got: %v",
				testCase.input,
				testCase.expectedOutput,
				output,
			)
		}
	}
}

func TestLexParen(t *testing.T) {
	testLexerCases(t, []lexTestCase{
		{
			input: "(",
			expectedOutput: []token{
				openParen,
			},
		},
		{
			input: ")",
			expectedOutput: []token{
				closeParen,
			},
		},
		{
			input: "()(()())",
			expectedOutput: []token{
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
			expectedOutput: []token{
				openParen,
				closeParen,
			},
		},
	})
}

func TestLexNumber(t *testing.T) {
	testLexerCases(t, []lexTestCase{
		{
			input: "1",
			expectedOutput: []token{
				newNumberToken("1"),
			},
		},
		{
			input: "123456",
			expectedOutput: []token{
				newNumberToken("123456"),
			},
		},
		{
			input: " \n\t123456\t\n",
			expectedOutput: []token{
				newNumberToken("123456"),
			},
		},
	})
}

func TestLexOperator(t *testing.T) {
	testLexerCases(t, []lexTestCase{
		{
			input: "+",
			expectedOutput: []token{
				opAdd,
			},
		},
		{
			input: "-",
			expectedOutput: []token{
				opSubtract,
			},
		},
		{
			input: "++--+-",
			expectedOutput: []token{
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
			expectedOutput: []token{
				opAdd,
				opSubtract,
			},
		},
	})
}

func TestLexCombos(t *testing.T) {
	testLexerCases(t, []lexTestCase{
		{
			input: "2 + 2",
			expectedOutput: []token{
				newNumberToken("2"),
				opAdd,
				newNumberToken("2"),
			},
		},
		{
			input: "\t(2 + 2) -\n (31 + 42)\n\n",
			expectedOutput: []token{
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
	testLexerCases(t, []lexTestCase{
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
