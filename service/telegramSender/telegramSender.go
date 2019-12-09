package telegramSender

import (
	"crypto/sha256"
	"currencyParser/cache"
	"currencyParser/service/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const SEND_EACH_SECONDS = 300

var telegramSender TelegramSender

type TelegramSender struct {
	Enabled bool
	Token   string
	ChatId  int64
}

func init() {
	telegramSender = TelegramSender{
		Enabled: config.GetConfig().Telegram.Enabled,
		Token:   config.GetConfig().Telegram.Token,
		ChatId:  config.GetConfig().Telegram.ChatId,
	}
}

func GetSender() *TelegramSender {
	return &telegramSender
}

func (sender *TelegramSender) SendMessage(msg string) error {
	if !sender.Enabled {
		return nil
	}

	isSendedHash := sha256.Sum256([]byte(msg))
	_, err := cache.Get(string(isSendedHash[:]))
	if err == nil {
		return nil
	}
	cache.Set(string(isSendedHash[:]), []byte("true"), SEND_EACH_SECONDS)

	bot, err := tgbotapi.NewBotAPI(sender.Token)
	if err != nil {
		return err
	}

	bot.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	messageConfig := tgbotapi.NewMessage(sender.ChatId, msg)
	_, err = bot.Send(messageConfig)

	return err
}