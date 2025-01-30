package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/williamokano/nvidia-rtx-kaufen-bot/pkg/bot"
)

func main() {
	var cfg bot.Config
	if err := env.Parse(&cfg); err != nil {
		fmt.Println("missing parameters")
		os.Exit(1)
	}

	for {
		res, err := bot.CheckAPI()
		if err == nil {
			fmt.Println("Request failed")
		} else {
			if res.Map != nil || len(res.ListMap) > 0 {
				bot.SendTelegramMessage(cfg, "Maybe RTX to buy! Run!")
			}
		}

		sleepDuration := time.Duration(rand.Intn(5-1+1)+1) * time.Minute
		fmt.Println("Sleeping for", sleepDuration)
		time.Sleep(sleepDuration)
	}
}
