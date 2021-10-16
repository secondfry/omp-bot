package theservice

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const START = 0
const PAGER = 3

func (c *InsuranceTheServiceCommander) ListText(cursor uint64, limit uint64) string {
	outputMsgText := "Here all the products: \n\n"

	products := c.theserviceService.List(cursor, limit)
	for _, p := range products {
		outputMsgText += p.Title
		outputMsgText += "\n"
	}

	return outputMsgText
}

func PrepareCallbackPath(idx int) string {
	if idx < 0 {
		idx = 0
	}

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: idx,
	})

	callbackPath := path.CallbackPath{
		Domain:       "insurance",
		Subdomain:    "theservice",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	return callbackPath.String()
}

func (c *InsuranceTheServiceCommander) ListKeyboard(cursor uint64) tgbotapi.InlineKeyboardMarkup {
	var row []tgbotapi.InlineKeyboardButton

	if c.theserviceService.HasBefore(cursor) {
		row = append(row, tgbotapi.NewInlineKeyboardButtonData("Prev page", PrepareCallbackPath(int(cursor)-PAGER)))
	}

	if c.theserviceService.HasAfter(cursor + PAGER) {
		row = append(row, tgbotapi.NewInlineKeyboardButtonData("Next page", PrepareCallbackPath(int(cursor)+PAGER)))
	}

	markup := tgbotapi.NewInlineKeyboardMarkup(
		row,
	)

	return markup
}

func (c *InsuranceTheServiceCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := c.ListText(START, PAGER)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
	msg.ReplyMarkup = c.ListKeyboard(START)
	c.bot.Send(msg)
}
