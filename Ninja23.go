package main

import (
	"fmt"
)

var y = 32
var z int = 42

func main() {

	fmt.Printf("%v\t%v\n", y, z)
	fmt.Printf("%T\t%T\n", y, z)
	fmt.Println(y, z)

}
