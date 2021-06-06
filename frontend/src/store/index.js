import Vue from 'vue'
import Vuex, { Store } from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    backend_url: "http://"+ window.location.hostname + ":8080/api/v1",
    load: false,
  },
  mutations: {
    toggleLoad(state){
      state.load = !state.load;
    }
  },
  actions: {
    
  },
  modules: {
  }
})
