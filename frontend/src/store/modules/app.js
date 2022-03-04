/**
 * @Author: Robby
 * @Date: 2022/3/2
 * @Filename: app.js
 * @Function: 存储侧边栏状态
 **/
import { getItem, setItem } from '@/utils/storage'
import { LANG, TAGS_VIEW } from '@/constant'

const state = {
  sidebarOpened: true,
  language: getItem(LANG) || 'zh',
  tagsViewList: getItem(TAGS_VIEW) || []
}

const mutations = {
  triggerSidebarOpened: (state) => {
    state.sidebarOpened = !state.sidebarOpened
  },
  setLanguage: (state, lang) => {
    setItem(LANG, lang)
    state.language = lang
  },
  addTagsView: (state, tag) => {
    // 如果tag已经存在，那么不添加
    const isFind = state.tagsViewList.find((item) => {
      return item.path === tag.path
    })
    // 如果不存在，那么存储到vuex和LocalStorage中
    if (!isFind) {
      state.tagsViewList.push(tag)
      setItem(TAGS_VIEW, state.tagsViewList)
    }
  },
  changeTagsView: (state, { index, tag }) => {
    state.tagsViewList[index] = tag
    setItem(TAGS_VIEW, state.tagsViewList)
  },
  // 当payload.type为other 表示删除其他，right表示删除右侧，index表示删除当前
  removeTagsView: (state, payload) => {
    if (payload.type === 'index') {
      // 删除数组中index所在的元素
      state.tagsViewList.splice(payload.index, 1)
    } else if (payload.type === 'other') {
      // 删除数组中index所在元素之后的元素，其中payload.index + 1删除的起始索引，state.tagsViewList.length - payload.index + 1是删除的个数
      state.tagsViewList.splice(
        payload.index + 1,
        state.tagsViewList.length - payload.index + 1
      )
      // 删除数组中index所在元素之前的元素
      state.tagsViewList.splice(0, payload.index)
    } else if (payload.type === 'right') {
      state.tagsViewList.splice(
        payload.index + 1,
        state.tagsViewList.length - payload.index + 1
      )
    }
    // 最后将LocalStorage中的tagsView更新
    setItem(TAGS_VIEW, state.tagsViewList)
  }
}

export default {
  namespaced: true,
  state,
  mutations
}
