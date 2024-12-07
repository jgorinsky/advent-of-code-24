package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type NotFound struct{}

func (m *NotFound) Error() string {
	return "Not Found"
}

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

var man = []rune{'<', '>', '^', 'v'}

type Map [][]rune

func direction(m rune) (Direction, error) {
	switch m {
	case '<':
		return LEFT, nil
	case '>':
		return RIGHT, nil
	case '^':
		return UP, nil
	case 'v':
		return DOWN, nil
	}

	return 0, &NotFound{}
}

func (d Direction) delta() (dx int, dy int) {
	switch d {
	case LEFT:
		return -1, 0
	case RIGHT:
		return 1, 0
	case UP:
		return 0, -1
	case DOWN:
		return 0, 1
	}

	return 0, 0
}

func (d Direction) turn() Direction {
	switch d {
	case LEFT:
		return UP
	case RIGHT:
		return DOWN
	case UP:
		return RIGHT
	case DOWN:
		return LEFT
	}

	return 0
}

func (m Map) whereMan() (x int, y int) {
	for i, row := range m {
		for j, col := range row {
			if slices.Contains(man, col) {
				return j, i
			}
		}
	}

	return -1, -1
}

func (m Map) boundsCheck(x int, y int) bool {
	return x < 0 || x >= len(m[0]) || y < 0 || y >= len(m)
}

func (m Map) move(x int, y int, dir Direction) {
	fmt.Printf("Man at (%v,%v) moving in %v\n", x, y, dir)
	if m.boundsCheck(x, y) {
		return
	}

	m[y][x] = 'X'
	dx, dy := dir.delta()
	if m.boundsCheck(x+dx, y+dy) {
		return
	}
	if m[y+dy][x+dx] == '#' {
		fmt.Println("Turning")
		m.move(x, y, dir.turn())
	} else {
		m.move(x+dx, y+dy, dir)
	}

}

func part1(mapp Map) {
	x, y := mapp.whereMan()
	d, err := direction(mapp[y][x])
	check(err)
	mapp.move(x, y, d)

	sum := 0
	for _, row := range mapp {
		for _, space := range row {
			if space == 'X' {
				sum += 1
			}
		}
	}
	fmt.Printf("%v\n", sum)
}

func part2(mapp [][]rune) {

}
func main() {

	f, err := os.Open("data/day6.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	mapp := Map{}
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, len(line))
		for i, c := range line {
			row[i] = c
		}
		mapp = append(mapp, row)
	}
	part1(mapp)
	part2(mapp)
}
