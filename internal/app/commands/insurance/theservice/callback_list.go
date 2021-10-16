package theservice

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c *InsuranceTheServiceCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)

	outputMsgText := c.ListText(uint64(parsedData.Offset), PAGER)
	msg := tgbotapi.NewEditMessageText(callback.Message.Chat.ID, callback.Message.MessageID, outputMsgText)
	markup := c.ListKeyboard(uint64(parsedData.Offset))
	msg.ReplyMarkup = &markup

	c.bot.Send(msg)
}
