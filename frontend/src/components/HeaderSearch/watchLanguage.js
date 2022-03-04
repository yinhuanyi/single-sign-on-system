/**
 * @Author: Robby
 * @Date: 2021/12/13
 * @Filename: watchLanguage.js
 * @Function:
 **/
import { watch } from 'vue'
import store from '@/store'

// cbs是函数列表，当language变化，会执行cbs中的函数
export function watchSwitchLang(...cbs) {
  watch(
    // 监控vuex中language的变化
    () => store.getters.language,
    //
    () => {
      cbs.forEach((cb) => cb())
    }
  )
}
