package command

import (
	"context"
	"errors"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"is-tgbot/pkg/logger"
)

type Command interface {
	GetCommand() string
	Handle(ctx context.Context, b *bot.Bot, update *models.Update)
}

func getChatId(update *models.Update) *int64 {
	var chatId *int64

	if update.Message != nil {
		chatId = &update.Message.Chat.ID
	} else if update.CallbackQuery != nil {
		chatId = &update.CallbackQuery.From.ID
	} else {
		logger.Log().Error(errors.New("cannot get chat id"), update.Message.Chat.Username)
		return nil
	}

	return chatId
}
