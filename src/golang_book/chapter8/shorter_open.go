package main

import (
	"fmt"
	"io/ioutil"
)

// Shorter version to read the entire file!

func main() {
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(bs)
	fmt.Println(str)
}
