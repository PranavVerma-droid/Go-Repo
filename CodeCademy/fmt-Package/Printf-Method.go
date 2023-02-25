/* Using fmt.Println() and fmt.Print() we have the ability to concatenate strings, i.e. combine different strings into a single string:

guess := "C"
fmt.Println("Is", guess, "your final answer?")
// Prints: Is C your final answer?

With fmt.Printf(), we can interpolate strings, or leave placeholders in a string and use values to fill in the placeholders. Let’s revisit the same example using fmt.Printf():

guess := "C"
fmt.Printf("Is %v your final answer?", guess)
// Prints: Is C your final answer?

The first argument we provide fmt.Printf() is the string: "Is %v your final answer?". The %v portion is our placeholder and is known as a verb in Go. Verbs are identified by the combination of a % character followed by a letter. The specific letter informs what goes fills in the placeholder, in this case, %v gets the value of "C" from our second argument, guess.

As long as we provide enough arguments, we can even add multiple placeholders:

selection1 := "soup"
selection2 := "salad"
fmt.Printf("Do I want %v or %v?", selection1, selection2)
// Prints: Do I want soup or salad?

Notice that the placement of the arguments matters! If we switched the position of selection1 and selection2, it would print: Do I want salad or soup?.
We’ll go over more verbs in the next exercise, but let’s first practice using fmt.Printf() with %v. */

package main

import "fmt"

func main() {
	animal1 := "cat"
	animal2 := "dog"

	// Add your code below:
	fmt.Printf("Are you a %v or a %v person?", animal1, animal2)

}
