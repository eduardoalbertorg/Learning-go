package main

import "fmt"

/*
1. at the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom TYPE “x
	type hotdog int
	var x hotdog
	var y int

2. in func main
	a. this should already be done
		i. print out the value of the variable “x”
		ii. print out the type of the variable “x”
		iii. assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
		iv. print out the value of the variable “x”
	b. now do this
		i. now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
		ii. then use the “=” operator to ASSIGN that value to “y”
		iii. print out the value stored in “y”
		iv. print out the type of “y”

*/

type hotdog int

var x hotdog
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
