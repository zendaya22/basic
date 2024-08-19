package main

import "fmt"

type Car struct{
	Name string
	Color string
}

func (c Car) getName()string{
	return c.Name
}

func (c Car) SayHello(){
	fmt.Println("Hello dari", c.Name, c.Color)
}

func main() {
	car := Car{
		Name: "audy",
		Color: "blue",
	}

	name := car.getName()
	fmt.Println(name)
	car.SayHello()
	

}
