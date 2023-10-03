package hello

import "fmt"

func SayHello() {
	ch := make(chan string)
	msg := "hello"
	go func() {
		ch <- msg
	}()
	fmt.Println("value from channel", <-ch)
}
