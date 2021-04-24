package main

import (
	"fmt"
)

//testing Git

var y = 42

func main() {

	fmt.Println(y)
	fmt.Printf("%T\n", y)
	//	fmt.Printf("%d\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	y = 911
	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	// s :=
	fmt.Println(s)

}
