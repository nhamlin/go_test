package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.WriteString("This is a test. This is only a test.\n\n\nThis is a test of the EBN.")
}
