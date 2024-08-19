package slices

import "fmt"

func Slice(){
	fruits := []string{"apple", "mango", "orange", "banana"}
	animal := [3]string{"cat", "dog", "elephant"}

	fmt.Println(fruits[2])
	fmt.Println(animal[2])
}