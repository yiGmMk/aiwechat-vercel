package chat

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
	"github.com/pwh-pwh/aiwechat-vercel/db"
)

func TestGpt(t *testing.T) {
	godotenv.Load("../conf/.env")
	db.ChatDbInstance, _ = db.GetChatDb()
	bot := GetChatBot("gpt")

	res := bot.Chat("testUser", "用10个字描述你的能力")

	fmt.Println(res)
}
