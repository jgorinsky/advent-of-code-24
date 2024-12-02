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

func part1(col1 []int, col2 []int) {
	slices.Sort(col1)
	slices.Sort(col2)

	total := 0
	for i := range col1 {
		total += int(math.Abs(float64(col1[i] - col2[i])))
	}

	fmt.Println(total)
}

func part2(col1 []int, col2 []int) {
	total := 0
	for _, v := range col1 {
		count := 0
		for _, v2 := range col2 {
			if v == v2 {
				count += 1
			}
		}
		total += v * count
	}

	fmt.Println(total)
}
func main() {

	f, err := os.Open("data/day1.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	col1 := []int{}
	col2 := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.Fields(line)
		first, err := strconv.Atoi(pair[0])
		check(err)
		second, err := strconv.Atoi(pair[1])
		check(err)
		col1 = append(col1, first)
		col2 = append(col2, second)
	}
	part1(col1, col2)
	part2(col1, col2)
}
