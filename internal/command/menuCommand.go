package command

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/wasd0/is-common/pkg/logger"
	"is-tgbot/internal/client"
	"is-tgbot/internal/keys"
	"is-tgbot/internal/model"
	"is-tgbot/internal/model/keyboard"
	"is-tgbot/internal/service"
	"is-tgbot/internal/utils"
)

type Menu struct {
	cache service.CacheService
}

func NewMenuCommand(cache service.CacheService) *Menu {
	return &Menu{cache: cache}
}

func (c *Menu) GetCommand() string {
	return keys.Menu
}

func (c *Menu) Handle(ctx context.Context, b *bot.Bot, update *models.Update) {
	balance := model.BalanceGetResponse{}
	c.cache.GetStruct(ctx, *getChatId(update), keys.RedisBalance, &balance)

	if balance.BalanceId == 0 {
		balance = c.getBalanceFromServer(ctx, update)
	}

	text := fmt.Sprintf("💸Ваш баланс: %s %s.\n🌍Страна: Россия\n🌐Оператор: -", *balance.Sum, balance.Currency)
	utils.SendKeyboard(ctx, keyboard.MainMenu, update, text, b)
}

func (c *Menu) getBalanceFromServer(ctx context.Context, update *models.Update) model.BalanceGetResponse {
	request := c.getBalanceRequest(ctx, update)
	balance, err := client.GetBalance(request)
	sum, curr := "0.0", "RUB"
	if err != nil {
		logger.Log().Error(err, "get balance error:")
	}
	if balance == nil {
		balance = &model.BalanceGetResponse{
			BalanceId: 0,
			Currency:  curr,
			Sum:       &sum,
		}
	}
	c.cache.SetStruct(ctx, *getChatId(update), keys.RedisBalance, balance)
	return *balance
}

func (c *Menu) getBalanceRequest(ctx context.Context, update *models.Update) model.BalanceGetRequest {
	var customerId *int64
	var telegramId *int64
	currencyCode := keys.CurrencyRub

	balance := model.BalanceGetResponse{}
	c.cache.GetStruct(ctx, *getChatId(update), keys.RedisBalance, &balance)
	if balance.BalanceId != 0 {
		currencyCode = balance.Currency
	}

	customer := model.CustomerResponse{}
	c.cache.GetStruct(ctx, *getChatId(update), keys.RedisCustomer, &customer)

	if customer.ID != 0 {
		customerId = &customer.ID
	} else if customer.TelegramID != nil {
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
