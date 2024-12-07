package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"sync"
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

type Loop struct{}

func (m *Loop) Error() string {
	return "Infinite Loop"
}

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

var man = []rune{'<', '>', '^', 'v'}

type Position struct {
	val     rune
	history []Direction
}

type Map [][]Position

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

func (m Map) clone() Map {
	newMap := make(Map, len(m))
	for i, row := range m {
		newMap[i] = make([]Position, len(row))
		copy(newMap[i], row)
	}
	return newMap
}

func (m Map) whereMan() (x int, y int) {
	for i, row := range m {
		for j, col := range row {
			if slices.Contains(man, col.val) {
				return j, i
			}
		}
	}

	return -1, -1
}

func (m Map) boundsCheck(x int, y int) bool {
	return x < 0 || x >= len(m[0]) || y < 0 || y >= len(m)
}

func (m Map) move(x int, y int, dir Direction) error {
	// fmt.Printf("Man at (%v,%v) moving in %v\n", x, y, dir)
	if m.boundsCheck(x, y) {
		return nil
	}

	if slices.Contains(m[y][x].history, dir) {
		return &Loop{}
	}

	m[y][x].val = 'X'
	m[y][x].history = append(m[y][x].history, dir)
	dx, dy := dir.delta()
	if m.boundsCheck(x+dx, y+dy) {
		return nil
	}
	lookAhead := m[y+dy][x+dx].val
	if lookAhead == '#' || lookAhead == 'O' {
		// fmt.Println("Turning")
		return m.move(x, y, dir.turn())
	} else {
		return m.move(x+dx, y+dy, dir)
	}

}

func (m Map) String() string {
	var b strings.Builder
	for _, row := range m {
		for _, pos := range row {
			fmt.Fprintf(&b, "%v", string(pos.val))
		}
		fmt.Fprintln(&b)
	}
	return b.String()
}

func part1(mapp Map) {
	x, y := mapp.whereMan()
	d, err := direction(mapp[y][x].val)
	check(err)
	mapp.move(x, y, d)

	sum := 0
	for _, row := range mapp {
		for _, space := range row {
			if space.val == 'X' {
				sum += 1
			}
		}
	}
	fmt.Printf("%v\n", sum)
}

func part2(mapp Map) {
	x, y := mapp.whereMan()
	d, err := direction(mapp[y][x].val)
	check(err)

	loops := make(chan int, len(mapp)*len(mapp[0]))
	var wg sync.WaitGroup
	for row := range len(mapp) {
		for pos := range len(mapp[0]) {
			if row == y && pos == x {
				continue
			}

			if mapp[y][x].val == '#' {
				continue
			}

			instance := mapp.clone()
			instance[row][pos].val = 'O'

			wg.Add(1)
			go func() {
				defer wg.Done()
				err := instance.move(x, y, d)

				if err != nil {
					loops <- 1
				}
			}()

		}
	}

	wg.Wait()
	close(loops)
	count := 0
	for range loops {
		count++
	}
	fmt.Printf("%v\n", count)

}
func main() {

	f, err := os.Open("data/day6.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	mapp := Map{}
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]Position, len(line))
		for i, c := range line {
			row[i] = Position{val: c, history: []Direction{}}
		}
		mapp = append(mapp, row)
	}
	part1(mapp.clone())
	part2(mapp.clone())
}
