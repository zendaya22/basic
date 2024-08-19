package main

import (
	"fmt"
	"runtime"
	"time"
)


func speak(total int, message string){
	for i := 0; i < total; i++ {
		fmt.Println("Speak", message, "- (", i+1, ")")
	}
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("running goroutine", runtime.NumCPU(), "cpu")
	
	go speak(2, "goroutine 1")
	go speak(2, "goroutine 2")
	go speak(2, "goroutine 3")
	time.Sleep(1 * time.Second)

}
