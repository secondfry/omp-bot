package subdomain

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/secondfry/omp-bot/internal/app/path"
)

func (c *InsuranceSubdomainCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the products: \n\n"

	products := c.subdomainService.List()
	for _, p := range products {
		outputMsgText += p.Title
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: 21,
	})

	callbackPath := path.CallbackPath{
		Domain:       "insurance",
		Subdomain:    "subdomain",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	c.bot.Send(msg)
}