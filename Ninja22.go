package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 19
	c := 20

	f := a == b
	g := b <= c
	h := b >= a
	i := a != b
	j := a < b
	k := c > b

	fmt.Printf("%T\n", f)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

}
