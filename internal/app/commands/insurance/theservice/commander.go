package theservice

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/insurance/theservice"
)

type InsuranceTheServiceCommander struct {
	bot               *tgbotapi.BotAPI
	theserviceService *theservice.DummyTheServiceService
}

func NewInsuranceTheServiceCommander(
	bot *tgbotapi.BotAPI,
) *InsuranceTheServiceCommander {
	theserviceService := theservice.NewDummyTheServiceService()

	return &InsuranceTheServiceCommander{
		bot:               bot,
		theserviceService: theserviceService,
	}
}

func (c *InsuranceTheServiceCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("InsuranceTheServiceCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *InsuranceTheServiceCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	default:
		c.Default(msg)
	}
}
