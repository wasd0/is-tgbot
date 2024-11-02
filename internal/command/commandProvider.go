package command

import (
	"is-tgbot/internal/app/serviceProvider"
	"is-tgbot/internal/keys"
)

type Provider struct {
	commands map[string]Command
	services *serviceProvider.ServiceProvider
}

func NewCommandProvider(services *serviceProvider.ServiceProvider) *Provider {
	commands := map[string]Command{
		keys.Profile:  NewProfileCommand(services.CacheService()),
		keys.Menu:     NewMenuCommand(services.CacheService()),
		keys.Settings: NewSettingsCommand(),
	}
	return &Provider{commands: commands, services: services}
}

func (cp *Provider) Get(command string) Command {
	return cp.commands[command]
}
