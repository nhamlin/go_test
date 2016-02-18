package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func getHash(filename string) (uint32, error) {
	f, err := os.Open(filename) // Open the file
	if err != nil {
		return 0, err
	}
	defer f.Close() // Remember to close opened files

	h := crc32.NewIEEE()   // Create a new header
	_, err = io.Copy(h, f) // Copy the file into the hasher - copy takes (dst, src) and returns (bytesWritten, error)
	if err != nil {        // We don't care how many bytes were written, but we do want to know if there was an error
		return 0, err
	}
	return h.Sum32(), nil
}

func main() {
	h1, err := getHash("test1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	h2, err := getHash("test2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(h1, h2, h1 == h2)
}
