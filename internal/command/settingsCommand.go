package command

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"is-tgbot/internal/keys"
	"is-tgbot/internal/model/keyboard"
	"is-tgbot/internal/utils"
)

type Settings struct {
}

func NewSettingsCommand() *Settings {
	return &Settings{}
}

func (s *Settings) GetCommand() string {
	return keys.Settings
}

func (s *Settings) Handle(ctx context.Context, b *bot.Bot, update *models.Update) {
	text := "Settings"
	utils.SendKeyboard(ctx, keyboard.Settings, update, text, b)
}
