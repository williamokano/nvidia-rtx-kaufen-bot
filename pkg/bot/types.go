package bot

type APIResponse struct {
	Success bool            `json:"success"`
	Map     *map[string]any `json:"map"`
	ListMap []any           `json:"listMap"`
}

type Config struct {
	TelegramChatID string `env:"TELEGRAM_CHAT_ID,required"`
	BotToken       string `env:"TELEGRAM_BOT_TOKEN,required"`
}
