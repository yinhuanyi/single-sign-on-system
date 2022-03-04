<template>
  <div class="header-search" :class="{ show: isShow }">
    <!-- click.stop的stop是禁止点击事件冒泡，如果外面元素还有监听click事件，不会被触发 -->
    <span @click.stop="onShowClick">
      <svg-icon class-name="search-icon" icon="search" />
    </span>
    <!--    <div class="search-icon" @click.stop="onShowClick">as</div>-->
    <el-select
      ref="headerSearchSelectRef"
      class="header-search-select"
      v-model="search"
      :filterable="true"
      default-first-option
      remote
      :remote-method="querySearch"
      placeholder="search"
      @change="onSelectChange"
    >
      <el-option
        v-for="option in searchOptions"
        :key="option.item.path"
        :label="option.item.title.join(' > ')"
        :value="option.item"
      />
    </el-select>
  </div>
</template>

<script setup>
import { watchSwitchLang } from '@/components/HeaderSearch/watchLanguage'
import { generateRoutes } from '@/components/HeaderSearch/fuseData'
import Fuse from 'fuse.js'
import { ref, computed, watch } from 'vue'
import { filterRoutes } from '@/utils/route'
import { useRouter } from 'vue-router'
import SvgIcon from '@/components/SvgIcon'

// 创建搜索数据源
const router = useRouter()
let searchPool = computed(() => {
  const filterResults = filterRoutes(router.getRoutes())
  return generateRoutes(filterResults)
})

// 基于searchPool数据源，配置fuse模糊搜索
let fuse
const initFuse = (searchPool) => {
  fuse = new Fuse(searchPool, {
    // 是否按优先级进行排序
    shouldSort: true,
    // 匹配长度超过这个值的才会被认为是匹配的
    minMatchCharLength: 1,
    // 将被搜索的键列表。 这支持嵌套路径、加权搜索、在字符串和对象数组中搜索。
    // name：搜索的键
    // weight：对应的权重
    keys: [
      {
        name: 'title',
        weight: 0.7
      },
      {
        name: 'path',
        weight: 0.3
      }
    ]
  })
}
initFuse(searchPool.value)

// 控制search输入框展示的变量
const isShow = ref(false)

// el-select 实例
const headerSearchSelectRef = ref(null)

const onShowClick = () => {
  isShow.value = !isShow.value
  // 出现了搜索框，让搜索框聚焦
  headerSearchSelectRef.value.focus()
}

const search = ref('')

// 存储fuse模糊匹配到的结果
const searchOptions = ref([])

// 搜索方法
const querySearch = (query) => {
  if (query !== '') {
    searchOptions.value = fuse.search(query)
  } else {
    searchOptions.value = []
  }
}

// 选中option的回调
const onSelectChange = (val) => {
  router.push(val.path)
}

const onClose = () => {
  // 让select选择器失去焦点
  headerSearchSelectRef.value.blur()
  // 让select隐藏
  isShow.value = false
  // 清空匹配结果
  searchOptions.value = []
}

// 监听isShow的变化，当isShow为true的时候
watch(isShow, (val) => {
  if (val) {
    // 让select选择器获取焦点
    headerSearchSelectRef.value.focus()
    // 当前为ture，那么给document添加一个点击事件，点击了就会调用onClose方法
    document.body.addEventListener('click', onClose)
  } else {
    document.body.removeEventListener('click', onClose)
  }
})

// 调用这个函数会被压栈，内部逻辑的执行当vuex的language变化的时候才会执行
watchSwitchLang(() => {
  searchPool = computed(() => {
    const filterResults = filterRoutes(router.getRoutes())
    return generateRoutes(filterResults)
  })
  initFuse(searchPool.value)
})
</script>

<style scoped lang="scss">
.header-search {
  font-size: 0 !important;
  // 这里由于把样式传递到组件内部了，那么操作子组件的样式，需要使用::v-deep
  ::v-deep .search-icon {
    cursor: pointer;
    font-size: 18px;
    vertical-align: middle;
  }
  .header-search-select {
    font-size: 18px;
    transition: width 0.2s;
    width: 0;
    overflow: hidden;
    background: transparent;
    border-radius: 0;
    display: inline-block;
    vertical-align: middle;

    ::v-deep .el-input__inner {
      border-radius: 0;
      border: 0;
      padding-left: 0;
      padding-right: 0;
      box-shadow: none !important;
      border-bottom: 1px solid #d9d9d9;
      vertical-align: middle;
    }
  }
  &.show {
    .header-search-select {
      width: 210px;
      margin-left: 10px;
    }
  }
}
</style>
