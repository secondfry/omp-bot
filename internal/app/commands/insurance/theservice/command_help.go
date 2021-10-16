package theservice

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var help = `/help__insurance__theservice — print list of commands
/get__insurance__theservice — get a entity
/list__insurance__theservice — get a list of your entity (💎: with pagination via telegram keyboard)
/delete__insurance__theservice — delete an existing entity

/new__insurance__theservice — create a new entity // not implemented (💎: implement list fields via arguments)
/edit__insurance__theservice — edit a entity      // not implemented`

func (c *InsuranceTheServiceCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, help)

	c.bot.Send(msg)
}
