
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
			continue
		}
		ProcessTask(task)
	}
}


type Result struct {
	Text string
	Err error
}


func Predict(task *Task) (chan string, chan Result) {

	stream := make(chan string)
	result := make(chan Result)

	go func(){
		callback := func(token string) bool {
			select {
			case stream <- token:
				return true
			case <- task.Stop:
				return false
			}
		}
	
		text, err := l.Predict(
			task.Question,
			llama.Debug,
			llama.SetTokenCallback(callback),
			llama.SetTokens(nTokens), 
			llama.SetThreads(nCpu),
			llama.SetTopK(90),
			llama.SetTopP(0.86),
			llama.SetStopWords(StopWord),
		)
		close(stream)
		result <- Result{text, err}
	}()
	
	return stream, result
}