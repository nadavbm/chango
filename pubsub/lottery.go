package pubsub

import (
	"fmt"

	"github.com/nadavbm/chango/strategy"
)

const (
	selectedBalls = 6
	ballsQuantity = 49
)

func Lottery() {
	balls := make(chan int)
	go publisher(balls)
	subscriber(balls)
}

func publisher(ch chan<- int) {
	numGenerator := strategy.Int{}
	for i := 0; i < selectedBalls; i++ {
		ch <- numGenerator.Integer(ballsQuantity, 1)
	}
	close(ch)
}

func subscriber(ch <-chan int) {
	fmt.Println("The selected balls are: \n***********************")
	for ball := range ch {
		fmt.Println(ball)
	}
}
