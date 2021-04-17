package main

import (
	"fmt"
)

var a int

type hotdog int

var b hotdog = 43

func main() {

	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	//s := fmt.Print("%t", b)
	//	fmt.Println(s)
	//fmt.Println(b)

}
