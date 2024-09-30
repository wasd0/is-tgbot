package keyboard

import (
	"github.com/go-telegram/bot/models"
	"is-tgbot/internal/keys"
)

var MainMenu = [][]models.InlineKeyboardButton{
	{
		{Text: "📱Купить номер", CallbackData: keys.BuyNumber},
	},
	{
		{Text: "💸Пополнить баланс", CallbackData: keys.Deposit},
		{Text: "⚙️Настройки", CallbackData: keys.Settings},
	},
	{
		{Text: "📑История активаций", CallbackData: keys.ActivationLog},
		{Text: "📑История пополнений", CallbackData: keys.DepositLog},
	}, {
		{Text: "👤Профиль", CallbackData: keys.Profile},
	},
}

var ProfileMenu = [][]models.InlineKeyboardButton{
	{
		{Text: "💸Пополнить баланс", CallbackData: keys.Deposit},
	},
	{
		{Text: "Ⓜ️Меню", CallbackData: keys.Menu},
	},
}
