
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