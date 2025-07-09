package messaging

import (
	"fmt"
	"time"

	"github.com/nadavbm/chango/decorator"
	"github.com/nadavbm/chango/singleton"
	"github.com/nadavbm/chango/strategy"
)

const chanSize = 3
const topicName = "messenger"

// SendReceive will send and receive from a topic
func SendReceive(logger decorator.Logger, cfg *singleton.Config) {
	topics := make(map[string](chan string))
	topics[topicName] = make(chan string, chanSize)
	generate := strategy.Str{}
	var stop = false
	go func() {
		for {
			time.Sleep(1 * time.Second)
			if stop {
				logger.Info("timer stopped")
				break
			}
			msg := generate.String(5) + " " + generate.String(3) + " " + generate.String(8)
			logger.Info("send to channel")
			topics[topicName] <- msg
		}
	}()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			msg2 := generate.String(2) + " " + generate.String(4) + " " + generate.String(4)
			if stop {
				logger.Info("timer stopped")
				break
			}
			logger.Info("message from farm channel: " + msg2)
			topics[topicName] <- msg2
		}
	}()

	timer := time.NewTimer(cfg.Duration)
	for {
		select {
		case <-topics[topicName]:
			logger.Info("message from farm channel: " + <-topics[topicName])
		case <-timer.C:
			stop = true
			timer.Stop()
			logger.Info(fmt.Sprintf("Pubsub is closed after %v. Sleep one second and close", cfg.Duration))
			time.Sleep(1 * time.Second)
			close(topics[topicName])
			return
		}
	}
}
