/* What if…? What if we’re hungry? If it’s raining? If the alarm’s ringing?

We would do something in response to these conditions.

if statements work very similarly to our own decision-making process. Let’s look at Go’s if statement:

alarmRinging := true
if alarmRinging {
  fmt.Println("Turn off the alarm!!")
}
In our example, we have a variable alarmRinging that has a value of true. Then we have an if statement that checks if the condition next to the if keyword is true. Then we have an opening curly brace { with code inside followed by a closing curly brace }. If the condition is true, then the code in between the curly braces { } is executed. In this case, "Turn off the alarm!!" is printed to the console.

In our if statement we could have provided parentheses, like so:

if (alarmRinging) {
  fmt.Println("Turn off the alarm!!")
}
The parentheses are optional and can make it easier to read, but typically, we’ll see the if statement written without parentheses. */

package main

import "fmt"

func main() {
	heistReady := true

	heistReady = false

	if heistReady {
		fmt.Print("Let's Go!")
	}
}
