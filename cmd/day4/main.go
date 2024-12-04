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

type Direction = []int

var UP = []int{0, 1}
var DOWN = []int{0, -1}
var LEFT = []int{-1, 0}
var RIGHT = []int{1, 0}
var UPLEFT = []int{-1, 1}
var UPRIGHT = []int{1, 1}
var DOWNLEFT = []int{-1, -1}
var DOWNRIGHT = []int{1, -1}
var ALL = []Direction{UP, DOWN, LEFT, RIGHT, UPLEFT, UPRIGHT, DOWNLEFT, DOWNRIGHT}

type Puzzle [][]rune

func (p Puzzle) search(x int, y int, keyword []rune, directions []Direction) int {
	if len(keyword) == 0 { // word found
		fmt.Printf("Found (%v, %v)\n", x, y)
		return 1
	}

	if x < 0 || x >= len(p[0]) || y < 0 || y >= len(p) { // out of bounds
		fmt.Printf("OOB (%v, %v)\n", x, y)
		return 0
	}

	if p[y][x] == keyword[0] { // potential match
		fmt.Printf("HIT (%v, %v) = %v\n", x, y, p[y][x])
		sum := 0
		for _, d := range directions {
			sum += p.search(x+d[0], y+d[1], keyword[1:], []Direction{d})
		}
		return sum
	}

	return 0

}

var XMAS = []rune{'X', 'M', 'A', 'S'}

func part1(puzzle Puzzle) {
	found := 0
	for y, p := range puzzle {
		for x := range p {
			found += puzzle.search(x, y, XMAS, ALL)
		}
	}
	fmt.Println(found)

}

func part2(puzzle [][]rune) {

	// fmt.Println(puzzle)
}
func main() {

	f, err := os.Open("data/day4.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	puzzle := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, len(line))
		for i, c := range line {
			row[i] = c
		}
		puzzle = append(puzzle, row)
	}
	part1(puzzle)
	part2(puzzle)
}
