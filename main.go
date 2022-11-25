package main

import (
	event_consumer "UrlBot/comsumer/event-consumer"
	"UrlBot/events/telegram"
	"UrlBot/storage/files"
	"flag"
	"log"

	tgClient "UrlBot/clients/telegram"
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New("api.telegram.org", mustToken()),
		files.New("files_storage"),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, 100)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tgToken",
		"",
		"token for acsess to tg bot",
	)
	flag.Parse()
	if *token == "" {
		log.Fatal("token is not valid")
	}
	return *token
}
