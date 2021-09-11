package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	second := 10

	go sendMessage(ch)

	for second > 0 {
		select {
			case <- ch :
				fmt.Println("메시지가 발송되었습니다.! ㅋㅋ")
				second = 0
			default : 
				fmt.Printf("%d초가 남았습니다.\n", second)

		}
		second--
		time.Sleep(time.Second)
	}
	
}

func sendMessage(ch chan string){
	var msg string
	fmt.Scanf("%s", msg)

	ch <- msg
}