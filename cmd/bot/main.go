package main

import (
	"github.com/bannovdaniil/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var productService = product.NewService()

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file TOKEN=")
	}
	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "help":
			helpCommand(bot, update.Message)
		case "list":
			listCommand(bot, update.Message)
		default:
			defaultAction(bot, update.Message)
		}
	}
}

// listCommand - action for /help command
//
// send help message
func listCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	productsListText := "List of our products: \n\n"
	for i, item := range productService.List() {
		productsListText += strconv.Itoa(i+1) + ". " + item.Title + "\n"
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, productsListText)
	bot.Send(msg)
}

// helpCommand - action for /help command
//
// send help message
func helpCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "/help - help tutorial.")
	bot.Send(msg)
}

// defaultAction - action for all other messages
//
// send repeat message
func defaultAction(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, "You send me text: "+message.Text)
	msg.ReplyToMessageID = message.MessageID

	bot.Send(msg)

}
