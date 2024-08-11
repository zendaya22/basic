package main

import "fmt"

// function as parameter

func odd(a int) bool {
	return a%2 != 0
}

func try(a int, checkOdd func(a int) bool) string {
	check := checkOdd(a)
	if check {
		return "odd"
	}
	return "even"
}

func main() {
	result := try(10, odd)
	fmt.Println(result)
}
