package command

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/redis/go-redis/v9"
	"is-tgbot/internal/client"
	"is-tgbot/internal/keys"
	"is-tgbot/internal/model"
	"is-tgbot/internal/model/keyboard"
	"is-tgbot/internal/storage"
	"is-tgbot/internal/utils"
	"is-tgbot/pkg/logger"
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

func (c *Menu) Handle(ctx context.Context, b *bot.Bot, update *models.Update, cache *redis.Client) {
	var balance model.BalanceGetResponse

	if cacheBalance := storage.GetStruct[model.BalanceGetResponse](cache, ctx, keys.RedisBalance); cacheBalance != nil {
		balance = *cacheBalance
	} else {
		balance = getBalanceFromServer(ctx, update, cache)
	}

	text := fmt.Sprintf("💸Ваш баланс: %s %s.\n🌍Страна: Россия\n🌐Оператор: -", *balance.Sum, balance.Currency)
	utils.SendKeyboard(ctx, keyboard.MainMenu, update, text, b)
}

func getBalanceFromServer(ctx context.Context, update *models.Update, cache *redis.Client) model.BalanceGetResponse {
	request := getBalanceRequest(ctx, update, cache)
	balance, err := client.GetBalance(request)
	sum, curr := "0.0", "RUB"
	if err != nil {
		logger.Log().Error(err, "get balance error:")
		balance.Sum = &sum
		balance.Currency = curr
	}
	storage.SetStruct(cache, ctx, keys.RedisBalance, balance)
	return balance
}

func getBalanceRequest(ctx context.Context, update *models.Update, cache *redis.Client) model.BalanceGetRequest {
	var customerId *int64
	var telegramId *int64
	currencyCode := keys.CurrencyRub

	if balance := storage.GetStruct[model.BalanceGetResponse](cache, ctx, keys.RedisBalance); balance != nil {
		currencyCode = balance.Currency
	}

	if customer := storage.GetStruct[model.CustomerResponse](cache, ctx, keys.RedisCustomer); customer != nil {
		if customer.ID != 0 {
			customerId = &customer.ID
		}
		telegramId = customer.TelegramID
	} else {
		telegramId = getChatId(update)
	}

	return model.BalanceGetRequest{
		CustomerId:   customerId,
		TelegramId:   telegramId,
		CurrencyCode: &currencyCode,
	}
}
