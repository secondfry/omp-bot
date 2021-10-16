package insurance

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	theService "github.com/ozonmp/omp-bot/internal/app/commands/insurance/theservice"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type InsuranceCommander struct {
	bot                 *tgbotapi.BotAPI
	theServiceCommander Commander
}

func NewInsuranceCommander(
	bot *tgbotapi.BotAPI,
) *InsuranceCommander {
	return &InsuranceCommander{
		bot: bot,
		// theServiceCommander
		theServiceCommander: theService.NewInsuranceTheServiceCommander(bot),
	}
}

func (c *InsuranceCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "theservice":
		c.theServiceCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("InsuranceCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *InsuranceCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "theservice":
		c.theServiceCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("InsuranceCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
