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
