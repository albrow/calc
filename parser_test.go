package calc

import (
	"reflect"
	"testing"
)

type parseTestCase struct {
	input          string
	expectedOutput *node
	expectedError  error
}

func testParseCases(t *testing.T, parseTestCases []parseTestCase) {
	for _, testCase := range parseTestCases {
		tokens, err := lex([]byte(testCase.input))
		if err != nil {
			t.Error(err)
			continue
		}
		output, err := parse(tokens)
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

func TestParseNumber(t *testing.T) {
	testParseCases(t, []parseTestCase{
		{
			input:          "42",
			expectedOutput: &node{},
		},
	})
}

func TestParseParens(t *testing.T) {
	testParseCases(t, []parseTestCase{
		{
			input:          "(42)",
			expectedOutput: &node{},
		},
	})
}

func TestParseOperators(t *testing.T) {
	testParseCases(t, []parseTestCase{
		{
			input:          "1 + 2",
			expectedOutput: &node{},
		},
		{
			input:          "2 - 3",
			expectedOutput: &node{},
		},
	})
}

func TestParseCombos(t *testing.T) {
	testParseCases(t, []parseTestCase{
		{
			input:          "(1 + 2) - 3",
			expectedOutput: &node{},
		},
		{
			input:          "(1) - 2",
			expectedOutput: &node{},
		},
		{
			input:          "(1 + (2) - 3) + (1 - (2 + 1))",
			expectedOutput: &node{},
		},
	})
}
