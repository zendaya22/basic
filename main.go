package main

import (
	"log"
	"sync"
)

func calculate(nums ...int) (total int) {
	for _, n := range nums {
		total += n
	}
	return
}

func main() {
	wg := sync.WaitGroup{}
	total := 0
	nums := []int{1, 2, 3, 4, 5}
	wg.Add(1)
	go func(numbers ...int) {
		// nilai total akan di overwrite disini
		total = calculate(numbers...)
		wg.Done()
	}(nums...)

	// prose menunggu sampai goroutine selesai
	wg.Wait()

	// menampilkan nilai total
	log.Println(total)
}
