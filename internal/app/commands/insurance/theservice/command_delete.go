package theservice

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *InsuranceTheServiceCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if idx < 0 || err != nil {
		log.Println("wrong args", args)
		return
	}

	status, err := c.service.Remove((uint64(idx)))
	if !status || err != nil {
		log.Printf("fail to get product with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Removed %d", idx),
	)

	c.bot.Send(msg)
}
