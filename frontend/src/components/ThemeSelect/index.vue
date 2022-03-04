<template>
  <!-- v-bind="$attrs" 当父组件传递非props属性的时候，会直接绑定到这里 -->
  <el-dropdown
    v-bind="$attrs"
    trigger="click"
    class="theme"
    @command="handleSetTheme"
  >
    <div>
      <el-tooltip :content="$t('msg.navBar.themeChange')">
        <svg-icon icon="change-theme"></svg-icon>
      </el-tooltip>
    </div>

    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item command="color">
          {{ $t('msg.theme.themeColorChange') }}
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>

  <!-- 展示弹出层 -->
  <div>
    <!-- @close是监听子组件的close事件，$event是获取到事件传递的值 -->
    <select-color v-model:visible="selectColorVisible" />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import SelectColor from '@/components/ThemeSelect/components/SelectColor'

const handleSetTheme = (command) => {
  selectColorVisible.value = true
}

// 设置dailog初始状态为隐藏
const selectColorVisible = ref(false)

</script>

<style scoped lang="scss">

</style>
