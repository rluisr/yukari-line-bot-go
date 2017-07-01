package controllers

import (
	_ "fmt"
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"

	m "github.com/rluisr/yukari-line-bot-go/models"
	c "github.com/rluisr/yukari-line-bot-go/conf"
)

func CreateInstance() (*linebot.Client) {
	bot, err := linebot.New(
		c.CHANNEL_SECRET,
		c.CHANNEL_TOKEN,
	)
	if err != nil {
		panic(err)
	}

	return bot
}

// arg: HP or FC {String}
func PushMessage(arg string) {
	var msg string

	if arg == "HP" {
		latestTitle, latestUrl := m.HPGetLatest()
		msg = c.CreateUpdateMsg("HP", latestTitle, latestUrl)
	}
	if arg == "FC" {
		latestTitle, latestUrl := m.FCGetLatest()
		msg = c.CreateUpdateMsg("FC", latestTitle, latestUrl)
	}

	m.PushMessage(msg)
}

func StartServer() {
	bot := CreateInstance()

	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}

		for _, event := range events {
			var userId string = event.Source.UserID

			switch event.Type {
			case linebot.EventTypeMessage:
				m.RecieveMessage(event.Message)
			case linebot.EventTypeFollow:
				m.RecieveFollow(userId, event)
			case linebot.EventTypeUnfollow:
				m.RecieveUnfollow(userId)
			}
		}
	})

	if err := http.ListenAndServe(":1337", nil); err != nil {
		log.Fatal(err)
	}
}
