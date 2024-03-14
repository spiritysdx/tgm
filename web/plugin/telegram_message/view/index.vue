<template>
  <div>
    <warning-bar title="需要提前在Telegram申请一个Bot的Token，务必注意Token不要泄露给任何人" />
    <div class="gva-form-box">
      <el-form
        ref="telegramMessageForm"
        label-position="right"
        label-width="80px"
        :model="form"
      > 
        <el-form-item label="token">
          <el-input v-model="form.token" />
        </el-form-item>
        <el-form-item label="chat_id">
          <el-input v-model="form.chat_id" />
        </el-form-item>
        <el-form-item label="消息内容">
          <el-input
            v-model="form.content"
            type="textarea"
          />
        </el-form-item>
        <el-form-item label="消息类型">
          <el-select v-model="form.message_type" placeholder="消息类型">
            <el-option label="Markdown" value="markdown"></el-option>
            <el-option label="HTML" value="html"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button @click="sendTelegramMessage">发送消息</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>

</template>

<script setup>
import WarningBar from '@/components/warningBar/warningBar.vue'
import { sendMessage } from '@/plugin/telegram_message/api/message.js'
import { ElMessage } from 'element-plus'
import { reactive, ref } from 'vue'

defineOptions({
  name: 'TelegramMessage',
})

const telegramMessageForm = ref(null)
const form = reactive({
  token: '',
  chat_id: '',
  content: '',
  message_type: 'markdown',
})
const sendTelegramMessage = async() => {
  const res = await sendMessage(form)
  if (res.code === 0) {
    ElMessage.success('发送成功,请查收')
  }
}
</script>

