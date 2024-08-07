package main

import (
	"fmt"
)

func main() {
	// constant

	const agency string = "Fast Tracks"
	fmt.Println(agency)

	// multiple 

	const (
		founded = 2001
		founder = "james bond"
	)

	fmt.Println(founded, founder)
	
	// iota
	
	const (
		_ = iota
		Economy 
		Compact
		Standard
		Fullsize
		Luxury
	)

	fmt.Println("lv :", Economy)
	fmt.Println("lv :",Compact)
	fmt.Println("lv :",Standard)
	fmt.Println("lv :",Fullsize)
	fmt.Println("lv :",Luxury)



}
