
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

// This function is a mess
func ProcessTask(task *Task) {

	log.Printf("Start processing task from user %d\n", task.UserID)
	log.Printf("The prompt is:\n%s\n", task.Question)

	// Start prediction
	stream, result :=  Predict(task)

	// Resulting generated text
	var answer string

	var counter int
	var issent bool
	for {
		select {
		case token := <- stream: 
			if !issent && strings.TrimSpace(token) != "" {
				answer += token
				msg := tgbotapi.NewMessage(task.UserID, answer)
				msg.ReplyMarkup = &stopButton
				sent, err := bot.Send(msg)
				if err != nil {
					log.Println("[ProcessTask] error sending answer:", err)
					continue