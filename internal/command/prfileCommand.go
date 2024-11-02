package command

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"is-tgbot/internal/client"
	"is-tgbot/internal/keys"
	"is-tgbot/internal/model"
	"is-tgbot/internal/model/keyboard"
	"is-tgbot/internal/service"
	"is-tgbot/internal/utils"
	"is-tgbot/pkg/logger"
)

type Profile struct {
	cache service.CacheService
}

func NewProfileCommand(cache service.CacheService) *Profile {
	return &Profile{cache: cache}
}

func (pc *Profile) GetCommand() string {
	return keys.Profile
}

func (pc *Profile) Handle(ctx context.Context, b *bot.Bot, update *models.Update) {
	request := getCustomerRequest(update)
	customer, err := client.GetCustomer(request)

	if err != nil {
		logger.Log().Error(err, "get customer info error:")
	} else {
		pc.cache.SetStruct(ctx, *getChatId(update), keys.RedisCustomer, customer)
	}

	id := customer.ID

	if customer.TelegramID != nil {
		id = *customer.TelegramID
	}

	createDate := customer.CreateDate

	text := fmt.Sprintf("Ваш ID: %d\n"+
		"Активировано номеров: 0\n"+
		"Арендовано номеров: 0\n"+
		"Дата создания аккаунта: %v", id, createDate.Format("2006-01-02"))

	utils.SendKeyboard(ctx, keyboard.ProfileMenu, update, text, b)
}

func getCustomerRequest(update *models.Update) model.CustomerGetRequest {
	return model.CustomerGetRequest{
		TelegramID: getChatId(update),
	}
}
