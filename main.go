package main

import (
	"fmt"
	"go_praco/service"
)

func one(xPtr *int) {

	*xPtr = 1
}
func main() {

	xPtr := new(int)

	one(xPtr)

	fmt.Println(xPtr) // x is 1
	service.UtcToLocal()
	fmt.Printf(" ")
	service.DateTime()
	fmt.Printf(" ")
	service.AutTime()
	fmt.Printf(" ")

}

// 1. How do you get the memory address of a variable?
var num = 1

// 2. How do you assign a value to a pointer?
// 3. How do you create a new pointer?
// 4. What is the value of x after running this program:
