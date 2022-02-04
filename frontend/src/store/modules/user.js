/**
 * @Author: Robby
 * @Date: 2021/12/7
 * @Filename: user.js
 * @Function:
 **/
import router from '@/router'
import { ACCESSTOKEN, REFRESHTOKEN } from '@/constant'
import { setItem, getItem } from '@/utils/storage'
import { login, ssoSessionSet } from '@/api/sys'

// import md5 from 'md5'

const state = {
  accessToken: getItem(ACCESSTOKEN) || '',
  refreshToken: getItem(REFRESHTOKEN) || ''
}

const mutations = {
  setAccessToken: (state, accessToken) => {
    state.accessToken = accessToken // 将token存储到vuex中
    setItem(ACCESSTOKEN, accessToken) // 将token存储LocalStorage中
  },
  setRefreshToken: (state, refreshToken) => {
    state.refreshToken = refreshToken
    setItem(REFRESHTOKEN, refreshToken)
  }
}

const actions = {
  login({ commit }, userInfo) {
    const { username, password } = userInfo
    // 异步请求都封装在Promise中
    return new Promise((resolve, reject) => {
      // login({ username: username.trim(), password: md5(password.trim()) })
      login({ username: username.trim(), password: password.trim() })
        .then((response) => {
          const { data } = response
          console.log(data)
          commit('setAccessToken', data.access_token)
          commit('setRefreshToken', data.refresh_token)
          router.push('/')
          resolve()
        })
        .catch((error) => {
          console.log('失败了')
          // 抛出异常
          reject(error)
        })
    })
  },

  ssoSessionSet({ commit }) {
    return new Promise((resolve, reject) => {
      ssoSessionSet()
        .then((response) => {
          // 这里是获取到SSO服务器的 【显示登录页面的数据】
          if (response.msg === 'Login') {
            // console.log(response)
            resolve()
          } else {
            reject(response)
          }
        })
        .catch((error) => {
          // 抛出异常
          reject(error)
        })
    })
  }
}

// 以一个独立的模块形式导出
export default {
  namespaced: true,
  state,
  mutations,
  actions
}
