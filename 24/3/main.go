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
	byteData := []byte(data)

	//PERF: ideally we do this in one pass instead of 3... O(n) is O(n) ig lol
	multIndeces := multRegex.FindAllIndex(byteData, -1)
	doIndeces := doRegex.FindAllIndex(byteData, -1)
	dontIndeces := dontRegex.FindAllIndex(byteData, -1)

	fmt.Println(multIndeces, doIndeces, dontIndeces)

	//2. process multiplications
	total := 0
	// for _, match := range allMatches {
	// 	total += multOp(match)
	// }

	return total

}

func main() {
	data := utils.FileParser("input.txt")
	ans := partOne(data)
	// ans := partTwo(data)
	fmt.Println(ans)
}
