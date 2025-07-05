package main

import (
	"fmt"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	explorePath(dir, 0)
}

func explorePath(path string, level int) {
	entries, err := os.ReadDir(path)
	entries = sortDir(entries)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, entry := range entries {
		decorator := ""
		if entry.IsDir() {
			decorator = "\\ "
		} else {
			decorator = "- "
		}

		fmt.Println(indentation(level) + decorator + entry.Name())
		newPath := path + "/" + entry.Name()
		if entry.IsDir() {
			explorePath(newPath, level+1)
		}
	}
}

func indentation(space int) string {
	padding := ""
	for range space {
		padding += "|"
		padding += "   "
	}
	return padding
}

func sortDir(entries []os.DirEntry) []os.DirEntry {
	var sortedDirEntry []os.DirEntry
	for _, entry := range entries {
		if entry.IsDir() {
			sortedDirEntry = append(sortedDirEntry, entry)
		}
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			sortedDirEntry = append(sortedDirEntry, entry)
		}
	}

	return sortedDirEntry
}

// \ another_folder
// |	\ binaryes
// |	|	\ another_folder
// |	|	|	- big file
// | 	|	|	- very big file
// |	| 	- leveler
// | 	\ nested
// |	| 	- file_name
// | 	| 	- test
