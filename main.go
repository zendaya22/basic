package main

import (
	"fmt"
)

func main() {

	// math
	c := 20
	k := 0.0
	// cast data type 
	k = float64(c) + 273.15
	fmt.Println(k)
	var tempInt int = 10
	var tempFloat float64 = float64(tempInt)
	fmt.Println(tempFloat)
	fmt.Printf("integer to float %.1f\n", 10.0)

}
