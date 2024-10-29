package command

import (
	"github.com/go-telegram/bot"
	"is-tgbot/internal/keys"
)

type Provider struct {
	bot      *bot.Bot
	commands map[string]Command
}

func NewCommandProvider(bot *bot.Bot) *Provider {
	commands := map[string]Command{
		keys.Profile:  NewProfileCommand(bot),
		keys.Menu:     NewMenuCommand(bot),
		keys.Settings: NewSettingsCommand(bot),
	}
	return &Provider{bot: bot, commands: commands}
}

func (cp *Provider) Get(command string) Command {
	return cp.commands[command]
}
