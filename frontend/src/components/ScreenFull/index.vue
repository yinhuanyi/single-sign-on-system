<template>
  <div @click="onToggle">
    <svg-icon
      :icon="isFullScreen ? 'exit-fullscreen' : 'fullscreen'"
    ></svg-icon>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import fullScreen from 'screenfull'
const isFullScreen = ref(false)

const onToggle = () => {
  fullScreen.toggle()
}

const change = () => {
  isFullScreen.value = fullScreen.isFullscreen
}

// 这里需要监听screenfull的变化
onMounted(() => {
  fullScreen.on('change', change)
})

onUnmounted(() => {
  fullScreen.on('change', change)
})
</script>

<style scoped lang="scss"></style>
