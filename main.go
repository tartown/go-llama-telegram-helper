package main

import (
	"log"
	"os"
	"runtime"
	"strconv"

	llama "github.com/go-skynet/go-llama.cpp"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


var apiToken = os.Getenv("TG_TOKEN")
var modelPath = os.Getenv("MODEL_PATH")
var nTokens int
var nCpu int

var SingleMessagePrompt string
var ReplyMessagePrompt string
var StopWord = os.Getenv("STOP_WORD")

var l *llama.LLama
var bot *tgbotapi.BotAPI
var qu *TaskQueue
var currentTask *Task


func main() {
	var err error

	if apiToken == "" || modelPath == "" {
		log.Fata