
/**
 * 业务中使用的ajax请求工具模块
 * @module utils/request
 * @see main.js
 */

import axios from 'axios'
import config from '../config'
import Vue from 'vue'


// 用来调用errorTip的vue实例
const vueInstance = new Vue()
// 构建得的请求对象
const request = axios.create(config)
// 请求拦截器
// request.interceptors.request.use(
//   config => {
//     // 触发loading效果
//     store.dispatch('setLoading', true)
//     return config
//   },
//   error => {
//     return Promise.reject(error)
//   }
// )
// 返回状态判断(添加响应拦截器)
// request.interceptors.response.use(

//   (res) => {
//     // 加载loading
//     store.dispatch('setLoading', false)
//     if (res.data.data instanceof Object) {
//       let key = Object.keys(res.data.data)
//       if (key.length === 0) {
//         errorTip(vueInstance, {}, '没有数据')
//       }
//     } else if (res.data.data instanceof Array) {
//       if (res.data.data.length === 0) {
//         errorTip(vueInstance, {}, '没有数据')
//       }
//     }
//     // 主动控制是否弹出提示
//     let tip = Url.parse(res.config.url, true).query['fet'] !== 'false'
//     let code = res.data.code
//     // 成功
//     if (code >= 200 && code < 300) {
//       // 如果请求不是get并且没有要求不提示，就提示
//       if (res.config.method !== 'get' && tip) successTip(vueInstance)
//       return res.data.data
//     }
//     // 如果数据请求失败
//     let message = ''
//     let prefix = res.config.method !== 'get' ? '操作失败：' : '请求失败：'
//     switch (code) {
//       case 400: message = prefix + '请求参数缺失'; break
//       case 401: message = prefix + '认证未通过'; break
//       case 404: message = prefix + '此数据不存在'; break
//       case 406: message = prefix + '条件不满足'; break
//       default: message = prefix + '服务器出错了'; break
//     }
//     let error = new Error(message)

//     if (tip) {
//       errorTip(vueInstance, error, message)
//     }
//     let result = { ...res.data, error: error }
//     process.env.NODE_ENV !== 'production' && console.log(error)
//     return result
//   },
//   (error, a, b) => {
//     store.dispatch('setLoading', false)
//     errorTip(vueInstance, error)
//     process.env.NODE_ENV !== 'production' && console.log(error)
//     return { data: null, code: 500, error: error }
//   }
// )

// 对axios的实例重新封装成一个plugin ,方便 Vue.use(xxxx)
export default request
