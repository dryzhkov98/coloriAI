package bot

import (
	keyboardbuilder "coloriAI/internal/keyboard"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) OnStart(msg *tgbotapi.Message) error {
	startKeyboard := keyboardbuilder.NewReplyBuilder().AddRow("📔 Дневник питания").AddRow("📊 Анализ Питания").AddRow("🗂 Справка").Build()

	reply := tgbotapi.NewMessage(msg.Chat.ID, "Привет! Я помощник в контроле твоего рациона")
	reply.ReplyMarkup = startKeyboard

	_, err := b.api.Send(reply)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) OnHelp(msg *tgbotapi.Message) {
	//... implement future
	fmt.Println("OnHelp")
}
