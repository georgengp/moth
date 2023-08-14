package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// chkErr will simply print the error and terminates the program.
func chkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// byteSizeAutorange returns human-readable values (K, M, G) given a size in bytes.
// humanizeByteSize could be a better name?
func byteSizeAutorange(size int64) string {
	switch {
	case size < 1024:
		return fmt.Sprintf("%v", size)
	case size < 10 * 1024:
		return fmt.Sprintf("%.1fK", float64(size) / 1024)
	case size < 1024 * 1024:
		return fmt.Sprintf("%vK", size / 1024)
	case size < 10 * 1024 * 1024:
		return fmt.Sprintf("%.1fM", float64(size) / (1024 * 1024))
	case size < 1024 * 1024 * 1024:
		return fmt.Sprintf("%vM", size / (1024 * 1024))
	case size < 10 * 1024 * 1024 * 1024:
		return fmt.Sprintf("%.1fG", float64(size) / (1024 * 1024 * 1024))
	case size < 1024 * 1024 * 1024 * 1024:
		return fmt.Sprintf("%vG", size / (1024 * 1024 * 1024))
	default:
		return string(size)
	}
}

func main() {
	cwd, err := os.Getwd() // Gets cwd. Where shell is pointing at!
	chkErr(err)

	dirEntries, err := os.ReadDir(cwd) // Acquire entries about the files and subdirectories inside.
	chkErr(err)
	
	// Sorting case-insensitively (ls behavior)
	sort.Slice(dirEntries, func(i, j int) bool {
		return strings.ToLower(dirEntries[i].Name()) < strings.ToLower(dirEntries[j].Name())
	})

	// Printing out the sizes
	for _, entry := range dirEntries {
		entryInfo, err := entry.Info()
		chkErr(err)

		// Subdirectories are handled normally, apparently.
		// Some rounding discrepancy between this tool and ls.
		fmt.Printf("%4v %v\n", byteSizeAutorange(entryInfo.Size()), entryInfo.Name())
	}
}
