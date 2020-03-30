// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	linesMappedToFilenames := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLinesAndMapToFile(os.Stdin, counts, linesMappedToFilenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLinesAndMapToFile(f, counts, linesMappedToFilenames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Println("Found in:")
			for _, fileName := range linesMappedToFilenames[line] {
				fmt.Printf("\t%s\n", fileName)
			}
		}
	}
}

func countLinesAndMapToFile(f *os.File, counts map[string]int, linesMappedToFilenames map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		countLine(input.Text(), counts)
		mapLineToFile(input.Text(), f.Name(), linesMappedToFilenames)
	}
	// NOTE: ignoring potential errors from input.Err()
}

func countLine(line string, counts map[string]int) {
	counts[line]++
}

func mapLineToFile(line string, fileName string, linesMappedToFilenames map[string][]string) {
	if !fileNameExists(fileName, linesMappedToFilenames[line]) {
		linesMappedToFilenames[line] = append(linesMappedToFilenames[line], fileName)
	}
}

func fileNameExists(fileName string, lineMappedToFilenames []string) bool {
	for _, s := range lineMappedToFilenames {
		if fileName == s {
			return true
		}
	}
	return false
}
