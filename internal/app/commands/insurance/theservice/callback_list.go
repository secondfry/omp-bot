package theservice

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c *InsuranceTheServiceCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) error {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		c.bot.Send(tgbotapi.NewMessage(callback.Message.Chat.ID, fmt.Sprintf("Unable to parse data: %+v", err)))
		return err
	}

	outputMsgText := c.ListText(uint64(parsedData.Offset), PAGER)
	msg := tgbotapi.NewEditMessageText(callback.Message.Chat.ID, callback.Message.MessageID, outputMsgText)
	markup := c.ListKeyboard(uint64(parsedData.Offset))
	msg.ReplyMarkup = &markup

	_, err = c.bot.Send(msg)
	return err
}
