# tgm

练手之作，如有错漏望告知，谢谢

仓库: https://github.com/spiritysdx/tgm

开发环境：GVA v2.6.1

当前插件版本：v0.0.4

- [x] 根据Bot的token，用户的chat_id，发送txt格式的文本内容
- [x] 支持html格式和markdown格式的文本内容
- [x] 支持使用单一/多个Bot的token(英文逗号分隔后输入)，多个token避免触发Bot的消息发送频率的风控，确保消息一定发出
- [x] 制作接口，通过Bot获取所在频道是否存在对应user_id的用户

具体部署说明见 server/plugin/telegram_bot 下的 README.md 说明