package model

type TelegramChat struct {
	Token string `json:"token"` // Telegram的Bot的Token
	ChatId  string `json:"chat_id"` // 发送的目标ID
	Content string `json:"content"` // 发送的内容
}
