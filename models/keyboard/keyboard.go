package keyboard

import "github.com/go-telegram/bot/models"

var MainMenu = [][]models.InlineKeyboardButton{
	{
		{Text: "📱Купить номер", CallbackData: "button_buyNumber"},
	},
	{
		{Text: "💸Пополнить баланс", CallbackData: "button_deposit"},
		{Text: "⚙️Настройки", CallbackData: "button_settings"},
	},
	{
		{Text: "📑История активаций", CallbackData: "button_activationLog"},
		{Text: "📑История пополнений", CallbackData: "button_depositLog"},
	}, {
		{Text: "👤Профиль", CallbackData: "button_profile"},
	},
}
