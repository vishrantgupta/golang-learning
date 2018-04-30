/*

# Learning
A variable can be declared in if condition and its scope is available in that if block
and not outside the block

There is no ternary if in Go, which I am fan of for shorthand, you'll need to use full if statement
even for basic conditions

*/

package main

import "fmt"

func main() {

	if (2 % 2 == 0) {
		fmt.Println("Divisible");
	} else {
		fmt.Println("Not divisible");
	}
	
	if i := 10; i < 11 {
		fmt.Println("Value of i =", i)
	}
	
	// this will give compile time error
	// fmt.Println("Value of i outside if =", i)
	
	// this does not work (invalid character U+003F '?')
	// fmt.Println(10 % 2 == 0 ? "Even" : "Old")

}