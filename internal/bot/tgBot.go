package bot

import (
	"context"
	"errors"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/wasd0/is/pkg/logger"
	"is-tgbot/models/keyboard"
	"os"
)

const initError = "bot init error"

func Start(ctx context.Context) {
	token := os.Getenv("TOKEN")

	if token == "" {
		logger.Log().Fatal(errors.New(initError), "TOKEN environment variable is empty")
	}

	opts := []bot.Option{
		bot.WithDefaultHandler(defaultHandler),
		bot.WithCallbackQueryDataHandler("button", bot.MatchTypePrefix, callbackHandler),
		bot.WithWorkers(10),
	}

	if b, err := bot.New(token, opts...); err != nil {
		logger.Log().Fatal(err, initError)
	} else {
		b.Start(ctx)
	}
}

func callbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.CallbackQuery == nil || update.CallbackQuery.Message.Message == nil {
		return
	}

	if _, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	}); err != nil {
		logger.Log().Error(err, "Callback error")
	}

	// mock logic
	if _, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Message.Chat.ID,
		Text:   "You selected the button: " + update.CallbackQuery.Data,
	}); err != nil {
		logger.Log().Error(err, "Send message error")
	}
}

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: keyboard.MainMenu,
	}

	if update.Message == nil {
		return
	}

	if _, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "💸Ваш баланс: 00.00 руб.\n🌍Страна: Россия\n🌐Оператор: -",
		ReplyMarkup: kb,
	}); err != nil {
		logger.Log().Error(err, "Send message error")
	}
}
