package main

import (
	"fmt"
)

func main() {

	//for init ; condition ; post {}

	for i := 0; i <= 10; i++ {

		for j := 0; j < 3; j++ {

			for k := 0; k < 4; k++ {

				for l := 0; l < 2; l++ {

					fmt.Printf("The Outer loop: %d\t The Inner loop: %d\t The Most-Inner loop: %d\t The l inner loop: %d\n", i, j, k, l)

				}

			}
		}
	}

}
