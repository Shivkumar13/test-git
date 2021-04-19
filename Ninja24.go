package main

import (
	"fmt"
)

func main() {

	var a int = 23
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)

	s := a << 1

	fmt.Printf("%d\t%b\t%#x\n", s, s, s)
}
