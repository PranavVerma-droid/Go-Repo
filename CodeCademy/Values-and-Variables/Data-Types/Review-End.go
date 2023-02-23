/* In this lesson you learned:

1. How to use literals in a Go program.
2. How to create constants that give names to values.
3. The basic types in Go: ints, floats, complexs, and strings.
4. The different numeric types and what values they accept.
5. How to create variables in Go with a specific type.
6. How to read an error in Go.
7. How to assign values to variables using =.
8. How to create variables in Go with inferred type.
9. How to declare multiple variables on a single line.
10. The “zero” values for variables that haven’t been assigned a value yet.
11. How Go decides the type for “default” int types depending on the architecture of the computer it’s running on.
12. How to update values in variables using = and other related operators.

What a lot of concepts! These tools will be very valuable to you along your journey as a programmer. Congratulations! */

package main

import "fmt"

func main() {
	var congrats string
	congrats = "Congratulations"
	congrats += "!!!"
	fmt.Println(congrats)

	var challenge string = "What else can you do?"
	fmt.Println(challenge)

	reminder := "Pratice is important!"
	fmt.Println(reminder)
}
