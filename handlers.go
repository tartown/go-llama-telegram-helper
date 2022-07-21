
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
		}

		if chars := []rune(update.Message.Text); string(chars[0]) == "/" {
			msg.Text = "There is no such command"
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
			return
		}

		// Do enqueue task
		task := Task{
			UserID: update.Message.From.ID,
			Stop: make(chan bool),
		}

		if reply := update.Message.ReplyToMessage; reply != nil && reply.From.ID == bot.Self.ID {
			task.WrapPrevContext(reply.Text, update.Message.Text)
		} else {
			task.WrapInRoles(update.Message.Text)
		}
		
		
		n, err := qu.Enqueue(&task)
		log.Println(err)
		if err != nil {
			if err == ErrOnePerUser {
				msg.Text = "You've already asked your question. You can edit the existing one until it's your turn"
			}
			if err == ErrQueueLimit {
				msg.Text = fmt.Sprintf("Now queue is full %d/%d. Wait one slot to be free at least.\nCheck queue /stats", n, qu.Limit)
			}
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
			return
		}
		msg.Text = fmt.Sprintf("Your question registered! Your queue is %d/%d.\nYou can edit your message until it's your turn", n, qu.Limit)
		sent, err := bot.Send(msg)
		if err != nil {
			log.Println(err)