package main

import "fmt"

//Ignore all of the things written below (my brain go zoom)
// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fib()
	fmt.Println(f(), f(), f(), f(), f())
}

//EZ CLAP LESGOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO
//FINNALLY DONE XD