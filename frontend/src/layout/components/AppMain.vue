<template>
  <div class="app-main">AppMain</div>
</template>

<script setup></script>

<style scoped lang="scss">
.app-main {
  // 浏览器可视区域的高度 100vh
  min-height: calc(100vh - 50px);
  width: 100%;
  position: relative;
  overflow: hidden;
  padding: 61px 20px 20px 20px;
  // border-box盒模型width和height包含border和padding
  box-sizing: border-box;
}
</style>
