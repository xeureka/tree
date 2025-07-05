package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var excluded map[string]bool

func main() {
	var excludFiles string
	flag.StringVar(&excludFiles, "x", "", "files to be excluded from view")
	flag.Parse()

	excluded = make(map[string]bool)
	processExcludedFiles(excludFiles)

	dir, _ := os.Getwd()
	explorePath(dir, 0)
}

func explorePath(path string, level int) {
	entries, err := os.ReadDir(path)
	entries = sortDir(entries)
	if err != nil {
		return
	}

	for _, entry := range entries {
		if excluded[entry.Name()] {
			continue
		}
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
	var sb strings.Builder
	for range space {
		sb.WriteString("|")
		sb.WriteString("   ")
	}
	return sb.String()
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

func processExcludedFiles(excludeFiles string) {
	// some hardcoded things
	excluded[".git"] = true
	excluded["node_modules"] = true

	files := strings.Split(excludeFiles, ",")
	for _, file := range files {
		excluded[file] = true
	}
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
