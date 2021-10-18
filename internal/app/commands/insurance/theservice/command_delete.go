package theservice

import (
	"errors"
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *InsuranceTheServiceCommander) Delete(msg *tgbotapi.Message) error {
	data := msg.CommandArguments()

	if data == "" {
		c.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Please provide arguments, like:\n/delete__insurance__theservice 2"))
		return errors.New("empty data")
	}

	idx, err := strconv.ParseUint(data, 10, 64)
	if err != nil {
		c.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("Unable to parse ID: %s", err)))
		return err
	}

	status, err := c.service.Remove(idx)
	if !status || err != nil {
		c.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("Unable remove TheService: %s", err)))
		return err
	}

	_, err = c.bot.Send(tgbotapi.NewMessage(
		msg.Chat.ID,
		fmt.Sprintf("Removed %d", idx),
	))
	return err
}
