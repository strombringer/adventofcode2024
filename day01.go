package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	s "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Preparation
	var input, err = os.ReadFile("./inputs/day01")
	check(err)

	var lines = s.Split(string(input), "\n")

	var leftNumbers = make([]int, len(lines))
	var rightNumbers = make([]int, len(lines))

	for i, line := range lines {
		var entry = s.Fields(line)
		if len(entry) == 0 {
			break
		}

		leftNumbers[i], _ = strconv.Atoi(entry[0])
		rightNumbers[i], _ = strconv.Atoi(entry[1])
	}

	slices.Sort(leftNumbers)
	slices.Sort(rightNumbers)

	// Part 1
	var sum = 0
	for i := range leftNumbers {
		var diff = leftNumbers[i] - rightNumbers[i]
		sum = sum + Abs(diff)
	}
	fmt.Println("Answer Part 1: ", sum)

	// Part 2
	var similarityMap = make(map[int]int)
	for _, entry := range rightNumbers {
		var currentValue = similarityMap[entry]
		similarityMap[entry] = currentValue + 1
	}

	var similaritySum = 0
	for _, entry := range leftNumbers {
		var similarity = entry * similarityMap[entry]
		similaritySum = similaritySum + similarity
	}
	fmt.Println("Answer Part 2: ", similaritySum)
}
