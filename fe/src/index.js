import 'babel-polyfill'
import Vue from 'vue'
import App from './App'
import router from './router'
import  UI from  './components'
// import {Pagination} from 'element-ui'
Vue.config.productionTip = false 
// Vue.use(Pagination)
Vue.use(UI)
new Vue({
    el:'#app',
    router,
    components:{App},
    template:'<App/>'
})