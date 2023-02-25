/* So far we’ve been checking on boolean values (variable assigned a true or false value). But, we can check more than a single value using comparison operators. Here are two commonly used comparison operators:

Operator	Meaning:
==			Is equal to
!=			Is NOT equal to

To use a comparison operator, we have two values (or operands) with a comparison operator in between the two values. The combination of values and the operator is evaluated by Go’s compiler which looks at the left value, compares it to the right value, and decides on a true or false value. Let’s take a look at == first:

"password1" == "password1" // Evaluates to true
"password1" == "badpass"   // Evaluates to false
When we use comparison operators, we check the left value against the right value. It can be helpful to think of comparison statements as questions. When the answer is “yes”, the statement evaluates to true, and when the answer is “no”, the statement evaluates to false. The above code’s first example would be asking: is "password1" the same as "password1"? Yes it is, so "password1" == "password1" evaluates to true. We can read the second example as, is "badpass" the same as "password1"? No, "badpass" is not the same as "password1" so it evaluates to false.

Let’s take a look at the != operator in action:

123 != 12 // Evaluates to true
123 != 123 // Evaluates to false
Above, the first line checks if 123 and 12 are not equivalent and since the integers are different values, it evaluates to true. This time we can think of the question as: is 123 NOT the same as 12? Yes, they’re not the same, so it evaluates to true. The same idea persists for the second example: is 123 not the same as 123? No, they are the same, so it evaluates to false.

It can be very helpful to take a second to think through how the operands are being compared before committing it to code. Let’s practice that now! */

package main

import "fmt"

func main() {
	lockCombo := "PranavVerma234"
	//LOL THIS IS NOT MY ACTUAL PASSWORD XD GO AHEAD AND TRY IT L
	//robAttempt := "1-1-1"
	var attempt string
	fmt.Print("Please Enter Your Password: ")
	fmt.Scan(&attempt)

	// Add your code below:
	if lockCombo == attempt {
		fmt.Println("The vault is now opened.")
	} else {
		fmt.Println("Wrong Pass! Closing the Vault!")
	}
}
