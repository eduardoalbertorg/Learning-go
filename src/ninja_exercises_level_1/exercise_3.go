package main

import "fmt"

/*
1. At the package level scope, assign the following values to the three variables
for x assign 42
for y assign “James Bond”
for z assign true

2. in func main
	a.use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
	b.print out the value stored by variable “s”

*/

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	// The %v prints the value and \t places a tab for easier reading.
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
