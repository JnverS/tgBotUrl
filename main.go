package main

import (
	"flag"
	"fmt"
	"log"

	"bot/clients/telegram"
)

func main() {
	tgClient := telegram.New("api.telegram.org", mustToken())
	fmt.Println(tgClient)
	//token = flags.Get(token)
	//tgClient = telegram.New(token)
	//fetcher = fetcher.New(tgClient)
	//processor = processor.New(tgClient)
	// consumer.Start(fetcher, processor)
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
