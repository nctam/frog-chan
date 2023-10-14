package server

type Config struct {
	BotToken    string       `yaml:"bot-token"`
	Channels    []string     `yaml:"channels"`
	TargetUsers []TargetUser `yaml:"target-users"`
	Toggles     []Toggle     `yaml:"toggles"`
}

type TargetUser struct {
	ID           string   `yaml:"id"`
	Nickname     string   `yaml:"nickname"`
	ChannelIDs   []string `yaml:"channel-ids"`
	Probability  float64  `yaml:"probability"`
	UniversalSet int      `yaml:"universal-set"`
}

type Toggle struct {
	Name    string `yaml:"name"`
	Enabled bool   `yaml:"enabled"`
}
