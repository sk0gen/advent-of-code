package utils

import (
	"log"
	"os"
	"strings"
)

func ReadLines(filePath string) []string {
	file, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")

	return lines
}
