package sender

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"is-tgbot/pkg/logger"
)

func SendKeyboard(ctx context.Context, keyboard [][]models.InlineKeyboardButton, update *models.Update, text string, b *bot.Bot) {
	var chatId any

	if update.Message != nil {
		chatId = update.Message.Chat.ID
	} else if update.CallbackQuery != nil {
		chatId = update.CallbackQuery.From.ID
	} else {
		return
	}

	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}

	if _, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatId,
		Text:        text,
		ReplyMarkup: kb,
	}); err != nil {
		logger.Log().Error(err, "Send message error")
	}
}
