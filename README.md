# tgm

练手之作，如有错漏望告知，谢谢

开发环境：GVA v2.6.1

当前插件版本：v0.0.3

- [x] 根据Bot的token，用户的chat_id，发送txt格式的文本内容
- [x] 支持html格式和markdown格式的文本内容
- [x] 支持使用单一/多个Bot的token(英文逗号分隔后输入)，多个token避免触发Bot的消息发送频率的风控，确保消息一定发出
- [ ] 制作接口，通过Bot获取所在频道是否存在对应tgid的用户
- [ ] 制作接口，通过Bot查询某个tgid的用户注册的时长，作用：校验用户是否频繁注销再注册刷权限
- [ ] 制作接口，通过Bot查询某个tgid的用户是否在行业黑名单数据库上，作用：屏蔽垃圾用户
