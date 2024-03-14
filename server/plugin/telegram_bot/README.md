## Telegram_Bot插件

开发者：spiritlhl

版本：v0.0.1

### 使用步骤

#### 1

查看 ```server/initialize/plugin.go``` 文件中是否已注册插件，如若未注册，在

```
func InstallPlugin
```

函数中插入

```
    PluginInit(PrivateGroup, telegram_bot.CreateTelegram_botPlug())
```

### 2. 配置说明

入参结构说明

```
type TelegramChat struct {
	Token string `json:"token"` // Telegram的Bot的Token
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

### 后台设置

![image](https://github.com/spiritysdx/tgm/assets/97792170/1c1468a9-a3dd-45ae-94fc-dd05c60f0eff)

路径：```/plugin/telegram_message/view/index.vue```

![image](https://github.com/spiritysdx/tgm/assets/97792170/7b40a5ec-78a5-47b6-9cfa-d4a219578369)

路径：```/telegram_bot/sendMessage```
