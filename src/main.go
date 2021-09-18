package main

import (
	"fmt"
	"time"
)

func main(){
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			ch1 <- true
		}
	}()

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			ch2 <- true
		}
	}()

	go func() {
		for {
			<- ch1
			fmt.Println("ch1 수신")
			<- ch2
			fmt.Println("ch2 수신")
		}
	}()

	time.Sleep(5 * time.Second)
}