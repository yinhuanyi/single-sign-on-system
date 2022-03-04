import { createStore } from 'vuex'
import getters from './getters'
import user from '@/store/modules/user'
import app from '@/store/modules/app'
import theme from '@/store/modules/theme'

export default createStore({
  modules: {
    user,
    app,
    theme
  },
  getters
})
