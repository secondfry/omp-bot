package theservice

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	theservice "github.com/ozonmp/omp-bot/internal/service/insurance/theservice"
)

func (c *InsuranceTheServiceCommander) New(msg *tgbotapi.Message) error {
	newservice := theservice.TheService{}

	println(msg.Text)

	err := json.Unmarshal([]byte(msg.Text), &newservice)
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
