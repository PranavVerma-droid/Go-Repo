package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s <filename>\n", os.Args[0])
        os.Exit(1)
    }
    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: could not open file \"%s\": %s\n", os.Args[1], err)
        os.Exit(1)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "Error reading file: %s\n", err)
        os.Exit(1)
    }
}
