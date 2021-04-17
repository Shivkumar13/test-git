package main

import "fmt"

type shiv int

var x shiv
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("\n%T\n", x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	var p = int(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", p)

}
