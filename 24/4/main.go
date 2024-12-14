package main

import (
	"aoc/utils"
	"fmt"
)

/* NOTE:


 */

var diagonalDir = [][]int{
	{-1, 1},  //NE
	{1, 1},   //SE
	{1, -1},  //SW
	{-1, -1}, //NW
}

var cardinalDir = [][]int{
	{-1, 0}, //N
	{0, 1},  //E
	{1, 0},  //S
	{0, -1}, //W
}

var word_1 = []byte{'X', 'M', 'A', 'S'}
var word_2 = []byte{'M', 'A', 'S'}

func findMAS(grid [][]byte, startRow int, startCol int, numRows int, numCols int) bool {

	if grid[startRow][startCol] != 'A' || startRow-1 < 0 || startRow+1 > numRows-1 || startCol-1 < 0 || startCol+1 > numCols-1 {
		return false
	}

	/*
	   M . M
	   . A .
	   S . S
	*/
	isTopM := grid[startRow-1][startCol-1] == 'M' && grid[startRow-1][startCol+1] == 'M' && grid[startRow+1][startCol-1] == 'S' && grid[startRow+1][startCol+1] == 'S'
	isLeftM := grid[startRow-1][startCol-1] == 'M' && grid[startRow+1][startCol-1] == 'M' && grid[startRow-1][startCol+1] == 'S' && grid[startRow+1][startCol+1] == 'S'
	isRightM := grid[startRow-1][startCol+1] == 'M' && grid[startRow+1][startCol+1] == 'M' && grid[startRow-1][startCol-1] == 'S' && grid[startRow+1][startCol-1] == 'S'
	isBotM := grid[startRow+1][startCol-1] == 'M' && grid[startRow+1][startCol+1] == 'M' && grid[startRow-1][startCol+1] == 'S' && grid[startRow-1][startCol-1] == 'S'

	if isTopM || isLeftM || isRightM || isBotM {
		return true
	}

	return false
}

func findXMAS(grid [][]byte, dir []int, startRow int, startCol int, numRows int, numCols int) bool {
	wordLen := len(word_1)
	yDir := dir[0]
	xDir := dir[1]

	xMax := (wordLen - 1) * xDir
	yMax := (wordLen - 1) * yDir

	if startRow+yMax < 0 || startRow+yMax > numRows-1 || startCol+xMax < 0 || startCol+xMax > numCols-1 {
		return false
	}

	for i, letter := range word_1 {
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

	//NOTE: this is just here for abstraction into the second part
	var directions = [][]int{}
	directions = append(directions, diagonalDir...)
	directions = append(directions, cardinalDir...)

	grid := utils.GridParse(data)
	numRows := len(grid)
	numCols := len(grid[0]) // assuming constant grid size
	count := 0

	// fmt.Println(numRows, numCols)
	// for i, row := range grid {
	// 	fmt.Printf("Row %d: %v\n", i, row)
	// }

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			// fmt.Print(string(grid[i][j]))
			for _, dir := range directions {
				isWord := findXMAS(grid, dir, i, j, numRows, numCols)
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
	grid := utils.GridParse(data)
	numRows := len(grid)
	numCols := len(grid[0]) // assuming constant grid size
	count := 0

	// fmt.Println(numRows, numCols)
	// for i, row := range grid {
	// 	fmt.Printf("Row %d: %v\n", i, row)
	// }

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			// fmt.Print(string(grid[i][j]))
			isWord := findMAS(grid, i, j, numRows, numCols)
			if isWord {
				// fmt.Print("*")
				count++
			}
		}

		// fmt.Print("\n")
	}

	return count
}

func main() {
	data := utils.FileParse("input.txt")
	// ans := partOne(data)
	ans := partTwo(data)
	fmt.Println(ans)
}
