package command

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"is-tgbot/internal/keys"
	"is-tgbot/internal/utils/sender"
	"is-tgbot/models/keyboard"
)

type Menu struct {
	bot *bot.Bot
}

func NewMenuCommand(bot *bot.Bot) *Menu {
	return &Menu{bot: bot}
}

func (c *Menu) GetCommand() string {
	return keys.Menu
}

func (pc *Menu) Handle(ctx context.Context, b *bot.Bot, update *models.Update) {
	text := "💸Ваш баланс: 00.00 руб.\n🌍Страна: Россия\n🌐Оператор: -"
	sender.SendKeyboard(ctx, keyboard.MainMenu, update, text, b)
}
