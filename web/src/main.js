import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

/*引入ElementUI*/
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
Vue.use(ElementUI);

// 使用bus通信
import Bus from '@/utils/bus.js'
Vue.use(Bus)

// 设置vue是否使用生成模式
Vue.config.productionTip = false

// 设置拦截器
import '@/permission'

export default new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
