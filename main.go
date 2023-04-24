package main

import (
	"assignment4/app"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

const (
	telegramBotToken = "5834358752:AAFOaJEKEDi27yWR90uRX9VCY0DmolaakSI"
)

func main() {
	bot := app.NewBot(telegramBotToken)

	log.Printf("Starting bot %s", bot.BotAPI.Self.UserName)
	updates, err := bot.BotAPI.GetUpdatesChan(tgbotapi.NewUpdate(0))
	if err != nil {
		log.Fatal(err)
	}

	counter := app.NewCounter()
	//wg := sync.WaitGroup{}
	//wg.Add(1)

	//count := make(chan int)
	//
	//go func() {
	//	var n int
	//	for {
	//		n++
	//		count <- n
	//	}
	//}()

	//go func() {
	//	defer wg.Done()
	for update := range updates {
		if update.Message == nil {
			continue
		}

		text := strings.ToLower(update.Message.Text)
		if text != "image" && text != "/image" {
			continue
		}

		bot.HandleUpdate(update.Message.Chat.ID)
		//n := <-count
		//log.Printf("Sent %d images", n)
		counter.Increment()
	}
	//}()
	//wg.Wait()
}
