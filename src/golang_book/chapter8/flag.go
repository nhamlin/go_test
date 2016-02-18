package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	maxp := flag.Int("max", 6, "the max value") // Define the flag. In this case, the flag is "-max". 6 is the default value
	flag.Parse()                                // Parse the flag
	fmt.Println(rand.Intn(*maxp))               // Generate a number between 0 and max
}
