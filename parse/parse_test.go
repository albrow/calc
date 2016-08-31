package calc

import (
	"fmt"
	"testing"

	"github.com/albrow/calc/ast"
	"github.com/albrow/calc/lex"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type parseTestCase struct {
	input          string
	expectedOutput ast.Node
	expectedError  error
}

type parseTestCaseWithFormat struct {
	input          string
	expectedOutput string
	expectedError  error
}

func testParseCases(t *testing.T, parseTestCases []parseTestCase) {
	for i, tc := range parseTestCases {
		tcInfo := fmt.Sprintf("test case: %d\ninput: %s", i, tc.input)
		tokens, err := lex.Lex([]byte(tc.input))
		require.NoError(t, err, tcInfo)
		output, err := Parse(tokens)
		if tc.expectedError != nil {
			require.Error(t, err, tcInfo)
			assert.Exactly(t, tc.expectedError, err, tcInfo)
		} else {
			require.NoError(t, err, tcInfo)
			assert.Exactly(t, tc.expectedOutput, output, "%s\n\nExpected:\n%s\n\nGot:\n%s\n\n", tcInfo, tc.expectedOutput.Format(0), output.Format(0))
		}
	}
}

func testParseCasesWithFormat(t *testing.T, parseTestCases []parseTestCaseWithFormat) {
	for i, tc := range parseTestCases {
		tcInfo := fmt.Sprintf("test case: %d\ninput: %s", i, tc.input)
		tokens, err := lex.Lex([]byte(tc.input))
		require.NoError(t, err, tcInfo)
		output, err := Parse(tokens)
		if tc.expectedError != nil {
			require.Error(t, err, tcInfo)
			assert.Exactly(t, tc.expectedError, err, tcInfo)
		} else {
			require.NoError(t, err, tcInfo)
			assert.Exactly(t, tc.expectedOutput, output.Format(0), "%s\n\nExpected:\n%s\n\nGot:\n%s\n\n", tcInfo, tc.expectedOutput, output.Format(0))
		}
	}
}

func justNumber(value string) ast.Node {
	base := ast.New()
	base.AddChild(&ast.Number{
		Value: value,
	})
	return base
}

func TestParse_Number(t *testing.T) {
	testParseCases(t, []parseTestCase{
		{
			input:          "42",
			expectedOutput: justNumber("42"),
		},
	})
}

func operation(a string, opClass ast.OpClass, b string) ast.Node {
	base := ast.New()
	base.AddChildren([]ast.Node{
		&ast.Number{
			Value: a,
		},
		&ast.Operator{
			Class: opClass,
		},
		&ast.Number{
			Value: b,
		},
	})
	return base
}

func TestParse_Operation(t *testing.T) {
	testParseCases(t, []parseTestCase{
		{
			input:          "2 + 3",
			expectedOutput: operation("2", ast.OpAdd, "3"),
		},
		{
			input:          "5 - 4",
			expectedOutput: operation("5", ast.OpSubtract, "4"),
		},
	})
}

var parensOutput0 = `|- base
  |- base
    |- 42
`

var parensOutput1 = `|- base
  |- base
    |- 1
    |- +
    |- 4
`

var parensOutput2 = `|- base
  |- base
    |- 1
    |- +
    |- 2
    |- -
    |- 3
    |- +
    |- 4
`

func TestParse_Parens(t *testing.T) {
	testParseCasesWithFormat(t, []parseTestCaseWithFormat{
		{
			input:          "(42)",
			expectedOutput: parensOutput0,
		},
		{
			input:          "(1 + 4)",
			expectedOutput: parensOutput1,
		},
		{
			input:          "(1 + 2 - 3 + 4)",
			expectedOutput: parensOutput2,
		},
	})
}

var combosOutput0 = `|- base
  |- base
    |- 1
    |- +
    |- 2
  |- -
  |- 3
`

var combosOutput1 = `|- base
  |- base
    |- 1
  |- -
  |- 2
`

var combosOutput2 = `|- base
  |- base
    |- 1
    |- +
    |- base
      |- 2
    |- -
    |- 3
  |- +
  |- base
    |- 4
    |- -
    |- base
      |- 5
      |- +
      |- 6
`

func TestParseCombos(t *testing.T) {
	testParseCasesWithFormat(t, []parseTestCaseWithFormat{
		{
			input:          "(1 + 2) - 3",
			expectedOutput: combosOutput0,
		},
		{
			input:          "(1) - 2",
			expectedOutput: combosOutput1,
		},
		{
			input:          "(1 + (2) - 3) + (4 - (5 + 6))",
			expectedOutput: combosOutput2,
		},
	})
}
