import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Articles from '../views/Articles.vue'
import Article from '../views/Article.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/articles',
    name: 'Articles',
    component: Articles
  },
  {
    path: '/article/:id',
    name: 'Article',
    component: Article
  },
  
]

const router = new VueRouter({
  routes
})

export default router
