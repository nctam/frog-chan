package server

type Config struct {
    BotToken       string   `yaml:"bot-token"`
    Channels       []string `yaml:"channels"`
    ExcludedUsers  []string `yaml:"excluded-users"`
    Probability    float64  `yaml:"probability"`
    UniversalSet   int      `yaml:"universal-set"`
    Env            string   `yaml:"env"`
    MessagePattern string   `yaml:"message-pattern"`
    PriorUsers     []string `yaml:"prior-users"`
}
