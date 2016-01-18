package calc

import (
	"reflect"
	"testing"
)

type testCase struct {
	input          string
	expectedOutput []token
}

func testLexerCases(t *testing.T, testCases []testCase) {
	for _, testCase := range testCases {
		output, err := lex([]byte(testCase.input))
		if err != nil {
			t.Error(err)
			continue
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

func TestParen(t *testing.T) {
	testLexerCases(t, []testCase{
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

func TestNumber(t *testing.T) {
	testLexerCases(t, []testCase{
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

func TestOperator(t *testing.T) {
	testLexerCases(t, []testCase{
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

func TestCombos(t *testing.T) {
	testLexerCases(t, []testCase{
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
