<!-- 对图标组件的封装 -->
<template>
  <!-- 展示外部图标 -->
  <div
    v-if="isExternal"
    :style="styleExternalIcon"
    class="svg-external-icon svg-icon"
    :class="className"
  ></div>

  <!-- 展示内部图标 -->
  <svg v-else class="svg-icon" :class="className" aria-hidden="true">
    <!--  iconName是基于icon字符串计算的  -->
    <use :xlink:href="iconName" />
  </svg>
</template>

<script setup>
// 这里是定义子组件的props
import { defineProps, computed } from 'vue'
// 这里给函数取了一个别名
import { isExternal as external } from '@/utils/validate'

// 接收父组件的props属性
const props = defineProps({
  // 图标名称
  icon: {
    type: String,
    required: true
  },
  // 图标类名称
  className: {
    type: String,
    default: ''
  }
})

// 判断当前图标是否是外部图标
const isExternal = computed(() => external(props.icon))

// 外部图标样式
const styleExternalIcon = computed(() => ({
  mask: `url(${props.icon}) no-repeat 50% 50%`,
  '-webkit-mask': `url(${props.icon}) no-repeat 50% 50%`
}))

// 内部图标：给所有的图标添加前缀, 因为在vue.config.js文件中，导入的时候加了icon前缀
const iconName = computed(() => `#icon-${props.icon}`)
</script>

<style scoped lang="scss">
.svg-icon {
  width: 1em;
  height: 1em;
  vertical-align: -0.15em;
  fill: currentColor;
  overflow: hidden;
}
.svg-external-icon {
  background-color: currentColor;
  mask-size: cover !important;
  display: inline-block;
}
</style>
