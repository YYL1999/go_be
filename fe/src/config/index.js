
/**
 * axios 配置模块
 * @module config
 * @see utils/request
*/

/**
 *  axios具体配置对象
 * @description 包含了基础路径/请求前后对数据对处理，自定义请求头的设置等
 */
const axiosConfig = {
    baseURL: process.env.RESTAPI_PREFIX,
    // 请求前的数据处理
    // transformRequest: [function (data) {
    //   return data
    // }],
    // 请求后的数据处理
    // transformResponse: [function (data) {
    //   return data
    // }],
    // 自定义的请求头
    // headers: {
    //   'Content-Type': 'application/json'
    // },
    // 查询对象序列化函数
    // paramsSerializer: function (params) {
    //   return qs.stringify(params)
    // },
    // 超时设置s
    timeout: 10000,
    // 跨域是否带Token 项目中加上会出错
    // withCredentials: true,
    // 自定义请求处理
    // adapter: function(resolve, reject, config) {},
    // 响应的数据格式 json / blob /document /arraybuffer / text / stream
    responseType: 'json',
    // xsrf 设置
    xsrfCookieName: 'XSRF-TOKEN',
    xsrfHeaderName: 'X-XSRF-TOKEN',
    // 下传和下载进度回调
    onUploadProgress: function (progressEvent) {
      Math.round(progressEvent.loaded * 100 / progressEvent.total)
    },
    onDownloadProgress: function (progressEvent) {
      Math.round(progressEvent.loaded * 100 / progressEvent.total)
    },
    // 最多转发数，用于node.js
    maxRedirects: 5,
    // 最大响应数据大小
    maxContentLength: 2000,
    // 自定义错误状态码范围
    validateStatus: function (status) {
      return status >= 200 && status < 300
    }
    // 用于node.js
    // httpAgent: new http.Agent({ keepAlive: true }),
    // httpsAgent: new https.Agent({ keepAlive: true })
  }
  /** 导出配置模块 */
  export default axiosConfig
  