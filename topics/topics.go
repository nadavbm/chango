package topics

import (
	"fmt"
	"time"
)

// TimerSize in seconds
const TimerSize = 6

// SendReceiveTopics will send and receive from a topic
func SendReceiveTopics() {
	fmt.Println("\nstart topics")
	topics := make(map[string](chan string))

	topics["farm"] = make(chan string, 3)
	msg := "hello farmer"

	var stop = false
	go func() {
		for {
			now := time.Now()
			time.Sleep(1 * time.Second)
			if stop {
				fmt.Println("timer stopped", now.Format(time.UnixDate))
				break
			}
			fmt.Println("send to channel", msg, now.Format(time.UnixDate))
			topics["farm"] <- msg
		}
	}()

	msg2 := "yeah"
	go func() {
		for {
			now := time.Now()
			time.Sleep(1 * time.Second)
			if stop {
				fmt.Println("timer stopped", now.Format(time.UnixDate))
				break
			}
			fmt.Println("send to channel", msg2, now.Format(time.UnixDate))
			topics["farm"] <- msg2
		}
	}()

	duration := time.Second * TimerSize
	timer := time.NewTimer(duration)

	for {
		select {
		case <-topics["farm"]:
			fmt.Println("message offload from farm channel", <-topics["farm"], time.Now().Format(time.UnixDate))
		case <-timer.C:
			stop = true
			timer.Stop()
			fmt.Println(fmt.Sprintf("Pubsub is closed after %v, channel length %d", duration, len(topics)), time.Now().Format(time.UnixDate))
			time.Sleep(1 * time.Second)
			close(topics["farm"])
			return
		}
	}
}
