package main

import (
	"fmt"
)

func main () {
	// for init; condition; post {}, looping in go
	for i :=0; i <= 10; i++ {
		for j := 0;
		 j < 3; j++ {
			 fmt.Printf("The outer loop: %d\t The inner loop: %d\n", i, j)
		 }

	}
	divide () 
	testing ()

}

func divide () {
	x := 83 / 40
	y := 83 % 40 
	fmt.Println(x, y)
}

func testing () {
	for i := 33; i < 122; i++{
		fmt.Printf("%v\t%#Un", i, i)
	}
}