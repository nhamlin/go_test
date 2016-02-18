package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	h := crc32.NewIEEE()    // Create a header
	h.Write([]byte("test")) // Write our data to it
	v := h.Sum32()          // Calculate the crc32 checksum
	fmt.Println(v)
}
