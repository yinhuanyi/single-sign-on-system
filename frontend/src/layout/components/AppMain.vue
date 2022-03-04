<template>
  <div class="app-main">
    <!--  下面是一个固定写法  -->
    <router-view v-slot="{ Component, route }">
      <transition name="fade-transform" mode="out-in">
        <!-- 这里是动态的渲染Component -->
        <keep-alive>
          <component :is="Component" :key="route.path" />
        </keep-alive>
      </transition>
    </router-view>
  </div>
</template>

<script setup>
import { watchSwitchLang } from '@/components/HeaderSearch/watchLanguage'
import { generateTitle } from '@/utils/i18n'
import { isTags } from '@/utils/tags'
import { watch } from 'vue'
import { useRoute } from 'vue-router'
import { useStore } from 'vuex'

const route = useRoute()
const store = useStore()

// 基于 生成title
const getTitle = (route) => {
  let title = ''
  // 如果路由中不存在meta，那么去路由最后一个字符串作为title
  if (!route.meta) {
    const pathArr = route.path.split('/')
    title = pathArr[pathArr.length - 1]
  } else {
    title = generateTitle(route.meta.title)
  }
  return title
}

// 监听route，获取to、from两个属性的值，将路由的信息保存到vuex和LocalStorage的tagsViewList中
watch(
  route,
  (to, from) => {
    // 如果路径在白名单里面，就不需要将路由path保存到tags中
    if (isTags(to.path)) {
      return
    }
    // 如果路由的path不在白名单，
    const { fullPath, meta, name, params, path, query } = to
    store.commit('app/addTagsView', {
      fullPath,
      meta,
      name,
      params,
      path,
      query,
      title: getTitle(route)
    })
  },
  {
    // 当组件创建的时候，watch就会执行一次
    immediate: true
  }
)

// 这是一个监控语言变化的回调方法，当语言切换的时候，函数会执行
watchSwitchLang(() => {
  store.getters.tagsViewList.forEach((route, index) => {
    store.commit('app/changeTagsView', {
      index,
      tag: {
        ...route,
        title: getTitle(route)
      }
    })
  })
})
</script>

<style scoped lang="scss">
.app-main {
  // 可视区域的最小高度 100vh，这里也减了43px
  min-height: calc(100vh - 50px - 43px);
  width: 100%;
  position: relative;
  overflow: hidden;
  // 这里的padding-top加了43px，是tagsView的高度
  padding: 104px 20px 20px 20px;
  // border-box盒模型width和height包含border和padding
  box-sizing: border-box;
}
</style>
