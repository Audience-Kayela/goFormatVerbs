package main

import "fmt"

func main() {
	// %f Floating point number
	fmt.Printf("A float: %f\n", 3.1415)
	// %d Decimal integer
	fmt.Printf("An integer: %d\n", 15)
	// %s String
	fmt.Printf("A string: %s\n", "Hello")
	// %t Booliean
	fmt.Printf("A boolean: %t\n", true)
	// %v Any value (Chooses appropriate format based on the value's type)
	fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)
	fmt.Printf("Values: %v %v %v\n", "", "\t", "\n")
	// %#v Any value, formatted as it would appear in a program useful for debugging.
	fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", false)
	// %T Type of supplied value(int, string, etc)
	fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true)
	// %% A literal percent sign
	fmt.Printf("Percent sign: %%\n")

}
