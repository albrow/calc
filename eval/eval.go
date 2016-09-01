package eval

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"github.com/albrow/calc/ast"
)

func Eval(tree ast.Node) (*big.Rat, error) {
	switch node := tree.(type) {
	case *ast.Number:
		return parseNumNode(node)
	case *ast.BaseNode:
		return evalNodes(tree.Children())
	default:
		return nil, fmt.Errorf("Unkown node type: %T", tree)
	}
}

func evalNodes(nodes []ast.Node) (*big.Rat, error) {
	if len(nodes) == 0 {
		return nil, errors.New("eval.evalNodes: cannot eval nodes of length 0")
	}
	var val *big.Rat
	switch first := nodes[0].(type) {
	case *ast.Number:
		num, err := parseNumNode(first)
		if err != nil {
			return nil, err
		}
		val = num
	case *ast.BaseNode:
		num, err := evalNodes(first.Children())
		if err != nil {
			return nil, err
		}
		val = num
	}
	var lastOp ast.OpClass
	for _, node := range nodes[1:] {
		switch n := node.(type) {
		case *ast.Operator:
			lastOp = n.Class
		case *ast.Number:
			operand, err := parseNumNode(n)
			if err != nil {
				return nil, err
			}
			applyOp(val, lastOp, operand)
		case *ast.BaseNode:
			operand, err := evalNodes(n.Children())
			if err != nil {
				return nil, err
			}
			applyOp(val, lastOp, operand)
		}
	}
	return val, nil
}

func parseNumNode(node *ast.Number) (*big.Rat, error) {
	i, err := strconv.ParseInt(node.Value, 10, 64)
	if err != nil {
		return nil, err
	}
	return big.NewRat(i, 1), nil
}

func applyOp(val *big.Rat, op ast.OpClass, operand *big.Rat) {
	switch op {
	case ast.OpAdd:
		val.Add(val, operand)
	case ast.OpSubtract:
		val.Sub(val, operand)
	default:
		panic(fmt.Sprintf("eval.applyOp: unkown operand: %d (%s)", op, op))
	}
}
