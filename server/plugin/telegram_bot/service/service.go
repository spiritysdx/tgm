package service

import (
	"gopkg.in/telebot.v3"
	"strconv"
	"time"
)

type TelegramBotService struct{}

func (e *TelegramBotService) SendTgMessage(token, chat_id, content, message_type string) (res string, err error) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 3 * time.Second}, // 机器人每分钟最多发送20条消息
	})
	if err != nil {
		return "", err
	}
	chatID, err := strconv.ParseInt(chat_id, 10, 64)
	if err != nil {
		bot.Stop()
		return "chatID cover failed", err
	}
	var parseMode telebot.ParseMode
	switch message_type {
	case "html":
		parseMode = telebot.ModeHTML
	case "markdown":
		parseMode = telebot.ModeMarkdown
	default:
		parseMode = telebot.ModeMarkdown
	}
	_, err = bot.Send(&telebot.Chat{ID: chatID}, content, telebot.SendOptions{ParseMode: parseMode})
	if err != nil {
		return "telebot send message failed", err
	}
	return "telebot send message success", nil
}
