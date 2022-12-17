[![Docker Pulls](https://img.shields.io/docker/pulls/tartown/go-llama-telegram-helper)](https://hub.docker.com/r/tartown/go-llama-telegram-helper)
[![Docker Image Size (tag)](https://img.shields.io/docker/image-size/tartown/go-llama-telegram-helper/latest)](https://hub.docker.com/r/tartown/go-llama-telegram-helper)

# 🦙 go-llama-telegram-helper

## What?
This is a chatbot for Telegram that utilizes the powerful [llama.cpp](https://github.com/ggerganov/llama.cpp). Give the live instance a try over here [@telellamabot](https://t.me/telellamabot)

## How?
[go-llama-telegram-helper](https://github.com/tartown/go-llama-telegram-helper) is crafted in Go and leverages [go-llama.cpp](https://github.com/go-skynet/go-llama.cpp) which is a binding to [llama.cpp](https://github.com/ggerganov/llama.cpp)

## Quick Start
Let's get started! The process is straightforward!

Parameters are accepted as environment variables.

1. `MODEL_PATH=/path/to/model`
2. `TG_TOKEN=your_telegram_bot_token_here`
3. `Q_SIZE=1000` - Task queue limit (optional: default 1000)
4. `N_TOKENS=1024` - Tokens to predi