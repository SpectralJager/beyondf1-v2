import Vue from 'vue'
import VueRouter from 'vue-router'
import Dashboard from '../views/Dashboard.vue'
import Articles from '../views/Articles.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/admin/dashboard',
    name: 'dashboard',
    component: Dashboard
  },
  {
    path: '/admin/articles',
    name: 'articles',
    component: Articles,
  }
]

const router = new VueRouter({
  routes
})

export default router
