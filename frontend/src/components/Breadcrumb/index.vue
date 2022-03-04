<template>
  <el-breadcrumb class="breadcrumb" separator="/">
    <transition-group name="breadcrumb">
      <el-breadcrumb-item
      v-for="(item, index) in breadcrumbData"
      :key="item.path"
    >
      <!-- 不可点击 -->
      <span class="no-redirect" v-if="index === breadcrumbData.length - 1">
        {{ generateTitle(item.meta.title) }}
      </span>
      <!--  可点击 -->
      <span class="redirect" v-else @click="onLinkClick(item)">
        {{ generateTitle(item.meta.title) }}
      </span>
    </el-breadcrumb-item>
    </transition-group>
  </el-breadcrumb>
</template>

<script setup>
import { watch, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { generateTitle } from '@/utils/i18n'

const route = useRoute()
const router = useRouter()
const store = useStore()

const linkHoverColor = ref(store.getters.cssVar.menuBg)

const onLinkClick = (item) => {
  router.push(item.path)
}

const breadcrumbData = ref([])
const getBreadcrumbData = () => {
  // 过滤的目的是，只需要有title的路由路径
  breadcrumbData.value = route.matched.filter(
    (item) => item.meta && item.meta.title
  )
}

// 监听route变化
watch(
  route,
  () => {
    getBreadcrumbData()
  },
  { immediate: true }
)
</script>

<style scoped lang="scss">
.breadcrumb {
  display: inline-block;
  font-size: 14px;
  line-height: 50px;
  margin-left: 8px;

  .redirect {
    color: #666;
    font-weight: 600;
    cursor: pointer;
  }

  .redirect:hover {
    // v-bind，可以使用js中的变量
    color: v-bind(linkHoverColor);
  }

  ::v-deep .no-redirect {
    color: #97a8be;
    cursor: text;
  }
}
</style>
