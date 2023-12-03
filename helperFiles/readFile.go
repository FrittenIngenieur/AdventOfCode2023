package helperFiles

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadInput(path string) (lines []string) {
	// Read File
	fmt.Println("Read input file")
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	file.Close()

	return fileLines
}
