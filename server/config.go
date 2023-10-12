package server

type Config struct {
	BotToken        string           `json:"bot_token"`
	DiscordChannels []DiscordChannel `json:"discord_channels"`
}

type DiscordChannel struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}
