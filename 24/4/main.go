package main

import (
	"aoc/utils"
	"fmt"
)

/* NOTE:


 */

var directions = [][]int{
	{-1, 0},  //N
	{-1, 1},  //NE
	{0, 1},   //E
	{1, 1},   //SE
	{1, 0},   //S
	{1, -1},  //SW
	{0, -1},  //W
	{-1, -1}, //NW
}

var word = []byte{'X', 'M', 'A', 'S'}

func findWord(grid [][]byte, dir []int, startRow int, startCol int, numRows int, numCols int) bool {
	wordLen := len(word)
	yDir := dir[0]
	xDir := dir[1]

	xMax := (wordLen - 1) * xDir
	yMax := (wordLen - 1) * yDir

	if startRow+yMax < 0 || startRow+yMax > numRows-1 || startCol+xMax < 0 || startCol+xMax > numCols-1 {
		return false
	}

	for i, letter := range word {
		r := startRow + yDir*i
		c := startCol + xDir*i

		// fmt.Println(r, c, numRows, numCols)

		if grid[r][c] != letter {
			return false
		}
	}

	return true
}

func partOne(data string) int {
	// fmt.Println(data)
	// fmt.Println(grid)

	grid := utils.GridParse(data)
	numRows := len(grid)
	numCols := len(grid[0]) // assuming constant grid size
	count := 0

	// fmt.Println(numRows, numCols)
	for i, row := range grid {
		fmt.Printf("Row %d: %v\n", i, row)
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			// fmt.Print(string(grid[i][j]))
			for _, dir := range directions {
				isWord := findWord(grid, dir, i, j, numRows, numCols)
				if isWord {
					// fmt.Print("*")
					count++
				}
			}
		}

		// fmt.Print("\n")
	}

	return count
}

func partTwo(data string) int {

	return -1
}

func main() {
	data := utils.FileParse("input.txt")
	ans := partOne(data)
	// ans := partTwo(data)
	fmt.Println(ans)
}
