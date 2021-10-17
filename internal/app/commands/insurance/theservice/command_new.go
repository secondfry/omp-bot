package theservice

import (
	"encoding/json"
	"errors"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	theservice "github.com/ozonmp/omp-bot/internal/service/insurance/theservice"
)

func (c *InsuranceTheServiceCommander) New(msg *tgbotapi.Message) error {
	newservice := theservice.TheService{}

	data := msg.CommandArguments()
	if data == "" {
		c.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Please provide arguments, like:\n/new__insurance__theservice {\"title\":\"example\"}"))
		return errors.New("empty data")
	}

	err := json.Unmarshal([]byte(data), &newservice)
	if err != nil {
		c.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("Couldn't parse JSON: %s", err)))
		return err
	}

	idx, err := c.service.Create(newservice)
	if err != nil {
		c.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("Couldn't create TheService: %s", err)))
		return err
	}

	c.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("Created #%d", idx)))
	return nil
}
