<template>
  <el-dropdown
    class="international"
    trigger="click"
    @command="handleSetLanguage"
  >
    <div>
      <!--  el-tooltip是用于显示提醒的  -->
      <el-tooltip content="国际化" :effect="effect">
        <svg-icon icon="language"></svg-icon>
      </el-tooltip>
    </div>

    <template #dropdown>
      <el-dropdown-menu>
        <!-- 这里的command=zh，传递的是handleSetLanguage的参数 -->
        <el-dropdown-item :disabled="language === 'zh'" command="zh">
          中文
        </el-dropdown-item>
        <el-dropdown-item :disabled="language === 'en'" command="en">
          English
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script setup>
import { computed, defineProps } from 'vue'
import { useStore } from 'vuex'
import { useI18n } from 'vue-i18n'
import { ElMessage } from 'element-plus'

defineProps({
  effect: {
    type: String,
    default: 'dark',
    // 对传递的字符串进行校验
    validator: (value) => {
      return ['dark', 'light'].indexOf(value) !== -1
    }
  }
})

const store = useStore()
// 获取到i18n的实例
const i18n = useI18n()
const language = computed(() => store.getters.language)

const handleSetLanguage = (lang) => {
  // 设置i18n中locale的值
  i18n.locale.value = lang
  // 修改vuex的lang值
  store.commit('app/setLanguage', lang)
  // 给一个消息提示
  ElMessage({
    message: i18n.t('msg.toast.switchLangSuccess'),
    type: 'success',
    duration: 2 * 1000
  })
}

</script>

<style scoped lang="scss"></style>
