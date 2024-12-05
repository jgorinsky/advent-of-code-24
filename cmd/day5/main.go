package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Pages []int

func (p Pages) find(page int, i int) bool {
	if i < 0 || i >= len(p) {
		return false
	}

	if p[i] == page {
		return true
	}

	return p.find(page, i-1)
}

func validateUpdates(first Rules, updates []Pages, correct bool) []Pages {
	found := []Pages{}
	for _, pages := range updates {
		good := true
		for i, p := range pages {
			if after, ok := first[p]; ok {
				for _, rule := range after {
					if pages.find(rule, i) {
						good = false
					}
				}
			}
		}

		if good && correct {
			found = append(found, pages)
		}

		if !good && !correct {
			found = append(found, pages)
		}
	}

	return found
}

func part1(first Rules, updates []Pages) {

	found := validateUpdates(first, updates, true)
	sum := 0
	for _, c := range found {
		sum += c[len(c)/2]
	}
	fmt.Printf("%v\n", sum)

}

func part2(first Rules, last Rules, updates []Pages) {
	fmt.Println()
	found := validateUpdates(first, updates, false)
	for _, page := range found {
		slices.SortStableFunc(page, func(a, b int) int {
			if after, ok := first[a]; ok {
				if slices.Contains(after, b) {
					return 1
				}
			}
			if before, ok := last[a]; ok {
				if slices.Contains(before, b) {
					return -1
				}
			}
			return 0
		})
	}
	sum := 0
	for _, c := range found {
		sum += c[len(c)/2]
	}

	fmt.Printf("%v\n", sum)
}

type Rules map[int][]int

func (r Rules) Load(key int, val int) {
	if existing, ok := r[key]; ok {
		r[key] = append(existing, val)
	} else {
		r[key] = []int{val}
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
	part1(first, updates)
	part2(first, last, updates)
}
