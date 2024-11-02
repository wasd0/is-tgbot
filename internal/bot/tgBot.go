package bot

import (
	"context"
	"errors"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/wasd0/is-common/pkg/logger"
	"is-tgbot/internal/command"
	"is-tgbot/internal/keys"
	"os"
)

const initError = "bot init error"

type isBot struct {
	provider *command.Provider
}

func Start(ctx context.Context, provider *command.Provider) {
	token := os.Getenv(keys.EnvToken)

	if token == "" {
		logger.Log().Fatal(errors.New(initError), "TOKEN environment variable is empty")
	}

	ib := &isBot{provider: provider}

	opts := []bot.Option{
		bot.WithDefaultHandler(ib.defaultHandler),
		bot.WithCallbackQueryDataHandler("button", bot.MatchTypePrefix, ib.callbackHandler),
		bot.WithWorkers(10),
		bot.WithDebugHandler(logger.Log().Infof),
		bot.WithErrorsHandler(func(err error) {
			logger.Log().Error(err, err.Error())
		}),
	}

	if b, err := bot.New(token, opts...); err != nil {
		logger.Log().Fatal(err, initError)
	} else {
		b.Start(ctx)
	}
}

func (ib *isBot) callbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

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

	handler := ib.provider.Get(data)

	if handler == nil {
		logger.Log().Errorf(errors.New("handler not found"), "Error handle command: %s", data)
		return
	}

	go handler.Handle(ctx, b, update)
}

func (ib *isBot) defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	go ib.provider.Get(keys.Menu).Handle(ctx, b, update)
}

func deleteMessage(ctx context.Context, b *bot.Bot, callback models.CallbackQuery) error {
	_, err := b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    callback.From.ID,
		MessageID: callback.Message.Message.ID,
	})

	return err
}
