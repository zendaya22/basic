package main

import "fmt"

type Fruits struct{
	Name string
	Weight int
}

func main() {
	var fruit1 = Fruits{
		Name: "banana",
		Weight: 20,
	}

	fruit2 := Fruits{
		Name: "apple",
		Weight: 30,
	}

	fruit3 := Fruits{"mango", 20}

	fmt.Println(fruit1)
	fmt.Println(fruit2)
	fmt.Println(fruit3)


}
