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

	//Range form of the for loop interates over a slice or map.

	//when ranging over a slice, two values are returned for each iteratio. The first is the index, and
	//the second is a copy of the element at that index.

	for i, v := range bs {
		fmt.Println(i, v)
	}

}
