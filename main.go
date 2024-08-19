package main

import "fmt"

func main(){
    // Diberikan sebuah array seperti berikut :
    animals := [...]string{"Cat", "Dog", "Pinguin", "Chicken", "Snake"}

    // Lalu, lengkapi variable variable berikut sesuai dengan expected-nya :
    mammals := animals[:2]     // expected : {Cat, Dog}
    notMammals := animals[2:]  // expected : {Pinguin, Chicken, Snake}
    haveLegs := animals[:4]    // expeccted : {Cat, Dog, Pinguin, Chicken}

    // Setelah itu, lakukan hal berikut :
    // 1. Ubah value Dog menjadi Cow
    // 2. Ubah value Pinguin menjadi Bird

    // dimulai dari sini
	animals[1] = "Cow"
	animals[2] = "Bird"
	


    // berakhir disini

    // Saat di print, harusnya hasilnya sesuai dengan expected
    fmt.Println(mammals)    // expected : {Cat, Cow}
    fmt.Println(notMammals) // expected : {Bird, Chicken, Snake}
    fmt.Println(haveLegs)   // expected : {Cat, Cow, Bird, Chicken}
    
}