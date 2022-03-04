<template>
  <div class="tags-view-container">
    <router-link
      v-for="(tag, index) in $store.getters.tagsViewList"
      :key="index"
      class="tags-view-item"
      :class="isActive(tag) ? 'active' : ''"
      :to="{ path: tag.fullPath }"
      :style="{
        backgroundColor: isActive(tag) ? $store.getters.cssVar.menuBg : '',
        borderColor: isActive(tag) ? $store.getters.cssVar.menuBg : ''
      }"
      @contextmenu.prevent="openMenu($event, index)"
    >
      <!-- 标题 -->
      {{ tag.title }}
      <!-- 小叉号图标 -->
      <i
        v-show="!isActive(tag)"
        class="el-icon-close"
        @click.prevent.stop="onCloseClick(index)"
      />
    </router-link>

    <!-- 引入ContextMenu组件 -->
    <context-menu v-show="visible" :style="menuStyle" :index="selectIndex" />
  </div>
</template>

<script setup>
import { useStore } from 'vuex'
import { ref, watch } from 'vue'
import ContextMenu from '@/components/TagsView/ContextMenu'
import { useRoute } from 'vue-router'

const route = useRoute()
const store = useStore()

// 判断tag是否被选中
const isActive = (tag) => {
  return tag.path === route.path
}

// 关闭tag的点击事件
const onCloseClick = (index) => {
  store.commit('app/removeTagsView', {
    type: 'index',
    index
  })
}

const visible = ref(false)

// 控制contextMenu组件的left和top
const menuStyle = ref({
  left: 0,
  top: 0
})

const selectIndex = ref(0)

const openMenu = (e, index) => {
  const { x, y } = e
  menuStyle.value.left = x + 'px'
  menuStyle.value.top = y + 'px'
  selectIndex.value = index
  visible.value = true
}

const closeMenu = () => {
  visible.value = false
}

watch(visible, (val) => {
  if (val) {
    // 如果是contextMenu展示出来了，那么给页面添加一个点击事件
    document.body.addEventListener('click', closeMenu)
  } else {
    // 如果contextMenu关闭了，那么将这个点击事件关闭
    document.body.removeEventListener('click', closeMenu)
  }
})

</script>

<style scoped lang="scss">
.tags-view-container {
  height: 34px;
  width: 100%;
  background: #fff;
  border-bottom: 1px solid #d8dce5;
  box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.12), 0 0 3px 0 rgba(0, 0, 0, 0.04);
  .tags-view-item {
    display: inline-block;
    position: relative;
    cursor: pointer;
    height: 26px;
    line-height: 26px;
    border: 1px solid #d8dce5;
    color: #495060;
    background: #fff;
    padding: 0 8px;
    font-size: 12px;
    margin-left: 5px;
    margin-top: 4px;
    &:first-of-type {
      margin-left: 15px;
    }
    &:last-of-type {
      margin-right: 15px;
    }
    &.active {
      color: #fff;
      &::before {
        content: '';
        background: #fff;
        display: inline-block;
        width: 8px;
        height: 8px;
        border-radius: 50%;
        position: relative;
        margin-right: 4px;
      }
    }
    // close 按钮
    .el-icon-close {
      width: 16px;
      height: 16px;
      line-height: 10px;
      vertical-align: 2px;
      border-radius: 50%;
      text-align: center;
      transition: all 0.3s cubic-bezier(0.645, 0.045, 0.355, 1);
      transform-origin: 100% 50%;
      &:before {
        transform: scale(0.6);
        display: inline-block;
        vertical-align: -3px;
      }
      &:hover {
        background-color: #b4bccc;
        color: #fff;
      }
    }
  }
}
</style>
