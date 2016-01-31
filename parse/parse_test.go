package calc

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/albrow/calc/ast"
	"github.com/albrow/calc/lex"
)

type parseTestCase struct {
	input          string
	expectedOutput ast.Node
	expectedError  error
}

func testParseCases(t *testing.T, parseTestCases []parseTestCase) {
	for _, testCase := range parseTestCases {
		tokens, err := lex.Lex([]byte(testCase.input))
		if err != nil {
			t.Error(err)
			continue
		}
		output, err := Parse(tokens)
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

// func justNumber(value string) node {
// 	base := &baseNode{}
// 	base.addChild(&numberNode{
// 		value: value,
// 	})
// 	return base
// }

// func TestParseNumber(t *testing.T) {
// 	testParseCases(t, []parseTestCase{
// 		{
// 			input:          "42",
// 			expectedOutput: justNumber("42"),
// 		},
// 	})
// }

// func TestParseParens(t *testing.T) {
// 	testParseCases(t, []parseTestCase{
// 		{
// 			input:          "(42)",
// 			expectedOutput: &baseNode{},
// 		},
// 	})
// }

// func TestParseOperators(t *testing.T) {
// 	testParseCases(t, []parseTestCase{
// 		{
// 			input:          "1 + 2",
// 			expectedOutput: &baseNode{},
// 		},
// 		{
// 			input:          "2 - 3",
// 			expectedOutput: &baseNode{},
// 		},
// 	})
// }

// func TestParseCombos(t *testing.T) {
// 	testParseCases(t, []parseTestCase{
// 		{
// 			input:          "(1 + 2) - 3",
// 			expectedOutput: &baseNode{},
// 		},
// 		{
// 			input:          "(1) - 2",
// 			expectedOutput: &baseNode{},
// 		},
// 		{
// 			input:          "(1 + (2) - 3) + (1 - (2 + 1))",
// 			expectedOutput: &baseNode{},
// 		},
// 	})
// }
