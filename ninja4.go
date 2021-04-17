package main

import (
	"fmt"
)

var a int

type shiv int

var x shiv

func main() {

	fmt.Println(x)
	fmt.Printf("\n%T\n", x)
	x = 42
	fmt.Println(x)

}
