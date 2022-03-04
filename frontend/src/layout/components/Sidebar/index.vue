<template>
  <div>
    <div class="logo-container">
      <el-avatar
        size="44"
        shape="square"
        src="https://element-plus.gitee.io/images/element-plus-logo.svg"
      ></el-avatar>
      <!-- 当侧边栏展开的时候显示系统title -->
      <h1 class="logo-title" v-if="$store.getters.sidebarOpened">
        后台管理系统
      </h1>
    </div>
    <!--   当菜单过长，el-scrollbar会出现滚动条 -->
    <el-scrollbar>
      <sidebar-menu></sidebar-menu>
    </el-scrollbar>
  </div>
</template>

<script setup>
import SidebarMenu from '@/layout/components/Sidebar/SidebarMenu'
</script>

<style scoped lang="scss">
.logo-container {
  height: 44px;
  padding: 10px 0 22px 0;
  display: flex;
  align-items: center;
  justify-content: center;
  .logo-title {
    margin-left: 10px;
    color: #fff;
    font-weight: 600;
    line-height: 50px;
    font-size: 18px;
    white-space: nowrap;
  }
  .el-avatar {
    --el-avatar-background-color: none;
  }
}

</style>
