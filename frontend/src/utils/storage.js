/**
 * @Author: Robby
 * @Date: 2022/1/27
 * @Filename: storage.js
 * @Function:
 **/

import Cookies from 'js-cookie'

// 存储数据
export const setItem = (key, value) => {
  // 将数组、对象类型的数据转化为 JSON 字符串进行存储
  if (typeof value === 'object') {
    value = JSON.stringify(value)
  }
  window.localStorage.setItem(key, value)
}

// 获取数据
export const getItem = (key) => {
  const data = window.localStorage.getItem(key)
  try {
    return JSON.parse(data)
  } catch (err) {
    return data
  }
}

// 删除数据
export const removeItem = (key) => {
  window.localStorage.removeItem(key)
}

// 删除所有数据
export const removeAllItem = (key) => {
  window.localStorage.clear()
}

export const removeAllCookie = () => {
  // 这个sessionid的名称是基于SSO的session名称指定的
  Cookies.remove('session_id')
}
