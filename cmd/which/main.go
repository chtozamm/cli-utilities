package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: which [filename ...]")
		return
	}

	files := os.Args[1:]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	// Iterate through each filename provided
	for _, file := range files {
		found := false

		// Iterate through each directory in the PATH
		for _, dir := range pathSplit {
			fullPath := filepath.Join(dir, file)

			// Append ".exe" extension for Windows if necessary
			if runtime.GOOS == "windows" {
				if !strings.HasSuffix(fullPath, ".exe") {
					fullPath += ".exe"
				}
			}

			if fileExists(fullPath) {
				fmt.Println(fullPath)
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("'%s' not found in PATH\n", file)
		}
	}
}

// fileExists checks if a file exists at a given path.
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
