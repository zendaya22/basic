package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

type User struct {
	Name   string
	Age    int
	Salary int
}

func main() {
	// step 1
	// read file to get all data, and put it into channel
	dataCh, err := readFile("data.json")
	if err != nil {
		panic(err)
	}

	data1 := cangeCurrencyToIDR(dataCh)
	data2 := cangeCurrencyToIDR(dataCh)
	data3 := cangeCurrencyToIDR(dataCh)


	datas := merge(data1, data2, data3)

	// tampilkan hasil dataCh
	for data := range datas {
		log.Printf("data penghsilan %v", data.Salary)
	}
	defer fmt.Println(<-datas)
}

func readFile(filename string) (<-chan User, error) {
	// membuat channel out untuk menampung data user
	// channel ini akan menjadi syarat untuk lanjut ke step selanjutnya
	outCh := make(chan User)

	// proses membaca file
	// hasil dari ReadFile adalah sebuah []byte
	data, err := os.ReadFile(filename)
	if err != nil {
		return outCh, err
	}

	users := []User{}
	// karena hasilnya adalah []byte dan kita tahu bahwa file itu adalah json
	// jadi bisa langsung di unmarshal
	// err = json.Unmarshal(data, &users)
	err = json.Unmarshal(data, &users)
	if err != nil {
		return outCh, err
	}

	log.Println("length users", len(users))

	go func() {
		// proses memasukkan data user ke channel out
		for _, user := range users {
			outCh <- user
		}

		// selalu close sebuah channel untuk menandakan channel tersebut telah selesai diisi
		// dan setelah ini, kita tidak boleh memasukkan data ke dalam channel tersebut.
		// Hanya boleh untuk membaca datanya
		close(outCh)
	}()

	return outCh, nil
}

func cangeCurrencyToIDR(dataCh <-chan User) <-chan User {
	outCh := make(chan User)

	go func() {

		for data := range dataCh {
			newData := data

			newData.Salary = newData.Salary * 10_000

			outCh <- newData
		}

		close(outCh)

	}()

	return outCh
}

func merge(dataCh ...<-chan User) <-chan User {
	out := make(chan User)

	wg := sync.WaitGroup{}

	for _, data := range dataCh {
		wg.Add(1)

		go func(data <-chan User) {
			log.Println("add", len(data), "into channel")
			for d := range data {
				out <- d
			}
			wg.Done()
		}(data)
	}

	go func() {
		wg.Wait()

		close(out)

	}()

	return out
}
