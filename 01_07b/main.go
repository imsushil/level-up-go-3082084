package main

import (
	"flag"
	"log"
)

type operatorType int
const (
	openBracket operatorType = iota
	closedBracket
	otherOperator
)
var bracketsMap = map[rune]rune {
	'{': '}',
	'(': ')',
	'[': ']',
}

func getOperatorType(op rune) operatorType {
	for ob, cb := range bracketsMap {
		switch op {
		case ob:
			return openBracket
		case cb:
			return closedBracket
		}
	}
	return otherOperator
}

// create a stack
type stack struct {
	items []rune
}

func (s *stack) push(op rune) {
	s.items = append(s.items, op)
}

func (s *stack) pop() *rune {
	n := len(s.items) - 1
	last := s.items[n]
	s.items = s.items[:n]
	return &last
}

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
	s := stack{}
	for _, e := range expr {
		switch (getOperatorType(e)) {
		case openBracket:
			s.push(e)
		case closedBracket:
			if len(s.items) == 0 {
				return false
			}
			last := s.pop()
			if last == nil || bracketsMap[*last] != e {
				return false
			}
		}
	}
	return len(s.items) == 0
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool){ 
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}
