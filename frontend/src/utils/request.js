/**
 * @Author: Robby
 * @Date: 2021/12/7
 * @Filename: request.js
 * @Function:
 **/
import axios from 'axios'
import { ElMessage } from 'element-plus'
import store from '@/store'
// import { isCheckTimeout } from '@/utils/auth'

const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API,
  withCredentials: true,
  timeout: 500000000
  // headers: { icode: '18B81CA2C36DE48F' }
})

// 请求拦截器
service.interceptors.request.use(
  (config) => {
    if (store.getters.accessToken) {
      // 判断用户token是否超时
      // if (isCheckTimeout()) {
      //   // 用户登出
      //   store.dispatch('user/logout')
      //   // 抛出异常
      //   return Promise.reject(new Error('token 失效'))
      // }

      // 获取用户token，设置Authorization请求头
      config.headers.Authorization = `Bearer ${store.getters.accessToken}`
      config.headers['Refresh-Token'] = store.getters.refreshToken
      // 添加Accept-Language请求头
      config.headers['Accept-Language'] = store.getters.language
    }

    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
service.interceptors.response.use(

  // 请求成功
  (response) => {
    const res = response.data
    // console.log(res)
    // 如果success字段是true，那么返回结果
    if (res.success) {
      return res
    } else {
      ElMessage({
        message: res.msg || '请求失败',
        type: 'error',
        duration: 5 * 1000
      })
      return Promise.reject(new Error(res.message || '请求失败'))
    }
  },

  // 请求失败
  (error) => {
    console.log(error)
    // 判断服务器端返回的状态码，根据状态码===401，让用户登出
    if (
      error.response &&
      error.response.data &&
      error.response.data.code === 401
    ) {
      store.dispatch('user/logout')
    }

    ElMessage({
      message: error.message || '请求失败',
      type: 'error',
      duration: 5 * 1000
    })

    return Promise.reject(error)
  }
)

export default service
