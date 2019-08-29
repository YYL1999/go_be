import 'babel-polyfill'
import Vue from 'vue'
import App from './App'
import router from './router'
import  UI from  './components'
import {Container} from 'element-ui'
// import {Pagination} from 'element-ui'
Vue.config.productionTip = false 
// Vue.use(Pagination)
Vue.use(UI)
Vue.use(Container)
new Vue({
    el:'#app',
    router,
    components:{App},
    template:'<App/>'
})