package utils

import (
	"bufio"
	"log"
	"os"
)

func MapFileLines(filePath string, cb func(line string)) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		cb(line)
	}

	if scannerErr := scanner.Err(); scannerErr != nil {
		log.Fatal(scannerErr)
	}
}
