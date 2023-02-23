/*
There is no shame in having your code fail to run. Programming errors and exceptions happen all the time and learning to read and understand them is an indispensable tool in a programmer’s toolbox.

When the Go compiler raises an error the program’s binary cannot get created and without the binary, our computer cannot execute the program’s code. Go tries to tell you what the issue is by offering you the following pieces of information: the name of the file where something went wrong, the line number in that file where Go noticed an issue, and the column number (the number of characters from the left) on that line where the error occurred. One common error occurs when Go language’s compiler recognizes that there is a defined variable that isn’t used in the program. Take for example:

	package main

	func main() {
  		var numberWheels int8
	}

When we attempt to run main.go and on line 4 of our file we define the variable numberWheels that we don’t use anywhere else in the program we will see the following message:

	./main.go:4:7: numberWheels declared and not used
Notice that the error message contains the file name (main.go), the line that causes the error to be raised (line 4) and specifically the location (7 characters to the right of the line break). If we remove the variable from our program or use it somewhere we appease Go’s compiler and can run our program.

This may seem like a downside to a language, something that keeps you from expressing your freedom and personality as a programmer, but it is designed to discover possible errors in your program sooner rather than later. An unused variable is a waste of memory, but it can also be a typo or an unintended omission. In this case, Go’s somewhat hardline stance is to keep programmers from wondering why their code isn’t working the way they expect by refusing to run until some action is taken on this unused definition.
*/

package main

import "fmt"

func main() {
	var jellybeanCounter int8

	// Go will raise errors both for
	// jellybeanCounter being unused
	// and for "fmt" being imported
	// and unused.

	// Uncomment the following line
	// and watch the program run
	// successfully!

	fmt.Println(jellybeanCounter)
}
