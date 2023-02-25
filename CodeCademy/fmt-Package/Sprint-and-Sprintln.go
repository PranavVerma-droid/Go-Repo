/* While we’ve been using fmt methods to print things out, remember, it’s the formatter package. We have other methods that don’t print strings, but format them instead like fmt.Sprint() and fmt.Sprintln().

grade := "100"
compliment := "Great job!"
teacherSays := fmt.Sprint("You scored a ", grade, " on the test! ", compliment)

fmt.Print(teacherSays)
// Prints: You scored a 100 on the test! Great job!

Take a closer look at teacherSays and how calling fmt.Sprint() doesn’t print out anything. Rather, it returned a value that we store in teacherSays. When a value is returned, it means that a function did some computation and is giving back the computed value. Afterward, we can use the returned value for later usage. We’ll elaborate more on returns in our Go functions lesson. For now, we should acknowledge that returning a value and printing it are two different things. In our code above, we’ve formatted one string by concatenating four separate strings. To see the value of teacherSays, we have to use a print statement.

fmt.Sprintln() works like fmt.Sprint() but it automatically includes spaces between the arguments for us (just like fmt.Println() vs fmt.Print()!):

quote = fmt.Sprintln("Look ma,", "no spaces!")
fmt.Print(quote) // Prints Look ma, no spaces!

Even though we didn’t add a trailing space in "Look ma," or a leading space in "no spaces!", quote is concatenated with a space in between: "Look ma, no spaces!". */

package main

import "fmt"

func main() {
	step1 := "Breathe in..."
	step2 := "Breathe out..."

	// Add your code below:
	meditation := fmt.Sprintln(step1, step2)
	fmt.Println(meditation)

}
