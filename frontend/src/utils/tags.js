/**
 * @Author: Robby
 * @Date: 2022/3/4
 * @Filename: tags.js
 * @Function:
 **/

const whiteList = ['/login', '/import', '/404', '/401']

// 判断path 是否需要被缓存
export function isTags(path) {
  return whiteList.includes(path)
}
