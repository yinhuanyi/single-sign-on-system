/**
 * @Author: Robby
 * @Date: 2022/1/25
 * @Filename: sys.js
 * @Function:
 **/

import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/login',
    method: 'POST',
    data: data
  })
}

// 发送请求到业务子系统
export function ssoSessionSet() {
  return request({
    url: '/sso_login',
    method: 'GET'
  })
}

// 请求用户信息
export function getUserInfo() {
  return request({
    url: '/user/profile',
    method: 'GET'
  })
}
