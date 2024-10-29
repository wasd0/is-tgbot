package command

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/redis/go-redis/v9"
)

type Command interface {
	GetCommand() string
	Handle(ctx context.Context, b *bot.Bot, update *models.Update, cache *redis.Client)
}

func getChatId(update *models.Update) *int64 {
	var chatId *int64

	if update.Message != nil {
		chatId = &update.Message.Chat.ID
	} else if update.CallbackQuery != nil {
		chatId = &update.CallbackQuery.From.ID
	}

	return chatId
}
