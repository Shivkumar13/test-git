package main

import (
	"fmt"
)

const (
	a = 2014
	b = iota + a
	c = iota + a
	d = iota + a
)

func main() {

	fmt.Println(a, b, c, d)

}
