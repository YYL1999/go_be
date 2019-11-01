import {getAllArticle} from 'api/article'
const article = {
  state:{
      list:[]
  },
  mutations:{
      SET_LIST:(state,sdata) => {
          state.list = sdata
      }
  },
  actions: {
      getList: ({commit}) => {
          let data = getAllArticle()
          commit('SET_LIST',data)
      }
  }
}
export default article