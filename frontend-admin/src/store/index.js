import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    token: "",
  },
  mutations: {
    logout(){
      state.token = "1";
    }
  },
  actions: {
  },
  modules: {
  }
})
