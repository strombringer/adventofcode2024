package main

import (
	"os"
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

func StringToInt(input string) int {
	var result, _ = strconv.Atoi(input)
	return result
}

func IsSameSign(x int, y int) bool {
	return (x >= 0) == (y >= 0)
}

func RemoveIndex(s []string, index int) []string {
	ret := make([]string, 0, len(s)-1)
	ret = append(ret, s[:index]...)
	ret = append(ret, s[index+1:]...)

	return ret
}

func GetInputData(path string) []string {
	var input, err = os.ReadFile(path)
	check(err)

	var lines = s.Split(string(input), "\n")
	return lines
}

func IsEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
