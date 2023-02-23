/*Programs, like the ones we write in Go, are excellent at processing and performing operations on data. But in programming, “data” can be so many different things. Data can be numbers, boolean values (true/false), strings (blocks of text), or combinations thereof.

In this lesson, we’ll be covering how to store “data” by creating and using variables in Go.

We’re going to look into the different ways we can interact with these values (like adding two numbers together, or creating a longer message from two strings). We’re also going to discuss data types, the specific divisions that Go uses to “expect” certain qualities from variables. By creating and assigning values to our variables, we’ll understand the limitations and benefits of using different data types.

We’ll also cover how to read Go error messages — creating and using variables adds a new level of complexity to our programs to make these errors much more likely. This will be a good learning process, life as a programmer involves reading and interpreting these error messages quite often!
*/

package main

import "fmt"

func main() {
  var stationName string
  var nextTrainTime int8
  var isUptownTrain bool
  
  stationName = "Union Square"
  nextTrainTime = 12
  isUptownTrain = false
  
  fmt.Println("Current station:", stationName)
  fmt.Println("Next train:", nextTrainTime, "minutes")
  fmt.Println("Is uptown:", isUptownTrain)
  
  stationName = "Grand Central"
  nextTrainTime = 3
  isUptownTrain = true
  
  fmt.Println("")
  fmt.Println("Current station:", stationName)
  fmt.Println("Next train:", nextTrainTime, "minutes")
  fmt.Println("Is uptown:", isUptownTrain)
}