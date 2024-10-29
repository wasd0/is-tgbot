package command

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/redis/go-redis/v9"
	"is-tgbot/internal/keys"
	"is-tgbot/internal/model/keyboard"
	"is-tgbot/internal/utils"
)

type Settings struct {
	bot *bot.Bot
}

func NewSettingsCommand(bot *bot.Bot) *Settings {
	return &Settings{bot: bot}
}

func (s *Settings) GetCommand() string {
	return keys.Settings
}

func (s *Settings) Handle(ctx context.Context, b *bot.Bot, update *models.Update, _ *redis.Client) {
	text := "Settings"
	utils.SendKeyboard(ctx, keyboard.Settings, update, text, b)
}
