package service

import (
	"errors"
	"fmt"
	"gopkg.in/telebot.v3"
	"strconv"
	"strings"
	"time"
)

type TelegramBotService struct{}

func (e *TelegramBotService) SendTgMessage(tokens, chatId, content, messageType string) (res string, err error) {
	// 英文逗号分隔token
	tokenList := strings.Split(tokens, ",")
	var lastError error
	// 多个token的时候轮询直至发送成功
	for index, token := range tokenList {
		bot, err := telebot.NewBot(telebot.Settings{
			Token:  token,
			Poller: &telebot.LongPoller{Timeout: 3 * time.Second},
		})
		if err != nil {
			// 初始化失败
			lastError = errors.New(fmt.Sprintf("bot initialise failed for token%d: %v", index, err))
			continue
		}
		chatID, err := strconv.ParseInt(chatId, 10, 64)
		if err != nil {
			// chat_id 转换失败
			lastError = errors.New(fmt.Sprintf("chatID cover failed for token%d: %v", index, err))
			continue
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
		res, err = bot.Send(&telebot.Chat{ID: chatID}, content, telebot.SendOptions{ParseMode: parseMode})
		if err != nil {
			// 发送失败
			lastError = errors.New(fmt.Sprintf("telebot send message failed for token%d: %v", index, err))
			continue
		}
		// 发送成功
		return fmt.Sprintf("telebot send message success with token%d: %v", index, res), nil
	}
	// 全部token发送都失败了
	return "telebot send message failed", lastError
}
