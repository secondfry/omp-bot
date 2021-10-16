package subdomain

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/secondfry/omp-bot/internal/app/path"
	"github.com/secondfry/omp-bot/internal/service/insurance/subdomain"
)

type InsuranceSubdomainCommander struct {
	bot              *tgbotapi.BotAPI
	subdomainService *subdomain.Service
}

func NewInsuranceSubdomainCommander(
	bot *tgbotapi.BotAPI,
) *InsuranceSubdomainCommander {
	subdomainService := subdomain.NewService()

	return &InsuranceSubdomainCommander{
		bot:              bot,
		subdomainService: subdomainService,
	}
}

func (c *InsuranceSubdomainCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("InsuranceSubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *InsuranceSubdomainCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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
