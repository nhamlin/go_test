package main

import "fmt"

func main() {
	x := 5
	zeroValue(x)
	fmt.Println(x) // x is still 0 because x is passed to zeroValue as a value (copy) and not the object itself (reference)

	zeroReference(&x)
	fmt.Println(x) // x is now 0

	//what if I used the & and not the pointer
	//zeroValue(&x)
	//fmt.Println(x) //an error is thrown because it's expecting a pointer to be passed to zeroValue
	// this is because the & operator finds the memory address of the variable, not the value.

	// Another way to get a pointer is to use the new function:
	y := new(int) // this will assign the default value of the int, which is 0
	fmt.Println("The memory location of y is:", y)

	// In the previous line, the value of y was expected, however, because y was invoked using the new function it is a pointer/refenence

	fmt.Println("And the default value of int is:", *y)
	one(y)
	fmt.Println("The new value of y is:", *y)

}

func zeroValue(x int) {
	x = 0
}

func zeroReference(x *int) {
	*x = 0 //By using the *, we are saying 'store the int 0 in the memory location x refers to'
}

func one(x *int) {
	*x = 1
}
