import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './registerServiceWorker'
import { MdApp, MdDrawer, MdDialog, MdCard, MdToolbar, MdList, MdIcon, MdButton, MdContent, MdTabs } from 'vue-material/dist/components'
import 'vue-material/dist/vue-material.min.css'
import 'vue-material/dist/theme/default.css'

Vue.use(MdApp)
Vue.use(MdDialog)
Vue.use(MdList)
Vue.use(MdIcon)
Vue.use(MdCard)
Vue.use(MdDrawer)
Vue.use(MdToolbar)
Vue.use(MdButton)
Vue.use(MdContent)
Vue.use(MdTabs)

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
