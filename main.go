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
func add(a int, b int)int{
	return a + b
}
func main() {

	a := 2
	b := 3

	fmt.Println("a = ", a, "b = ",b)
	
	swap(&a, &b)
	
	fmt.Println("a = ", a, "b = ",b)
}
