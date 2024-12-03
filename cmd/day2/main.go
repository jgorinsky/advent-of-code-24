package main

import (
	"bufio"
	"fmt"
	"math"
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

type Rule func([]int) bool

func increasing(n int, m int) bool {
	return n-m < 0
}

func decreasing(n int, m int) bool {
	return !increasing(n, m)
}

func incdec(report []int) bool {
	result := true

	inc := increasing(report[0], report[1])

	for i, n := range report[:len(report)-1] {
		test := increasing(n, report[i+1])
		if test && !inc {
			result = false
		} else if decreasing(n, report[i+1]) && inc {
			result = false
		}
	}

	return result
}

func span(report []int) bool {
	result := true

	for i, level := range report[:len(report)-1] {
		span := int(math.Abs(float64(level - report[i+1])))
		if span == 0 || span > 3 {
			result = false
		}
	}

	// fmt.Printf("span report %v %v\n", report, result)
	return result
}

func permute(report []int) [][]int {
	reports := [][]int{report}
	for i := range report {
		tmp := make([]int, len(report))
		copy(tmp, report)
		reports = append(reports, slices.Concat(tmp[:i], tmp[i+1:]))
	}
	return reports
}

func part1(reports [][]int) {

	safe := 0

	rules := []Rule{incdec, span}

	for _, report := range reports {
		isSafe := true
		for _, rule := range rules {
			if !rule(report) {
				isSafe = false
			}
		}
		if isSafe {
			safe += 1
		}
	}

	fmt.Println(safe)
}

func part2(reports [][]int) {
	safeCount := 0

	for _, r := range reports {
		safe := incdec(r) && span(r)
		if !safe {
			for _, t := range permute(r) {
				if incdec(t) && span(t) {
					safeCount += 1
					break
				}
			}
		} else {
			safeCount += 1
		}

	}

	fmt.Println(safeCount)
}
func main() {

	f, err := os.Open("data/day2.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	reports := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		reportStr := strings.Fields(line)
		report := make([]int, len(reportStr))
		for i, v := range reportStr {
			level, err := strconv.Atoi(v)
			check(err)
			report[i] = level
		}
		reports = append(reports, report)
	}
	part1(reports)
	part2(reports)
}
