import router from './router'
import store from './store'
// import md5 from 'md5'

// 白名单
const whiteList = ['/login']
/**
 * 路由前置守卫
 * to: 要到哪里去
 * from：你从哪里来
 * next：是否要去
 */
router.beforeEach(async (to, from, next) => {
  if (store.getters.accessToken) {
    if (to.path === '/login') {
      next('/') // 如果用户访问的是/login，那么跳转到首页
    } else {
      next() // 如果访问的不是/login，那么直接跳转即可
    }
  } else {
    // 不存在token，如果访问白名单路由，那么直接跳转即可
    if (whiteList.indexOf(to.path) > -1) {
      // 这个不能少，如果只是next('/login')，那么会进入到重定向无限循环中
      await store.dispatch('user/ssoSessionSet')
      next()
    } else {
      await store.dispatch('user/ssoSessionSet')
      next('/login')
    }
  }
})
