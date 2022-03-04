<template>
  <el-menu
    :collapse="!$store.getters.sidebarOpened"
    :unique-opened="true"
    :background-color="$store.getters.cssVar.menuBg"
    :text-color="$store.getters.cssVar.menuText"
    :active-text-color="$store.getters.cssVar.menuActiveText"
    :router="true"
    :default-active="activeMenu"
  >
    <sidebar-item
      v-for="item in routes"
      :key="item.path"
      :route="item"
    ></sidebar-item>
  </el-menu>
</template>

<script setup>
import { computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { filterRoutes, generateMenus } from '@/utils/route'
import SidebarItem from '@/layout/components/Sidebar/SidebarItem'

// 获取所有路由配置
const router = useRouter()
const route = useRoute()
// 通过filterRoutes过滤掉脱离层级的子路由，通过
const routes = computed(() => {
  // 通过filterRoutes过滤掉脱离层级的子路由
  const filterRoutesResult = filterRoutes(router.getRoutes())
  // 通过generateMenus生成想要的路由信息
  return generateMenus(filterRoutesResult)
})

// console.log(JSON.stringify(routes.value))

// 获取当前路由的path路径
const activeMenu = computed(() => {
  const { path } = route
  return path
})

</script>

<style scoped lang="scss"></style>
