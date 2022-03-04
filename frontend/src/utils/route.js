/**
 * @Author: Robby
 * @Date: 2022/3/2
 * @Filename: route.js
 * @Function: 除脱离层级的路由配置，再生成符合menu的路由数据
 **/

import path from 'path'

// 获取子集路由
function getChildrenRoutes(routes) {
  const result = []
  routes.forEach((route) => {
    if (route.children && route.children.length > 0) {
      // 如果路由有children，全部添加到result里面
      result.push(...route.children)
    }
  })
  return result
}

// 过滤路由配置数据，删除脱离路由层级的路由配置
export function filterRoutes(routes) {
  // 获取所有的子路由
  const childrenRoutes = getChildrenRoutes(routes)
  // 将子路由从routes中过滤掉
  return routes.filter((route) => {
    return !childrenRoutes.find((childrenRoute) => {
      return childrenRoute.path === route.path
    })
  })
}

// 判断data是否为空
function isNull(data) {
  if (!data) return true
  if (JSON.stringify(data) === '{}') return true
  if (JSON.stringify(data) === '[]') return true
}

// 根据过滤的路由配置数据，创建menu规则的数据
export function generateMenus(routes, basePath = '') {
  const result = []
  routes.forEach((item) => {
    // 如果route不存在children 且不存在meta，说明是子路由，直接丢弃
    if (isNull(item.children) && isNull(item.meta)) return
    // 如果route存在children，但是不存在meta，说明是嵌套路由，需要对children再次遍历，
    if (!isNull(item.children) && isNull(item.meta)) {
      result.push(...generateMenus(item.children)) // Todo: 这里是递归调用
    }
    // item如果不存在children，但是存在meta
    // 将basePath与路由的path合并
    const routePath = path.resolve(basePath, item.path)
    // Todo：路由分离之后，可能存在同名的父路由情况，
    let route = result.find((item) => item.path === routePath)
    // 如果是空的，说明当前路由未加入到result
    if (!route) {
      route = {
        ...item,
        path: routePath,
        children: []
      }
      if (route.meta.icon && route.meta.title) {
        result.push(route)
      }
    }
    // 如果存在children和meta
    if (!isNull(item.children)) {
      route.children.push(...generateMenus(item.children, route.path))
    }
  })
  return result
}
