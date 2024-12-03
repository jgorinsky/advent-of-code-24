package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Stack []Op

func (s *Stack) Push(n Op) {
	*s = append(*s, n)
}

func (s *Stack) Pop() Op {
	tmp := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return tmp
}

func (s Stack) Print() {
	fmt.Printf("%v\n", s)
}

var pattern = regexp.MustCompile(`(mul|\(|\)|\d+|,){1}`)

func tokenize(s string) []string {
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		loc := pattern.FindIndex(data)
		if loc == nil {
			return 0, nil, bufio.ErrFinalToken
		} else if loc[0] == 0 { // Next token is valid
			return loc[1], data[0:loc[1]], nil
		} else { // Next token is invalid
			return loc[0], data[0:loc[0]], nil
		}
	}

	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(split)

	result := []string{}
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}

var ops = []string{"mul"}

type Op = struct {
	op string
	x  int
	y  int
}

func parse(tokens []string, ops *Stack) *Stack {
	if len(tokens) == 0 {
		return ops
	}
	op, err := parseOp(tokens)
	if err != nil {
		return parse(tokens[1:], ops)
	} else {
		ops.Push(*op)
		return parse(tokens[6:], ops)
	}
}

func parseOp(tokens []string) (*Op, error) {
	op := tokens[0]
	if !slices.Contains(ops, op) {
		return nil, errors.New("invalid op name")
	}
	if tokens[1] != "(" {
		return nil, errors.New("invalid op: no open paren")
	}
	arg1, arg2, err := parseArgs(tokens[2:])
	if err != nil {
		return nil, err
	}

	if tokens[5] != ")" {
		return nil, fmt.Errorf("invalid op: no close paren %v", tokens[2])
	}

	return &Op{op, arg1, arg2}, nil

}

func parseArgs(tokens []string) (int, int, error) {
	arg1, err := strconv.Atoi(tokens[0])

	if err != nil {
		return 0, 0, err
	}

	if tokens[1] != "," {
		return 0, 0, fmt.Errorf("invalid arguments: no comma in token %v", tokens[1])
	}

	arg2, err := strconv.Atoi(tokens[2])
	if err != nil {
		return 0, 0, err
	}

	return arg1, arg2, nil

}

func part1(sections []string) {
	ops := Stack{}
	for _, section := range sections {
		parse(tokenize(section), &ops)
	}
	sum := 0
	for _, op := range ops {
		sum += (op.x * op.y)
	}
	fmt.Printf("%v\n", sum)
}

func part2(sections []string) {

}
func main() {

	f, err := os.Open("data/day3.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	part1(lines)
	part2(lines)
}
