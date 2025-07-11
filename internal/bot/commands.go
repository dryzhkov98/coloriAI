package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// Command represents an executable action or operation within the system, typically triggered by user input or program logic.
type Command struct{}

func (b *Bot) handleCommand(msg *tgbotapi.Message) error {
	switch msg.Command() {
	case "start":
		err := b.OnStart(msg)
		if err != nil {
			return err
		}
	case "help":
		b.OnHelp(msg)

	}
	return nil
}
