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

var UP = []int{0, -1}
var DOWN = []int{0, 1}
var LEFT = []int{-1, 0}
var RIGHT = []int{1, 0}
var UPLEFT = []int{-1, -1}
var UPRIGHT = []int{1, -1}
var DOWNLEFT = []int{-1, 1}
var DOWNRIGHT = []int{1, 1}
var ALL = []Direction{UP, DOWN, LEFT, RIGHT, UPLEFT, UPRIGHT, DOWNLEFT, DOWNRIGHT}

type Puzzle [][]rune
type Match = struct {
	x int
	y int
	d Direction
}

type NotFound struct{}

func (m *NotFound) Error() string {
	return "Not Found"
}

func (p Puzzle) search(x int, y int, keyword []rune, directions []Direction) ([]Match, error) {
	if len(keyword) == 0 { // word found
		d := directions[0]
		return []Match{{x + (d[0] * -1), y + (d[1] * -1), directions[0]}}, nil
	}

	if x < 0 || x >= len(p[0]) || y < 0 || y >= len(p) { // out of bounds
		return nil, &NotFound{}
	}

	if p[y][x] == keyword[0] { // potential match
		sum := []Match{}
		for _, d := range directions {
			result, err := p.search(x+d[0], y+d[1], keyword[1:], []Direction{d})
			if err == nil {
				sum = append(sum, result...)
			}

		}
		return sum, nil
	}

	return nil, &NotFound{}

}

var XMAS = []rune{'X', 'M', 'A', 'S'}

func part1(puzzle Puzzle) {
	found := 0
	for y, p := range puzzle {
		for x := range p {
			matches, err := puzzle.search(x, y, XMAS, ALL)
			if err == nil {
				found += len(matches)
			}
		}
	}
	fmt.Println(found)

}

type Coord []int

func (c Coord) Equals(o Coord) bool {
	return c[0] == o[0] && c[1] == o[1]
}

func commonCenters(matches []Match) int {
	centers := []Coord{}
	fmt.Println()

	// find the center of each match (the 'A' in MAS)
	for _, m := range matches {
		x := m.x + (m.d[0] * -1)
		y := m.y + (m.d[1] * -1)
		centers = append(centers, Coord{x, y})
	}

	// count the overlapping centers
	result := 0
	for i, c := range centers {
		for j := range centers {
			if i != j && c.Equals(centers[j]) {
				result += 1
			}
		}
	}

	// divide by two to remove the duplicate matches
	return result / 2
}

var MAS = []rune{'M', 'A', 'S'}
var DIAGS = []Direction{DOWNRIGHT, DOWNLEFT, UPRIGHT, UPLEFT}

func part2(puzzle Puzzle) {
	all := []Match{}
	for y, p := range puzzle {
		for x := range p {
			matches, err := puzzle.search(x, y, MAS, DIAGS)
			if err == nil {
				all = append(all, matches...)
			}
		}
	}
	fmt.Printf("%v\n", commonCenters(all))

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
