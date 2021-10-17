package theservice

import (
	"errors"
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *InsuranceTheServiceCommander) Get(msg *tgbotapi.Message) error {
	data := msg.CommandArguments()

	if data == "" {
		c.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Please provide arguments, like:\n/get__insurance__theservice 2"))
		return errors.New("empty data")
	}

	idx, err := strconv.ParseInt(data, 10, 0)
	if idx < 0 || err != nil {
		c.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("Unable to parse ID: %s", err)))
		return err
	}

	product, err := c.service.Describe(uint64(idx))
	if err != nil {
		c.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("Couldn't get TheService: %s", err)))
		return err
	}

	c.bot.Send(tgbotapi.NewMessage(
		msg.Chat.ID,
		product.String(),
	))
	return nil
}
