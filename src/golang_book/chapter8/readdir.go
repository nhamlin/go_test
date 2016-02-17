package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Open(".")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1) //-1 (or 0) gets all of the files. Any positive number will return up to that many files.
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, fi := range fileInfos {
		fmt.Printf("%-15s | Last Modified: %s\n", fi.Name(), fi.ModTime())
	}
}
