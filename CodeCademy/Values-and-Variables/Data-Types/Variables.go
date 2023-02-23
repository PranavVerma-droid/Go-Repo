/*
Now that we have some background on what types are, we can talk about what variables are and how we make and use them. A variable is a named value (like a constant) with the added feature that it can change during the running of a program. If we’re working with a value in various places in our program, we can store that value in a variable to easily access it later.

Variables are defined with the var keyword and two pieces of information: the name that we will use to refer to them and the type of data stored in the variable. Since variables can be updated we don’t even need to assign a value initially. Here’s a couple of variable definitions:

var lengthOfSong uint16
var isMusicOver bool
var songRating float32

Above, we created three variables:

1. An unsigned 16-bit integer called lengthOfSong.
2. A boolean value called isMusicOver.
3. A 32-bit float called songRating.
Notice that our variable names also follow the same naming convention of constants, using camelCase with a descriptive name.

Create a variable called jellybeanCounter, that will store the number of jellybeans in a jar. Give it a type of int8.
Successfully doing this will result in an error! We’ll talk more about errors in Go and how to read them in the next exercise, so continue when you’re ready.
*/

package main

func main() {
	// Create the variable jellybeanCounter
	// here and make its type int8
	var jellybeanCounter int8

}
