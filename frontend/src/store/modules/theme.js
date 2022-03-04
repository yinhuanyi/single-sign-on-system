/**
 * @Author: Robby
 * @Date: 2022/3/4
 * @Filename: theme.js
 * @Function:
 **/
import variables from '@/styles/variables.scss'
import { MAIN_COLOR, DEFAULT_COLOR } from '@/constant'
import { getItem, setItem } from '@/utils/storage'

const state = {
  mainColor: getItem(MAIN_COLOR) || DEFAULT_COLOR,
  variables: variables
}

const mutations = {
  setMainColor(state, newColor) {
    state.mainColor = newColor
    state.variables.menuBg = newColor
    setItem(MAIN_COLOR, newColor)
  }
}

const actions = {}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
