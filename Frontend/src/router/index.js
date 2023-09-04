import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '../views/HomeView.vue';
import EmailView from '../views/EmailView.vue';
import ContentEmailView from '../views/ContentEmailView.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/mails',
      name: 'mails',
      component: EmailView
    },
    {
      path: '/mails/:id',
      name: 'mail',
      component: ContentEmailView
    }
  ]
})

export default router
