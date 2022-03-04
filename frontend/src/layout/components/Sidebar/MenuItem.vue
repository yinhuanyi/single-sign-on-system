<!--有children的二级菜单的样式组件，用于显示title和icon-->
<template>
  <!-- elementui的icon：如果icon字符串以el-icon开头 -->
  <i v-if="icon.includes('el-icon')" class="sub-el-icon" :class="icon"></i>
  <!-- 自定义的icon -->
  <svg-icon v-else :icon="icon"></svg-icon>
  <!-- 文本 -->
  <span class='title'> {{ generateTitle(title) }}</span>
</template>

<script setup>
import { defineProps } from 'vue'
import SvgIcon from '@/components/SvgIcon'
import { generateTitle } from '@/utils/i18n'

defineProps({
  title: {
    type: String,
    required: true
  },
  icon: {
    type: String,
    required: true
  }
})

</script>

<style scoped lang="scss">
.title{
  margin-left: 10px;
  font-weight: 600;
  //line-height: 50px;
  font-size: 15px;
  //white-space: nowrap;
}
</style>
