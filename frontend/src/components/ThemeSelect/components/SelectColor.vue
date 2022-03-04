<template>
  <!-- :model-value属性是用于控制弹出层关闭和显示 -->
  <el-dialog
    :title="$t('msg.universal.title')"
    :model-value="visible"
    @close="closed"
    width="22%"
  >
    <!-- 弹出层的内容 -->
    <div class="content">
      <p class="title">{{ $t('msg.theme.themeColorChange') }}</p>
      <el-color-picker
        v-model="mColor"
        :predefine="predefineColors"
      ></el-color-picker>
    </div>
    <!-- 弹出层的底部 -->
    <template #footer>
      <el-button @click="closed">{{ $t('msg.universal.cancel') }}</el-button>
      <el-button type="primary" @click="confirm">{{ $t('msg.universal.confirm') }}</el-button>
    </template>
  </el-dialog>
</template>

<script setup>
import { useStore } from 'vuex'
import { generateNewStyle, writeNewStyle } from '@/utils/theme'
import { defineProps, ref, defineEmits } from 'vue'

defineProps({
  visible: {
    type: Boolean,
    required: true
  }
})

const store = useStore()

// 定义一个事件，并且指定参数名称为visible
const emits = defineEmits(['update:visible'])

// 当用户点击弹出层的取消按钮时调用
const closed = () => {
  // 子组件抛出update事件，并且指定给visible为false
  emits('update:visible', false)
}

// 当用户点击弹出层的确定按钮时调用
const confirm = async () => {
  const newStyleText = await generateNewStyle(mColor.value)
  writeNewStyle(newStyleText)
  store.commit('theme/setMainColor', mColor.value)
  closed()
}

// 预定义色值
const predefineColors = [
  '#ff4500',
  '#ff8c00',
  '#ffd700',
  '#90ee90',
  '#00ced1',
  '#1e90ff',
  '#c71585',
  '#000000',
  'rgba(255, 69, 0, 0.68)',
  'rgb(255, 120, 0)',
  'hsv(51, 100, 98)',
  'hsva(120, 40, 94, 0.5)',
  'hsl(181, 100%, 37%)',
  'hsla(209, 100%, 56%, 0.73)',
  '#c7158577'
]
// 默认色值
const mColor = ref('#000000')

</script>

<style scoped lang="scss">
.content {
  text-align: center;
  .title {
    margin-bottom: 12px;
  }
}
</style>
