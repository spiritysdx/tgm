package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/telegram_bot/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/telegram_bot/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Telegram_botApi struct{}

// SendMessage
// @Tags Telegram_bot
// @Summary 指定chat_id发送消息
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param    token chat_id content  "发送消息必须的参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /telegram_bot/routerName [post]
func (p *Telegram_botApi) SendMessage(c *gin.Context) {
	var plug model.TelegramChat
	bindErr := c.ShouldBindJSON(&plug)
	if bindErr != nil {
		response.FailWithMessage(bindErr.Error(), c)
		return
	}
	res, err := service.ServiceGroupApp.SendTgMessage(plug.Token, plug.ChatId, plug.Content, plug.MessageType)
	if err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Error(err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithDetailed(res, "发送成功", c)
	}
}
