package main

import "fmt"

func main(){
	fmt.Println("hello ")

	const agency string = "fast track"
	var name string = "dimas"

	// another variabel

	user_name := "dimas"

	// integer

	var total_cars int = 50

	// float aouto

	startingPrice := 29.99

	fmt.Println("hello", name)
	fmt.Println("welcome to",agency)
	fmt.Println("user name", user_name)
	fmt.Println("Out fleet consists of", total_cars, "cars")
	fmt.Printf("Our prices start at %v\n", startingPrice)
	fmt.Println("take your pick")

	insuranceIncluded := false
	fmt.Println(insuranceIncluded)
	
}