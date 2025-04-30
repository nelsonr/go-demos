package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func countLines(path string) int {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\r\n")

	return len(lines)
}

func getGoFileLinesCount(path string) int {
	var linesSum int
	var wg sync.WaitGroup

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".go") {
			goFilePath := filepath.Join(path, file.Name())

			// Spawn worker that counts number of lines of a file
			wg.Add(1)

			go func() {
				fmt.Printf("Counting lines of file %s\n", goFilePath)
				linesSum += countLines(goFilePath)

				wg.Done()
			}()
		}
	}

	wg.Wait()

	return linesSum
}

func main() {
	var totalLinesSum int
	var wg sync.WaitGroup

	// Get absolute path to the parent directory
	path, err := filepath.Abs("../")
	if err != nil {
		log.Fatal(err)
	}

	// Get the list of files/folders in the parent directory
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	// Filter to just the directories
	for _, entry := range entries {
		if entry.IsDir() {
			dirPath := filepath.Join(path, entry.Name())

			// Spawn worker to count the sum of lines of each Go file within each directory
			wg.Add(1)

			go func() {
				fmt.Printf("Counting line sum of path: %s\n", dirPath)
				totalLinesSum += getGoFileLinesCount(dirPath)

				wg.Done()
			}()
		}
	}

	wg.Wait()

	fmt.Printf("\nThe total lines count of all Go files is: %d", totalLinesSum)
}
