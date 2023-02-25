/* https://pkg.go.dev/fmt#hdr-Printing

In addition to %v, Go has a variety of useful verbs (check their documentation for a comprehensive list). Let’s go over a few in this exercise, starting with %T:

specialNum := 42
fmt.Printf("This value's type is %T.", specialNum)
// Prints: This value's type is int.

quote := "To do or not to do"
fmt.Printf("This value's type is %T.", quote)
// Prints: This value's type is string.

As seen by the example above, the verb %T prints out the type of the second argument.

Now look at %d:

votingAge := 18
fmt.Printf("You must be %d years old to vote.", votingAge)
// Prints: You must be 18 years old to vote.

Using %d we can interpolate a number into a string! If we need to include a float:

gpa := 3.8
fmt.Printf("You're averaging: %f.", gpa)
// Prints: You're averaging 3.800000.

With %f, we can limit how precise we are by including a value between the % and f like: %.2f. If we include this in our code:

gpa := 3.8
fmt.Printf("You're averaging: %.2f.", gpa)
// Prints: You're averaging 3.80.

Let’s explore these new verbs in our own code. */

package main

import "fmt"

func main() {
	floatExample := 1.75
	// Edit the following Printf for the FIRST step
	fmt.Printf("Working with a %T", floatExample)

	fmt.Println("\n***") // Added for spacing

	yearsOfExp := 3
	reqYearsExp := 15
	// Edit the following Printf for the SECOND step
	fmt.Printf("I have %d years of Go experience and this job is asking for %d years.", yearsOfExp, reqYearsExp)

	fmt.Println("\n***") // Added for spacing

	stockPrice := 3.50
	// Edit the following Printf for the THIRD step
	fmt.Printf("Each share of Gopher feed is $%.2f!", stockPrice)
}
