package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

const AFTER = 1
const BEFORE = -1

type Pages []int

func (p Pages) find(page int, i int, dir int) bool {
	if i < 0 || i >= len(p) {
		return false
	}

	if p[i] == page {
		return true
	}

	return p.find(page, i+dir, dir)
}

func part1(first Rules, last Rules, updates []Pages) {
	fmt.Printf("%v\n", first)
	fmt.Printf("%v\n", last)
	fmt.Printf("%v\n", updates)

	correct := []Pages{}
	for _, pages := range updates {
		good := true
		for i, p := range pages {
			after, ok := first[p]
			if ok {
				for _, rule := range after {
					if pages.find(rule, i, BEFORE) {
						good = false
					}
				}
			}

			before, ok := last[p]
			if ok {
				for _, rule := range before {
					if pages.find(rule, i, AFTER) {
						good = false
					}
				}
			}
		}

		if good {
			correct = append(correct, pages)
		}
	}

	sum := 0
	for _, c := range correct {
		sum += c[len(c)/2]
	}

	fmt.Printf("%v\n", correct)
	fmt.Printf("%v\n", sum)

}

func part2(first Rules, last Rules, updates []Pages) {

}

type Rules map[int][]int

func (r *Rules) Load(key int, val int) {
	if existing, ok := (*r)[key]; ok {
		(*r)[key] = append(existing, val)
	} else {
		(*r)[key] = []int{val}
	}
}
func main() {

	f, err := os.Open("data/day5.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	first := Rules{}
	last := Rules{}
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}
		line := scanner.Text()
		vals := strings.Split(line, "|")
		f, err := strconv.Atoi(vals[0])
		check(err)
		l, err := strconv.Atoi(vals[1])
		check(err)
		first.Load(f, l)
		last.Load(l, f)
	}

	updates := []Pages{}
	for scanner.Scan() {
		line := scanner.Text()
		pagesStr := strings.Split(line, ",")
		pages := make([]int, len(pagesStr))
		for i, page := range pagesStr {
			pageNum, err := strconv.Atoi(page)
			check(err)
			pages[i] = pageNum
		}
		updates = append(updates, pages)
	}
	part1(first, last, updates)
	part2(first, last, updates)
}
