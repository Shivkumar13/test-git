package main

import (
	"fmt"
)

var y = 42
var z string = `"Shaken, 

not stirred"`

func main() {

	fmt.Println(y)
	fmt.Printf("%T\n", y)
	// z = 43
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
