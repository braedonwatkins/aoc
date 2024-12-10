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
