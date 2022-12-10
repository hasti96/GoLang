package main

import (
	"fmt"

	"github.com/hsti96/discord-ping/bot"
	"github.com/hsti96/discord-ping/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
