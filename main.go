package main

import (
	"fmt"
	"strings"
)

func main(){
	
	str := "Hello, World!"
	length := len(str)
	fmt.Println(length)

	str1 := "dimas"
	str2 := "DIMAs"
	// compare
	fmt.Println(strings.EqualFold(str1, str2))
	// index in string
	fmt.Println(strings.Index(str, "H"))
	// substring
	subs := str[strings.Index(str, "H"): 5]
	fmt.Println(subs)
	a := str[:13]
	b := str[:]
	fmt.Println(a)
	fmt.Println(b)
	// upper and lowwer
	str6 := "Go for web development"
	fmt.Println(strings.ToLower(str6))
	fmt.Println(strings.ToUpper(str6))
	// contain to check 
}