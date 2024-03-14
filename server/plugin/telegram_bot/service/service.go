package service

import (
	"gopkg.in/telebot.v3"
	"strconv"
	"time"
)

type TelegramBotService struct{}

func (e *TelegramBotService) SendTgMessage(token, chatId, content, messageType string) (res string, err error) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 3 * time.Second}, // 机器人每分钟最多发送20条消息
	})
	if err != nil {
		return "bot initialise failure", errors.New(fmt.Sprintf("bot initialise failed: %v", err))
	}
	chatID, err := strconv.ParseInt(chatId, 10, 64)
	if err != nil {
		bot.Stop()
		return "chatID cover failed", errors.New(fmt.Sprintf("chatID cover failed: %v", err))
	}
	var parseMode telebot.ParseMode
	switch messageType {
	case "html":
		parseMode = telebot.ModeHTML
	case "markdown":
		parseMode = telebot.ModeMarkdown
	default:
		parseMode = telebot.ModeMarkdown
	}
	_, err = bot.Send(&telebot.Chat{ID: chatID}, content, telebot.SendOptions{ParseMode: parseMode})
	if err != nil {
		return "telebot send message failed", errors.New(fmt.Sprintf("telebot send message failed: %v", err))
	}
	return "telebot send message success", nil
}
