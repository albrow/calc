package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/albrow/calc/eval"
	"github.com/albrow/calc/lex"
	"github.com/albrow/calc/parse"
)

func main() {
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		result, err := parseAndEval(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result.RatString())
		fmt.Print("> ")
	}
}

func parseAndEval(input string) (*big.Rat, error) {
	tokens, err := lex.Lex([]byte(input))
	if err != nil {
		return nil, err
	}
	tree, err := parse.Parse(tokens)
	if err != nil {
		return nil, err
	}
	result, err := eval.Eval(tree)
	if err != nil {
		return nil, err
	}
	return result, nil
}
