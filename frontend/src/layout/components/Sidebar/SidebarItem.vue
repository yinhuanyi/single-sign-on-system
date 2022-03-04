<!--二级菜单-->
<template>
  <!-- 当menu数据中有children的时候，展示el-submenu -->
  <el-sub-menu v-if="route.children.length > 0" :index="route.path">
    <template #title>
      <menu-item :title="route.meta.title" :icon="route.meta.icon"></menu-item>
    </template>
    <!-- 这里是自己渲染自己，就是渲染子菜单 -->
    <sidebar-item
      v-for="item in route.children"
      :key="item.path"
      :route="item"
    ></sidebar-item>
  </el-sub-menu>

  <!--如果menu数据中没有children的时候，展示el-menu-item-->
  <el-menu-item v-else :index="route.path">
    <menu-item :title="route.meta.title" :icon="route.meta.icon"></menu-item>
  </el-menu-item>
</template>

<script setup>
import { defineProps } from 'vue'
import MenuItem from '@/layout/components/Sidebar/MenuItem'

defineProps({
  route: {
    type: Object,
    required: true
  }
})
</script>

<style scoped lang="scss"></style>
