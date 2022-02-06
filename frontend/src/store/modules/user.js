/**
 * @Author: Robby
 * @Date: 2021/12/7
 * @Filename: user.js
 * @Function:
 **/
import router from '@/router'
import { ACCESSTOKEN, REFRESHTOKEN } from '@/constant'
import { setItem, getItem, removeAllItem, removeAllCookie } from '@/utils/storage'
import { getUserInfo, login, ssoSessionSet } from '@/api/sys'

const state = {
  accessToken: getItem(ACCESSTOKEN) || '',
  refreshToken: getItem(REFRESHTOKEN) || '',
  userInfo: {}
}

const mutations = {
  setAccessToken: (state, accessToken) => {
    state.accessToken = accessToken // 将token存储到vuex中
    setItem(ACCESSTOKEN, accessToken) // 将token存储LocalStorage中
  },
  setRefreshToken: (state, refreshToken) => {
    state.refreshToken = refreshToken
    setItem(REFRESHTOKEN, refreshToken)
  },
  setUserInfo: (state, userInfo) => {
    state.userInfo = userInfo // 将用户信息存储到state中
  }
}

const actions = {

  // 获取session_id
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
  },

  // 获取token
  login({ commit }, userInfo) {
    const { username, password } = userInfo
    // 异步请求都封装在Promise中
    return new Promise((resolve, reject) => {
      // login({ username: username.trim(), password: md5(password.trim()) })
      login({ username: username.trim(), password: password.trim() })
        .then((response) => {
          const { data } = response
          // console.log(data)
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

  // 获取用户信息
  async getUserInfo({ commit }) {
    const res = await getUserInfo()
    // 如果状态码是20302，那么需要更新accessToken和refreshToken，然后重新发起请求
    if (res.code === 20302) {
      commit('setAccessToken', res.data.access_token)
      commit('setRefreshToken', res.data.refresh_token)
      router.push('/')
      // 这里需要return，否则会运行下面的commit
      return res

    // 如果状态码是20303，那么需要清空accessToken、refreshToken、userinfo，然后重新跳转到/路由
    } else if (res.code === 20303) {
      commit('setAccessToken', '')
      commit('setRefreshToken', '')
      commit('setUserInfo', {})
      removeAllCookie()
      removeAllItem()
      router.push('/')
      return res
    }
    commit('setUserInfo', res.data)
    return res
  }
}

// 以一个独立的模块形式导出
export default {
  namespaced: true,
  state,
  mutations,
  actions
}
