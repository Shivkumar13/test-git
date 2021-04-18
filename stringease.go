package main

import "fmt"

func main() {

	s := "HELLO PLAYGROUND"

	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Printf("%x\n", bs)
	fmt.Printf("%X\n", bs)
	fmt.Printf("%p\n", bs)

	for i := 0; i < len(s); i++ {

		fmt.Printf("%#U", s[i])
	}

	for i, v := range s {
		fmt.Println(i, v)
	}

}
