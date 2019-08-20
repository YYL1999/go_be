import Vue from 'vue'
import Vuex from 'vuex'
import article from './article'
Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'
export default new Vuex.Store({
    state:{},
    modules:{
      article
    },
    strict:debug
})