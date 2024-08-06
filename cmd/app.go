package cmd

import "fmt"

type Initializer struct{}

func One(xPtr *int) {
	*xPtr = 1
}

func Init() {
	xPtr := new(int)
	One(xPtr)
	fmt.Println(*xPtr) // x is 1
}