package subdomain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *InsuranceSubdomainCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list products",
	)

	c.bot.Send(msg)
}
