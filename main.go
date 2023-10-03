package main

import (
	"github.com/nadavbm/chango/hello"
	"github.com/nadavbm/chango/topics"
)

func main() {
	hello.SayHello()
	topics.SendReceiveTopics()
}
