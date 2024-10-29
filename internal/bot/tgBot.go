package bot

import (
	"context"
	"errors"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/redis/go-redis/v9"
	"is-tgbot/internal/command"
	"is-tgbot/internal/keys"
	"is-tgbot/pkg/logger"
	"os"
)

const initError = "bot init error"

var provider *command.Provider

var cache *redis.Client

func Start(ctx context.Context, redis *redis.Client) {
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
		provider = command.NewCommandProvider(b)
		cache = redis
		b.Start(ctx)
	}
}

func callbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

	if update == nil {
		return
	}

	if update.CallbackQuery == nil || update.CallbackQuery.Message.Message == nil {
		return
	}

	callback := *update.CallbackQuery

	if _, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: callback.ID,
		ShowAlert:       false,
	}); err != nil {
		logger.Log().Error(err, "Callback error")
	}

	logger.Log().Infof("TgID: %v, button: %s", callback.From.ID, callback.Data)

	data := callback.Data

	if err := deleteMessage(ctx, b, callback); err != nil {
		logger.Log().Errorf(err, "delete message error, chat id: %d", callback.From.ID)
	}

	handler := provider.Get(data)

	if handler == nil {
		logger.Log().Errorf(errors.New("handler not found"), "Error handle command: %s", data)
		return
	}

	go handler.Handle(ctx, b, update, cache)
}

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	go provider.Get(keys.Menu).Handle(ctx, b, update, cache)
}

func deleteMessage(ctx context.Context, b *bot.Bot, callback models.CallbackQuery) error {
	_, err := b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    callback.From.ID,
		MessageID: callback.Message.Message.ID,
	})

	return err
}
