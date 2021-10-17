package theservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	theservice "github.com/ozonmp/omp-bot/internal/service/insurance/theservice"
)

func (c *InsuranceTheServiceCommander) Edit(msg *tgbotapi.Message) error {
	newservice := theservice.TheService{}

	println(msg.Text)

	data := msg.CommandArguments()
	if data == "" {
		c.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Please provide arguments, like:\n/edit__insurance__theservice 1 {\"title\":\"example\"}"))
		return errors.New("empty data")
	}

	parts := strings.Split(data, " ")
	if len(parts) < 2 {
		c.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Please provide arguments, like:\n/edit__insurance__theservice 1 {\"title\":\"example\"}"))
		return errors.New("too few arguments")
	}

	idx, err := strconv.ParseInt(parts[0], 10, 0)
	if idx < 0 || err != nil {
		c.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("Unable to parse ID: %s", err)))
		return err
	}

	json_data := strings.Join(parts[1:], " ")
	err = json.Unmarshal([]byte(json_data), &newservice)
	if err != nil {
		c.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("Couldn't parse JSON: %s", err)))
		return err
	}

	err = c.service.Update(uint64(idx), newservice)
	if err != nil {
		c.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("Couldn't create TheService: %s", err)))
		return err
	}

	c.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("Changed #%d", idx)))
	return nil
}
