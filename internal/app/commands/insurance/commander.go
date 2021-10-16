package insurance

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/secondfry/omp-bot/internal/app/commands/insurance/subdomain"
	"github.com/secondfry/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type InsuranceCommander struct {
	bot                *tgbotapi.BotAPI
	subdomainCommander Commander
}

func NewInsuranceCommander(
	bot *tgbotapi.BotAPI,
) *InsuranceCommander {
	return &InsuranceCommander{
		bot: bot,
		// subdomainCommander
		subdomainCommander: subdomain.NewInsuranceSubdomainCommander(bot),
	}
}

func (c *InsuranceCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "subdomain":
		c.subdomainCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("InsuranceCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *InsuranceCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "subdomain":
		c.subdomainCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("InsuranceCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
