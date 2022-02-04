<template>
  <div class="app-wrapper">
    <!-- 左侧 menu -->
    <!-- 这里的style是行内样式，目的是根据menuBg变量，改变menu背景色 -->
    <sidebar
      class="sidebar-container"
      :style="{ backgroundColor: variables.menuBg }"
    />
    <!-- 右侧容器 -->
    <div class="main-container">
      <div class="fixed-header">
        <!-- 顶部的 navbar -->
        <navbar />
      </div>
      <!-- 内容区 -->
      <app-main />
    </div>
  </div>
</template>

<script setup>
import Navbar from '@/layout/components/Navbar'
import Sidebar from '@/layout/components/Sidebar'
import AppMain from '@/layout/components/AppMain'
// 导入样式变量
import variables from '@/styles/variables.scss'
</script>

<style scoped lang="scss">
@import '~@/styles/mixin.scss';
@import '~@/styles/variables.scss';

.app-wrapper {
  // 调用clearfix这个mixin
  @include clearfix;
  // 针对body相对定位
  position: relative;
  height: 100%;
  width: 100%;
}

.fixed-header {
  // 相对于浏览器窗口定位
  position: fixed;
  top: 0;
  right: 0;
  z-index: 9;
  // Navbar的宽度是屏幕宽度 - Menu宽度
  width: calc(100% - #{$sideBarWidth});
}
</style>
