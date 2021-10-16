package theservice

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (*InsuranceTheServiceCommander) Edit(*tgbotapi.Message) error {
	return errors.New("not implemented")
}
