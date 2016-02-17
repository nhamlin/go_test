package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		// if path != "." && info.IsDir() {
		// 	return filepath.SkipDir
		// }
		if info.IsDir() { // if the file is actually a directory entry, suffix the name with a slash
			fmt.Printf("%-35s | Last Modified: %s\n", path+"/", info.ModTime())
		} else {
			fmt.Printf("%-35s | Last Modified: %s\n", path, info.ModTime())
		}
		return nil
	})
}
