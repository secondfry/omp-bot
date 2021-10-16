package theservice

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (*InsuranceTheServiceCommander) New(*tgbotapi.Message) error {
	return errors.New("not implemented")
}
