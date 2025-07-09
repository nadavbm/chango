package main

import (
	"github.com/nadavbm/chango/cmd"
)

func main() {
	cmd.ParseCommandArgs()
	cmd.Execute()
}
