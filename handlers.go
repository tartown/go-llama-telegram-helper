
package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func ProcessUpdate(update tgbotapi.Update) {
	// If we've gotten a message update.
	if update.Message != nil {

		msg := tgbotapi.MessageConfig{
			BaseChat: tgbotapi.BaseChat{
				ChatID:           update.Message.Chat.ID,
			},
			DisableWebPagePreview: true,
		}

		if update.Message.Text == "/start" {
			msg.Text = "Just ask question"
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
			return
		}

		if update.Message.Text == "/queue" {
			_, n := qu.Load(update.Message.From.ID)

			switch n {
			case -1:
				if currentTask != nil && currentTask.UserID == update.Message.From.ID {
					msg.Text = "It's your turn now!!!"
				} else {
					msg.Text = "Hey! You haven't asked question yet!"
				}
			case 0:
				msg.Text = "Hold a second, you're next"
			default:
				msg.Text = fmt.Sprintf("Hold on! Your queue is %d", n)
			}
			
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
			return