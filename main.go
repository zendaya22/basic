package main

import (
	"fmt"
)
// pointer

func swap(a *int, b *int){
	var temp int = *a
	*a = *b
	*b = temp
}

func main() {

	a := 2
	b := 3
	c := 4
	fmt.Println("a = ", a, "b = ",b)
	fmt.Println(c)
	swap(&a, &b)
	
	fmt.Println("a = ", a, "b = ",b)
}
