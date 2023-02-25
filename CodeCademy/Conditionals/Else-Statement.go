/*
If we’re hungry we would go to eat something. But if we’re not hungry then we don’t. The idea is that we have a backup plan or something we can default to in case our condition isn’t met.

We can provide a default option to our conditional (if statement) by adding an else statement:

isHungry := false
if isHungry {
  fmt.Println("Eat the cookie")
} else {
  fmt.Println("Step away from the cookie...")
}
In the example above, isHungry is a variable with a value of false. We’ve set up an if statement like we saw in the previous exercise. Immediately after the if statement’s closing brace is the else keyword and its opening brace, all on the same line.
The else statement also contains a block of code wrapped by a set of curly braces. The code inside the block will execute by default if the if condition is false. Notice, the else statement does not accept a condition. */

package main

import "fmt"

func main() {
	heistReady := false

	if heistReady {
		fmt.Println("Let's go!")
	} else {
		fmt.Println("Act normal.")
	}
}
