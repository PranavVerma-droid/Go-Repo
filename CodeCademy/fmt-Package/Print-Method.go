/*
fmt.Println() allows us to print to the terminal and view the data that we’re working with. It has some defaulting styling built-in that makes viewing data easier for us. fmt.Println() prints its arguments (data provided within its parentheses ( )) with an included space in between each argument and adds a line break at the end. Take for example:

fmt.Println("Println", "formats", "really well")
fmt.Println("Right?")

Which prints:

Println formats really well
Right?

Notice that our first print statement has 3 arguments and each one has an accompanying space between each argument even though we never outright included one in our code. For our second print statement, the argument is printed on the next line since Println() adds a line break for us.

However, there are times we might not want the default formatting. In such cases, using fmt.Print() would be more appropriate:

fmt.Print("The answer is", ": ")
fmt.Print("12")

The above code snippet would print:

The answer is: 12

Notice that there’s no default spacing added when fmt.Print() has multiple arguments. Also, since fmt.Print() doesn’t add a line break after printing, the argument for the second print statement print on the same line as the first print statement’s.
Let’s compare this for ourselves. */

package main

import "fmt"

func main() {
  fmt.Println("Let's first see how", "the Println() method works.")
  fmt.Println("Notice that each statement adds a newline for us.")
  fmt.Println("There's also a default space", "between the string arguments.")
  
  // Add your code below:
  fmt.Print("Print", "is", "different")
  fmt.Print("See?")
  
}
