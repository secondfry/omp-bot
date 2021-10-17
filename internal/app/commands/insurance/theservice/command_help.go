package theservice

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var help = `/help__insurance__theservice — print list of commands
/get__insurance__theservice <idx> — get an entity
/list__insurance__theservice — get a list of entities with pagination
/delete__insurance__theservice <idx> — delete an existing entity

/new__insurance__theservice <json> — create a new entity
/edit__insurance__theservice <idx> <json> — edit an entity`

func (c *InsuranceTheServiceCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, help)

	c.bot.Send(msg)
}
