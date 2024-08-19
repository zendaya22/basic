package main

import (
	"fmt"
	"runtime"
	"sync"
)


func speak(total int, message string, wg *sync.WaitGroup){

	if wg != nil {
		defer wg.Done()
	}

	for i := 0; i < total; i++ {
		fmt.Println("Speak", message, "- (", i+1, ")")
	}
}

func main() {

	wg := sync.WaitGroup{}

	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("running goroutine", runtime.NumCPU(), "cpu")
	wg.Add(3)
	go speak(2, "goroutine 1", &wg)
	go speak(2, "goroutine 2", &wg)
	go speak(2, "goroutine 3", &wg)
	

	wg.Wait()

	fmt.Println("done")


}
