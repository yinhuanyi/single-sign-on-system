/**
 * @Author: Robby
 * @Date: 2022/3/4
 * @Filename: i18n.js
 * @Function:
 **/

import i18n from '@/i18n'

export function generateTitle(title) {
  return i18n.global.t('msg.route.' + title)
}
