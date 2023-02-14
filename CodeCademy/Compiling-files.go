/*
Compiling Files
Now we know what Go is for, let’s learn how to use the Go Compiler to compile a file into an executable.

In our terminal, we type in go build followed by the name of our file and press Enter. If we wanted to run a file called greet.go, the command will look like:

go build greet.go
While nothing obvious is shown after we run our command, Go has created our program’s executable file. If we type in the command ls, we’ll see our original Go program and its executable file.

ls
greet     greet.go
To execute the file, we call:

./greet
*/

package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

