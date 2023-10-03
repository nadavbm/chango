package topics

import (
	"fmt"
	"time"
)

// TimerSize in seconds
const TimerSize = 4

// SendReceiveTopics will send and receive from a topic
func SendReceiveTopics() {
	topics := make(map[string](chan string))

	topics["farm"] = make(chan string, 2)
	msg := "hello farmer"

	go func() {
		for {
			fmt.Println("send to channel")
			topics["farm"] <- msg
		}
	}()

	duration := time.Second * TimerSize
	timer := time.NewTimer(duration)

	for {
		select {
		case <-topics["farm"]:
			fmt.Println(<-topics["farm"])
		case <-timer.C:
			timer.Stop()
			fmt.Println(fmt.Sprintf("Pubsub is closed after %v", duration))
			time.Sleep(1 * time.Second)
			close(topics["farm"])
			return
		}
	}
}
