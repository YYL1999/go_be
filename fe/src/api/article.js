import request from '../utils/request'
/**
 * 获取全部文章
 */
export function getAllArticle() {
    return request.get('/getallarticle')
}