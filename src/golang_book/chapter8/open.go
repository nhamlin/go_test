package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt") // Open the file. "err" will hold any error message that is generated
	if err != nil {                  // If the error is not nothing then handle the errors\
		fmt.Println(err)
		return
	}
	defer file.Close() // Defer the file close to the end of the function, which is main() in this case.

	stat, err := file.Stat() // Get the file statistics
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Last updated:", stat.ModTime())

	bs := make([]byte, stat.Size()) // Create an array/slice (buffer) of bytes the size of the file
	_, err = file.Read(bs)          // Read the file into the buffer
	if err != nil {
		fmt.Println(err)
		return
	}

	str := string(bs) // Create a string out of the bytes in the buffer
	fmt.Println(str)
}
