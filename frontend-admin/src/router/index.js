import Vue from 'vue'
import VueRouter from 'vue-router'
import Dashboard from '../views/Dashboard.vue'
import Articles from '../views/Articles.vue'
import Login from '../views/Login.vue'

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
  },
  {
    path: '/admin/login',
    name: 'login',
    component: Login
  }
]

const router = new VueRouter({
  routes
})

export default router
