package eval

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/albrow/calc/ast"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEval(t *testing.T) {
	addTree := ast.New()
	addTree.AddChildren([]ast.Node{
		&ast.Number{
			Value: "2",
		},
		&ast.Operator{
			Class: ast.OpAdd,
		},
		&ast.Number{
			Value: "3",
		},
	})
	subTree := ast.New()
	subTree.AddChildren([]ast.Node{
		&ast.Number{
			Value: "5",
		},
		&ast.Operator{
			Class: ast.OpSubtract,
		},
		&ast.Number{
			Value: "3",
		},
	})
	testCases := []struct {
		tree     ast.Node
		expected *big.Rat
	}{
		{
			tree: &ast.Number{
				Value: "42",
			},
			expected: big.NewRat(42, 1),
		},
		{
			tree:     addTree,
			expected: big.NewRat(5, 1),
		},
		{
			tree:     subTree,
			expected: big.NewRat(2, 1),
		},
	}
	for i, tc := range testCases {
		tcInfo := fmt.Sprintf("test case: %d\ninput:\n%s\n", i, tc.tree.Format(0))
		actual, err := Eval(tc.tree)
		require.NoError(t, err)
		assert.Exactly(t, tc.expected, actual, "\nexpected: %0.3f\ngot:      %0.3f\n%s", ratFloat(tc.expected), ratFloat(actual), tcInfo)
	}
}

func ratFloat(rat *big.Rat) float64 {
	f, _ := rat.Float64()
	return f
}
