package theservice

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var help = `/help__insurance__service — print list of commands
/get__insurance__service — get a entity
/list__insurance__service — get a list of your entity (💎: with pagination via telegram keyboard)
/delete__insurance__service — delete an existing entity

/new__insurance__service — create a new entity // not implemented (💎: implement list fields via arguments)
/edit__insurance__service — edit a entity      // not implemented`

func (c *InsuranceTheServiceCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, help)

	c.bot.Send(msg)
}
