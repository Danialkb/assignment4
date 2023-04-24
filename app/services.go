package app

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
)

const unsplashAccessKey = "YnEM6VRmv9EkYiUQ24jy0sUSW4NblVMjYZGm4P-UNFg"

type UnsplashPhotoProvider interface {
	GetRandomPhoto() (*UnsplashPhoto, error)
}

type UnsplashPhotoProviderV1 struct {
	accessKey string
}

func (s UnsplashPhotoProviderV1) GetRandomPhoto() (*UnsplashPhoto, error) {
	url := fmt.Sprintf("https://api.unsplash.com/photos/random?client_id=%s", s.accessKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var photo UnsplashPhoto
	err = json.NewDecoder(resp.Body).Decode(&photo)
	if err != nil {
		return nil, err
	}

	if len(photo.Errors) > 0 {
		return nil, fmt.Errorf(photo.Errors[0].Detail)
	}

	return &photo, nil
}

type MessageSenderV1 struct {
	unsplashProvider UnsplashPhotoProvider
}

func (m MessageSenderV1) SendPhoto(chatID int64, bot *tgbotapi.BotAPI) {
	photo, _ := m.unsplashProvider.GetRandomPhoto()
	resp, err := http.Get(photo.URL["regular"])
	fmt.Println(photo.URL["regular"])
	if err != nil {
		log.Printf("Error downloading photo: %v", err)
	}

	defer resp.Body.Close()
	msg1 := tgbotapi.NewPhotoUpload(chatID, tgbotapi.FileReader{
		Name:   photo.ID + ".jpg",
		Reader: resp.Body,
		Size:   -1,
	})
	msg1.Caption = "Here's your random photo!"
	_, err = bot.Send(msg1)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}

func NewMessageSender() *MessageSenderV1 {
	return &MessageSenderV1{unsplashProvider: UnsplashPhotoProviderV1{accessKey: unsplashAccessKey}}
}
