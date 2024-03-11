package service

import (
	"gopkg.in/telebot.v3"
	"strconv"
	"time"
)

type TelegramBotService struct{}

func (e *TelegramBotService) SendTgMessage(token, chatid, content string) (res string, err error) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return "", err
	}

	chatID, err := strconv.ParseInt(chatid, 10, 64)
	if err != nil {
		bot.Stop()
		return "chatID cover failed", err
	}

	_, err = bot.Send(&telebot.Chat{ID: chatID}, content)
	if err != nil {
		return "telebot send message failed", err
	}

	return "telebot send message success", nil
}
