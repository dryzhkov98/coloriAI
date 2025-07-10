package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Command struct{}

func (b *Bot) handleCommand(msg *tgbotapi.Message) error {
	switch msg.Command() {
	case "start":
		err := b.OnStart(msg)
		if err != nil {
			return err
		}
		break
	case "help":
		b.OnHelp(msg)
		break
	}
	return nil
}
