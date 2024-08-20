package main

import (
	"io"
	"log"
	"os"
	"time"
)

func main(){
	f, _ := os.Open("text.txt")

	chunkSize := 15

	text := ""

	for {
		chunk := make([]byte, chunkSize)

		bytesRead, err := f.Read(chunk)
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		if bytesRead == 0 {
			break
		}

		log.Println(string(chunk))
		
		text += string(chunk)
		time.Sleep(2 * time.Second)
	}
	
	f.Close()
	log.Println(text)
	
}