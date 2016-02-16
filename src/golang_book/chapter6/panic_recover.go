package main

import "fmt"

func main() {
	// This *will* run because it's deferred until the end of the function, which happens when the panic() is called
	defer func() {
		str := recover() // recover() stops the panic and returns the value that was passed to the call to panic
		fmt.Println(str)
	}()
	panic("PANIC!")
	str := recover() // This will never happen because the panic() is like die()
	fmt.Println(str) // This will never happen because the panic() is like die()
}

// panic() is like throwing an error
// defer is like the finally in a try..catch..finally block
