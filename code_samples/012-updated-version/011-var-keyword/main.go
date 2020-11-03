package main

import (
	"fmt"
)

type chicken int
var x chicken
var y int 

func main() {
	
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 3
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)

}




