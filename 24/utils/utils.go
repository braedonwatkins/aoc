package utils

import (
	"log"
	"os"
)

func FileParser(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Error in fileParser():", err)
	}

	return string(data)

}
