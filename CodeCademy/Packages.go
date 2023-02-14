/*
Now that we understand how to compile and run Go programs, let’s take a look at Go packages.

Projects can contain many .go files, organized into packages. Each package is like a directory: .go files to do with one part of your program can go in one package, other .go files to do with something else can go into another package. If we were writing a calculator program, we could put the files for the calculation in package calc and the files for input & output in package io.

Our Program So Far

package main 
 
import "fmt" 
 
func main () {
  fmt.Println("Hello World") 
}

Package Declaration
Let’s focus on the first line package main. This line is called a package declaration and it tells the Go compiler to which package this ‘.go’ file belongs. If this package declaration is ‘package main’, then this program will be compiled into an executable.

Whitespace
Next is a blank line. Go generally ignores these blank lines, they’re considered whitespace (new lines, spaces, and tabs). While our program doesn’t need the line break, it makes our code easier to read.

Import Statement
Then we have an import statement, import "fmt". The import keyword allows us to use code from other packages, in this case the Println function from the fmt package. Note how the package name is enclosed with double quotes ".

Now that we know how we declare & use Go packages, let’s discuss functions.


1.
In main.go, create the package declaration.
Click “Run” to check our work.

2.
Under the package declaration, import the "fmt" package.
We’ll continue building this program in the next exercise!
Note: Running the program (using the command go run main.go) will result in an error, instead, click the Run button after importing the "fmt" package.
*/

package main

import "fmt"