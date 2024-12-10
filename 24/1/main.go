package main

import (
	"aoc/utils"
	"slices"
	"strconv"
	"strings"
)

func sortStringNums(a string, b string) int {
	//NOTE: ignoring errors, possible edge case maybe
	aNum, _ := strconv.Atoi(a)
	bNum, _ := strconv.Atoi(b)

	if aNum > bNum {
		return 1
	} else {
		return -1
	}
}

func parseSpaceSeparatedLists(listString string) [][]string {
	//1. parse into two separate []string
	lists := strings.Fields(listString)
	// fmt.Println(lists)

	leftList := []string{}
	rightList := []string{}

	//NOTE: assuming lists are same length
	for i := 0; i < len(lists); i++ {
		if i%2 == 0 {
			leftList = append(leftList, lists[i])
		} else {
			rightList = append(rightList, lists[i])

		}
	}

	return [][]string{leftList, rightList}
}

func partOne(data string) int {
	lists := parseSpaceSeparatedLists(data)
	leftList, rightList := lists[0], lists[1]

	// fmt.Println(leftList)
	// fmt.Println(rightList)

	//2. sort each []string
	slices.SortFunc(leftList, sortStringNums)
	slices.SortFunc(rightList, sortStringNums)

	//3. for less than combined length calculate distance (like .reduce)
	totalDistance := 0
	for i := 0; i < len(leftList) && i < len(rightList); i++ {
		//NOTE: ignoring errors, possible edge case maybe
		left, _ := strconv.Atoi(leftList[i])
		right, _ := strconv.Atoi(rightList[i])

		difference := (left - right)
		if difference < 0 {
			difference *= -1
		}

		totalDistance += difference
	}

	return totalDistance
}

func partTwo(data string) int {
	lists := parseSpaceSeparatedLists(data)
	leftList, rightList := lists[0], lists[1]

	rightSet := make(map[string]int)
	for _, str := range rightList {
		rightSet[str] += 1
	}

	totalDistance := 0
	for i := 0; i < len(leftList) && i < len(rightList); i++ {
		left := leftList[i]
		if val, hasVal := rightSet[left]; hasVal {
			leftNum, _ := strconv.Atoi(left)
			totalDistance += leftNum * val
		}
	}
	return totalDistance
}

func main() {
	_ = utils.FileParse("test.txt")
	// ans := partOne(data)
	// ans := partTwo(data)
	// fmt.Println(ans)
}
