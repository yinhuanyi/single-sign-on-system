/**
 * @Author: Robby
 * @Date: 2022/3/2
 * @Filename: index.js
 * @Function:
 **/

import { createI18n } from 'vue-i18n'
import zhLocale from './lang/zh'
import enLocale from './lang/en'
import store from '@/store'
import { getItem } from '@/utils/storage'
import { LANG } from '@/constant'

function getLanguageFromVuex() {
  return store && store.getters && store.getters.language
}

function getLanguageFromLS() {
  return getItem(LANG)
}

const messages = {
  en: {
    msg: {
      ...enLocale
    }
  },
  zh: {
    msg: {
      ...zhLocale
    }
  }
}

const i18n = createI18n({
  legacy: false, // vue3使用composition api
  globalInjection: true, // 全局使用t函数
  locale: getLanguageFromVuex() || getLanguageFromLS() || 'zh',
  messages
})

export default i18n
