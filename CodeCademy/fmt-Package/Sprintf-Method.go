/* If we need to interpolate a string, without printing it, then we can use fmt.Sprintf().

Just like fmt.Printf(), fmt.Sprintf() can also use verbs:

correctAns := "A"
answer := fmt.Sprintf("And the correct answer is… %v!", correctAns)

fmt.Print(answer) // Prints: And the correct answer is… A!
fmt.Sprintf() works very similarly to fmt.Printf(), the major difference is that fmt.Sprintf() returns its value instead of printing it out! */

package main

import "fmt"

func main() {
	template := "I wish I had a %v."
	pet := "puppy"

	var wish string

	// Add your code below:

	wish = fmt.Sprintf(template, pet)
	fmt.Println(wish)
}
