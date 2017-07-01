package models

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"

	c "github.com/rluisr/yukari-line-bot-go/conf"
	l "github.com/rluisr/yukari-line-bot-go/libs"
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

func PushMessage(msg string) {
	bot := CreateInstance()

	users := l.GetUsers()

	for _, userId := range users {
		if _, err := bot.PushMessage(userId, linebot.NewTextMessage(msg)).Do(); err != nil {
			// 友だち追加してない場合エラーがでる
			// あえて何もしない
		}
	}
}

// Todo 何かする必要性ある？
func RecieveMessage(msg linebot.Message) {
	fmt.Println(msg)
}

// # 友達登録されたときに呼び出される
// # ===========================
func RecieveFollow(userId string) {
	l.InsertUserToDB("Users", userId)
}

// # ブロックされたときに呼び出される
// # ===========================
func RecieveUnfollow(userId string) {
	l.DeleteUserFromDB(userId)
}
