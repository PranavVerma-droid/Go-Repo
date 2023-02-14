// Link for Packages:
// https://pkg.go.dev/std

/*
Previously, we imported a single package, fmt. This package is useful but it is only one in a list of many included with Go.

The standard packages are so useful that you will often use multiple packages in each .go file.

Importing Multiple Packages
To import multiple packages, we can add multiple import statements:

import "package1"
import "package2"

Or use a single import with parentheses:

import (
  "package1"
  "package2"
)

Package Aliases
We can also provide an alias to a package by specifying an alias name before the package name.

import (
  p1 "package1"
  "package2"
)

In the example above we’ve aliased package1 as p1 and now we can call functions from package1 by using p1 like:

p1.SampleFunc()

Instead of:

package1.SampleFunc() 

1.
In main.go, we have our “Hello World” program set up.

Let’s import another package, "time".

2.
Assign an alias of t to the time package.

3.
Inside of main, let’s use the time package. Add the line:

fmt.Println(t.Now())
This line will print the current time using both the fmt and time packages.

Note: the time printed will be UTC time zone.

4.
Run the program using the command line. Look at the time printed out!
*/

package main

import "fmt"
import t "time"

func main() {
	fmt.Println("Hello World")
  fmt.Println(t.Now())
}

/*

Output:

$ go run main.go
Hello World
2023-02-14 17:09:54.524489098 +0000 UTC m=+0.000037241
$ 

*/


