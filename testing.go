package main

import "fmt"

func main() {

	fmt.Println("Hello World from Hyderabad, Hi-Tech city.")

	foo()

	fmt.Println("Back to Main and checking Loop inside main")

	i := 1
	for i <= 100 {
		fmt.Println(i)
		i = i + 1
	}

	bar()

}

//Control Flow can be 1] sequential
//2] Loop/Iterative
//3] Contditional

func foo() {

	fmt.Println("I am executing inside Foo now")
}

func bar() {

	fmt.Println("in function bar")
}
