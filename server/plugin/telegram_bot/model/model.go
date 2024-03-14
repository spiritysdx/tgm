package model

type TelegramChat struct {
	Token       string `json:"token"`        // Telegram的Bot的Token
	ChatId      string `json:"chat_id"`      // 发送的目标ID
	Content     string `json:"content"`      // 发送的内容
	MessageType string `json:"message_type"` // 发送内容的文本类型 html 或 markdown 为空则默认为 markdown
}
