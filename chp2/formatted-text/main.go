package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Print integers using %%d: %d|\n", 10)
	fmt.Printf("You can set the width of the printed number, left aligned: %5d|\n", 10)
	fmt.Printf("You can make numbers right-aligned with a given width: %-5d|\n", 10)
	fmt.Printf("The width can be filled with 0s: %05d|\n", 10)

	fmt.Printf("You can use multiple arguments: %d %s %v\n", 10, "yes", true)
	fmt.Printf("You can refer to the same argument multiple times : %d %s %[2]s  %v\n", 10, "yes", true)
	fmt.Printf("But if you use an index n, the next argument will be selected from n+1 : %d %s %[2]s %[1]v  %v\n", 10, "yes", true)
	fmt.Printf("Use %%v to use the default format for the type: %v %v %v\n", 10, "yes", true)
	fmt.Printf("For floating point, you can specify precision: %5.2f\n", 12.345657)
	fmt.Printf("For floating point, you can specify precision: %5.2f\n", 12.0)
	type S struct {
		IntValue    int
		StringValue string
	}

	s := S{
		IntValue:    1,
		StringValue: `foo "bar"`,
	}

	// Print the field values of a structure, in the order they are declared
	fmt.Printf("%v\n", s)
	// {1 foo "bar"}

	// Print the field names and values of a structure
	fmt.Printf("%+v\n", s)
	// {IntValue:1 StringValue:foo "bar"}
}
