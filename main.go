package main

import (
	"log"

	"github.com/k0yote/simplebot/bot"
)

func main() {
	var (
		requestURL   = "http://localhost:8080/payment/v1/shop/free/purchase"
		requestCount = 1000

		headerKey   = "x-service-api-key"
		headerValue = "58895e41-2789-1142-1fc8-48fcc7d2b9e5"
	)

	bot, err := bot.NewBot(requestURL, requestCount, &bot.CustomHeader{
		Key:   headerKey,
		Value: headerValue,
	})

	if err != nil {
		log.Fatalln(err)
	}

	if err := bot.BotRequest(); err != nil {
		log.Fatalln(err)
	}

	log.Println("bot end")
}
