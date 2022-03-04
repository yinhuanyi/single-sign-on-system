/**
 * @Author: Robby
 * @Date: 2021/12/7
 * @Filename: rules.js
 * @Function:
 **/

import i18n from '@/i18n'

export const validatorPassword = (rule, value, callback) => {
  if (value.length > 5) {
    callback() // 表示通过
  } else {
    callback(new Error(i18n.global.t('msg.login.passwordRule')))
  }
}
