/**
 * @Author: Robby
 * @Date: 2021/12/7
 * @Filename: rules.js
 * @Function:
 **/

// import i18n from '@/i18n'

export const validatorPassword = (rule, value, callback) => {
  if (value.length > 5) {
    callback() // 表示通过
  } else {
    callback(new Error('用户名长度必须大于5位数'))
  }
}
