package theservice

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/insurance/theservice"
)

type TheServiceCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

type InsuranceTheServiceCommander struct {
	bot     *tgbotapi.BotAPI
	service service.TheServiceService
}

func NewInsuranceTheServiceCommander(
	bot *tgbotapi.BotAPI,
	service service.TheServiceService,
) *InsuranceTheServiceCommander {
	return &InsuranceTheServiceCommander{
		bot:     bot,
		service: service,
	}
}

func (c *InsuranceTheServiceCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	var err error

	switch callbackPath.CallbackName {
	case "list":
		err = c.CallbackList(callback, callbackPath)
	default:
		log.Printf("InsuranceTheServiceCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}

	if err != nil {
		log.Printf("HandleCallback failed! [%+v] [%+v]", callbackPath, err)
	}
}

func (c *InsuranceTheServiceCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	var err error

	switch commandPath.CommandName {
	case "help":
		err = c.Help(msg)
	case "get":
		err = c.Get(msg)
	case "list":
		err = c.List(msg)
	case "delete":
		err = c.Delete(msg)
	case "new":
		err = c.New(msg)
	case "edit":
		err = c.Edit(msg)
	default:
		err = c.Default(msg)
	}

	if err != nil {
		log.Printf("HandleCommand failed! [%+v] [%+v]", commandPath, err)
	}
}
