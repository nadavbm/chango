package hello

import "fmt"

// SayHello will send hello to channel and print the value received from the channel
func SayHello() {
	ch := make(chan string)
	msg := "hello"
	go func() {
		ch <- msg
	}()
	fmt.Println("value from channel", <-ch)
}
