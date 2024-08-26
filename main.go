package main

import (
	"fmt"
	"os"
)

func main(){
	defer fmt.Println("success generate text")
	text := "halo iam dimas a pasionate backend engginer from indonesia\n you can call me dimas"

	textBytes := []byte(text)

	err := os.WriteFile("text.txt", textBytes, 0666)

	if err != nil {
		panic(err)
	}
}