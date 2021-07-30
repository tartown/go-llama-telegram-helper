
package main

import (
	"log"
	"strings"
	"time"

	llama "github.com/go-skynet/go-llama.cpp"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


var stopButton = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Stop", "/stop"),
	),
)

func ProcessQueue() {
	for {
		task, err := qu.Dequeue()
		currentTask = task
		if err == ErrQueueEmpty {
			time.Sleep(time.Second * 2)