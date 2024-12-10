package main

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// takes in string of form `mul(\d,\d)`
func multOp(multStr string) int {
	leftParenIndex := strings.Index(multStr, "(")
	commaIndex := strings.Index(multStr, ",")
	rightParenIndex := strings.Index(multStr, ")")

	str1 := multStr[leftParenIndex+1 : commaIndex]
	str2 := multStr[commaIndex+1 : rightParenIndex]

	arg1, _ := strconv.Atoi(string(str1))
	arg2, _ := strconv.Atoi(string(str2))

	// fmt.Println(leftParenIndex, commaIndex, rightParenIndex)
	// fmt.Println(str1, str2)

	return arg1 * arg2
}

var multRegex = regexp.MustCompile(`mul\(\d+,\d+\)`)
var doRegex = regexp.MustCompile(`do\(\)`)
var dontRegex = regexp.MustCompile(`don't\(\)`)

func partOne(data string) int {
	//1. get all instances of correctly formatted mul
	allMatches := multRegex.FindAllString(data, -1)

	// fmt.Println(allMatches)

	//2. process multiplications
	total := 0
	for _, match := range allMatches {
		total += multOp(match)
	}

	return total
}

func partTwo(data string) int {
	isEnabled := true
	byteData := []byte(data)
	total := 0

	for i := 0; i < len(byteData); {
		if !isEnabled {
			nextDoIndeces := doRegex.FindIndex(byteData[i:])
			if nextDoIndeces == nil {
				return total
			}
			i += nextDoIndeces[1]
			isEnabled = true
		} else {
			nextMultIndeces := multRegex.FindIndex(byteData[i:])
			nextDontIndeces := dontRegex.FindIndex(byteData[i:])

			if nextMultIndeces == nil {
				return total
			}

			if nextDontIndeces == nil || nextMultIndeces[1] < nextDontIndeces[1] {
				multStr := multRegex.Find(byteData[i:])
				total += multOp(string(multStr))
				i += nextMultIndeces[1]
			} else {
				i += nextDontIndeces[1]
				isEnabled = false
			}
		}
	}
	return total
}

func main() {
	data := utils.FileParse("input.txt")
	// ans := partOne(data)
	ans := partTwo(data)
	fmt.Println(ans)
}
