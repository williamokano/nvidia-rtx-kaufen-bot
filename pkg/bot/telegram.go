package bot

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func SendTelegramMessage(cfg Config, message string) {
	telegramURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", cfg.BotToken)
	jsonBody := fmt.Sprintf(`{"chat_id":"%s","text":"%s"}`, cfg.TelegramChatID, message)

	req, err := http.NewRequest("POST", telegramURL, ioutil.NopCloser(strings.NewReader(jsonBody)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Message sent, response code:", resp.StatusCode)
}
