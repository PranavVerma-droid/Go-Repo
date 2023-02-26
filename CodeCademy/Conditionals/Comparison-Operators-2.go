/* There are more comparison operators that we haven’t covered and they may seem familiar from math class:

Operator	Meaning:
<			Less than
>			Greater than
<=			Less than or equal to
>=			Greater than or equal to

Like the previous exercise, we’re still using the comparison operators to compare the left hand value against the right hand value. For instance:

100 < 200 // Evaluates to true
We can read this example as, “Is 100 less than 200? Yes, that is true.

Let’s look at one more example with <=:

100.5 <= 100.5 // Evaluates to true
Just like the less than operator (<), we want to check if the left hand value is less than OR equal to the right hand value. Since 100.5 is less than or equal to 100.5, it evaluates as true. */

package main

import "fmt"

func main() {
	vaultAmt := 2356468

	// Add your code below:
	if vaultAmt >= 200000 {
		fmt.Println("We're going to need more bags.")
		//YEAAAA MORE BAGS LESGOOOOO
	}

}
