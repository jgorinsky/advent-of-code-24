package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Stack []rune

func (s *Stack) Push(n rune) {
	*s = append(*s, n)
}

func (s *Stack) Pop() rune {
	tmp := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return tmp
}

func (s Stack) Print() {
	fmt.Println(string(s))
}

func part1(sections []string) {
	stack := Stack{}
	for _, section := range sections {
		for _, s := range section {
			stack.Push(s)
		}
	}
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
