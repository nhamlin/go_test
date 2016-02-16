package main

import "fmt"

func main() {
	// A closure refers to a lambda function whose open bindings (free variables) have been
	// closed by (or bound in) the lexical environment, resulting in a closed expression, or closure. (Wikipedia)

	// Note: Closure is not the same as an anonymous function.
	// Anonymous functions are a function literal without a name.
	// Closure: an instance of a function, a value, whose non-local variables have been bound to values or
	//          storage locations. The non-local variables have to be bound to values.
	// Closures are variables with a closure function as a value

	nextEven := makeEvenGenerator()
	ab := [5]int{1, 2, 3, 4, 5}
	for range ab {
		fmt.Println(nextEven())
	}

}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
