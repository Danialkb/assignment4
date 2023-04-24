package app

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"sync"
)

type UnsplashPhoto struct {
	ID     string            `json:"id"`
	URL    map[string]string `json:"urls"`
	Errors []struct {
		Detail string `json:"detail"`
	} `json:"errors"`
}

type UnsplashClient struct {
}

type Counter struct {
	counter int
	mu      *sync.Mutex
}

func NewCounter() *Counter {
	var cnt int
	var mut sync.Mutex

	return &Counter{counter: cnt, mu: &mut}
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.counter++
	fmt.Println(c.counter)
	c.mu.Unlock()
}

type Bot struct {
	BotAPI        *tgbotapi.BotAPI
	messageSender *MessageSenderV1
}

func NewBot(telegramToken string) *Bot {
	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Fatal(err)
	}
	return &Bot{BotAPI: bot, messageSender: NewMessageSender()}
}

func (b Bot) HandleUpdate(chatID int64) {
	b.messageSender.SendPhoto(chatID, b.BotAPI)

}
