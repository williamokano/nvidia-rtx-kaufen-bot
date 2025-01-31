package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/williamokano/nvidia-rtx-kaufen-bot/pkg/bot"
)

func main() {
	botToken, exist := os.LookupEnv("TELEGRAM_BOT_TOKEN")
	if !exist {
		fmt.Println("TELEGRAM_BOT_TOKEN environment variable not set")
		os.Exit(1)
	}

	subscribersList := make(map[int64]bool)
	if subsParam, exist := os.LookupEnv("TELEGRAM_SUBSCRIBERS"); exist {
		subs := strings.Split(subsParam, ",")
		for _, sub := range subs {
			convertedChannelID, err := strconv.Atoi(sub)
			if err != nil {
				fmt.Println("Error converting channel id to int")
			}

			subscribersList[int64(convertedChannelID)] = true
		}
	}

	var telegramBot = bot.NewBot(botToken, subscribersList)

	for {
		res, err := bot.CheckAPI()
		if err != nil {
			fmt.Println("Request failed")
			fmt.Println(err)
		} else {
			fmt.Println(res)
			if res.Map != nil || len(res.ListMap) > 0 {
				telegramBot.Broadcast("Maybe RTX to buy! Run!")
			}
		}

		// Sleep at least 1 minute, up to 6 minutes
		sleepDuration := time.Duration(rand.Intn(300)+60) * time.Second
		fmt.Println("Sleeping for", sleepDuration)
		time.Sleep(sleepDuration)
	}
}
