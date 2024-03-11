## Telegram_Bot插件

开发者：spiritlhl

版本：v0.0.1

### 使用步骤

#### 1. 前往GVA主程序下的initialize/router.go 在 Routers 方法最末尾按照你需要的及安全模式添加本插件

```
    PluginInit(PrivateGroup, telegram_bot.CreateTelegram_botPlug())
```

### 2. 配置说明

入参结构说明

```
type TelegramChat struct {
    Token string // BOT的token
	ChatId  string `json:"chat_id"` // 发送的目标ID
	Content string `json:"content"` // 发送的内容
}
```

### 3. 可直接调用的接口

发送消息的接口： /telegram_bot/sendMessage [post]

入参：
```
type TelegramChat struct {
	Token string `json:"token"` // Telegram的Bot的Token
	ChatId  string `json:"chat_id"` // 发送的目标ID
	Content string `json:"content"` // 发送的内容
}
```