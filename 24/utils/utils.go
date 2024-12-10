package utils

import (
	"log"
	"os"
	"strings"
)

func FileParse(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Error in fileParser():", err)
	}

	return string(data)

}

func LineParse(data string) []string {
	data = strings.TrimRight(data, "\n") // remove any extra lines at the end

	return strings.Split(data, "\n")
}

func GridParse(data string) [][]byte {
	lines := LineParse(data)
	grid := make([][]byte, len(lines))
	for i := range grid {
		grid[i] = make([]byte, len(lines[i]))
	}

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			grid[i][j] = lines[i][j]
		}
	}
	return grid
}
