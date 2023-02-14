/*
Now we have completed the basics of using packages, let’s move on to comments.

Before Commenting…
“Don’t comment bad code — rewrite it.” — Brian W. Kernighan and P. J. Plaugher.

Try to make code as clean and self-explanatory before adding comments. They should be the icing on the cake rather than the filling!

Comments are useful for:
Explaining what the code does & why something was done a certain way
Outlining important or fragile blocks of code, which require extra care
Noting down what we need to do when we are writing the code
Disabling code without deleting it
Adding information to be picked up by the go doc tool (more on that later)

There are two types of comments in Go.

Line Comments
Line comments start with a // and the rest of the line is ignored by the compiler.

// This entire line is ignored by the compiler
// fmt.Println("Does NOT print")
fmt.Println("This gets printed!") // This part gets ignored
Note how you can add a // after the code, without affecting it.

Block Comments
Block comments can span multiple lines. They start with a /* and end with a * /, enveloping everything inside:

/*
This is ignored.
This is also ignored. 
fmt.Println("This WON'T print!") */
/*
In the example above we’ve commented out all three lines using a block comment.


Now let’s use some comments in our own code! */

package main

import "fmt"

func main() {
  //Are we racing or coding?
	/*fmt.Println("Ready")
  fmt.Println("Set")*/
  fmt.Println("Gooooo!")
}

/*
Output:

$ go run main.go
# command-line-arguments
./main.go:6:7: syntax error: unexpected we at end of statement
./main.go:6:26: invalid character U+003F '?'
$ go main.go
go main.go: unknown command
Run 'go help' for usage.
$ go help
Go is a tool for managing Go source code.

Usage:

        go <command> [arguments]

The commands are:

        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         add dependencies to current module and install them
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

        buildconstraint build constraints
        buildmode       build modes
        c               calling between Go and C
        cache           build and test caching
        environment     environment variables
        filetype        file types
        go.mod          the go.mod file
        gopath          GOPATH environment variable
        gopath-get      legacy GOPATH go get
        goproxy         module proxy protocol
        importpath      import path syntax
        modules         modules, module versions, and more
        module-get      module-aware go get
        module-auth     module authentication using go.sum
        packages        package lists and patterns
        private         configuration for downloading non-public code
        testflag        testing flags
        testfunc        testing functions
        vcs             controlling version control with GOVCS

Use "go help <topic>" for more information about that topic.

$ go build main.go
$ ./main.go
bash: ./main.go: Permission denied
$ go run main.go
Gooooo!
$ 

*/