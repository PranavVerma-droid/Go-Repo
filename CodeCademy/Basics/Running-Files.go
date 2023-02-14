/*
Running Files
Great, we were able to use the Go compiler to build an executable. We can run that executable as many times as we want.

But what happens if we ever wanted to change our program? We would have to compile another executable file and then run that file. Imagine having to do that every single time!

Thankfully, we have the go run command followed by the name of the Go file. We can use the go run command like so:

go run greet.go
The go run command compiles the code (like go build) and executes it. Unlike go build, go run will NOT create an executable file in our current folder. */

package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
