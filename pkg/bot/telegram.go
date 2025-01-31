package bot

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Chat struct {
	ID int64 `json:"id"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Bot struct {
	token       string
	subscribers map[int64]bool
}

func NewBot(token string, subs map[int64]bool) *Bot {
	return &Bot{
		token:       token,
		subscribers: subs,
	}
}

func (bot *Bot) Broadcast(text string) {
	for id := range bot.subscribers {
		err := bot.SendTelegramMessage(id, text)
		if err != nil {
			log.Printf("Error sending message to channel %d: %v", id, err)
		}
	}
}

func (bot *Bot) SendTelegramMessage(chatId int64, message string) error {
	telegramURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", bot.token)
	jsonBody := fmt.Sprintf(`{"chat_id":"%d","text":"%s"}`, chatId, message)

	req, err := http.NewRequest("POST", telegramURL, io.NopCloser(strings.NewReader(jsonBody)))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending message: %w", err)
	}
	defer resp.Body.Close()

	fmt.Println("Message sent, response code:", resp.StatusCode)

	return nil
}
