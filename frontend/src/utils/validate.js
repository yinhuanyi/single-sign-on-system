/**
 * @Author: Robby
 * @Date: 2022/1/25
 * @Filename: validate.js
 * @Function:
 **/

// 判断path是否是外部资源
export function isExternal(path) {
  // 判断path是否以https开头
  return /^(https?:|mailto:|tel:)/.test(path)
}
