import store from './store'
import { createApp } from 'vue'
import App from './App.vue'
import installElementPlus from './plugins/element'
import router from './router'
import installIcons from '@/icons'
import '@/styles/index.scss'
import '@/permission'
import i18n from '@/i18n'

const app = createApp(App)
installElementPlus(app)
installIcons(app)

app.use(store).use(router).use(i18n).mount('#app')
