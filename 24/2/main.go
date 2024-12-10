package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func parseLines(data string) []string {
	return strings.Split(data, "\n")
}

func abs(num int) int {
	if num < 0 {
		num *= -1
	}

	return num
}

func partOne(data string) int {
	lines := parseLines(data)
	// fmt.Printf("%q ", lines)

	numSafe := 0
	for _, line := range lines {
		elements := strings.Split(line, " ")
		if len(elements) < 2 {
			continue
		}

		firstNum, _ := strconv.Atoi(elements[0])
		secondNum, _ := strconv.Atoi(elements[1])

		if secondNum > firstNum {
			var i int
			for i = 0; i < len(elements)-1; i++ {
				curNum, _ := strconv.Atoi(elements[i])
				nextNum, _ := strconv.Atoi(elements[i+1])
				absDiff := abs(curNum - nextNum)

				if absDiff > 3 || absDiff < 1 || nextNum < curNum {
					break
				}
			}
			// fmt.Println(i)
			if i >= len(elements)-1 {
				numSafe++
			}
		} else {
			var i int
			for i = 0; i < len(elements)-1; i++ {
				curNum, _ := strconv.Atoi(elements[i])
				nextNum, _ := strconv.Atoi(elements[i+1])
				absDiff := abs(curNum - nextNum)

				if absDiff > 3 || absDiff < 1 || nextNum > curNum {
					break
				}
			}
			// fmt.Println(i)
			if i >= len(elements)-1 {
				numSafe++
			}

		}
	}

	return numSafe
}

// FIXME: for some reason this fails the test but passes the input ??
func partTwo(data string) int {
	lines := parseLines(data)
	// fmt.Printf("%q ", lines)

	numSafe := 0
	for _, line := range lines {
		isProbDampened := false

		elements := strings.Split(line, " ")
		if len(elements) < 2 {
			continue
		}

		firstNum, _ := strconv.Atoi(elements[0])
		secondNum, _ := strconv.Atoi(elements[1])

		if secondNum > firstNum {
			var i int
			for i = 0; i < len(elements)-1; i++ {
				curNum, _ := strconv.Atoi(elements[i])
				nextNum, _ := strconv.Atoi(elements[i+1])
				absDiff := abs(curNum - nextNum)

				if absDiff > 3 || absDiff < 1 || nextNum < curNum {
					if !isProbDampened {
						isProbDampened = true
					} else {
						break
					}
				}
			}
			// fmt.Println(i)
			if i >= len(elements)-1 {
				numSafe++
			}
		} else {
			var i int
			for i = 0; i < len(elements)-1; i++ {
				curNum, _ := strconv.Atoi(elements[i])
				nextNum, _ := strconv.Atoi(elements[i+1])
				absDiff := abs(curNum - nextNum)

				if absDiff > 3 || absDiff < 1 || nextNum > curNum {
					if !isProbDampened {
						isProbDampened = true
					} else {
						break
					}
				}
			}
			// fmt.Println(i)
			if i >= len(elements)-1 {
				numSafe++
			}

		}
	}

	return numSafe

}

func main() {
	data := utils.FileParse("input.txt")
	ans1 := partOne(data)
	ans2 := partTwo(data)
	fmt.Println(ans1, ans2)
}
