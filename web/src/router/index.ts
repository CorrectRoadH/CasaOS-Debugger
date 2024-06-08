import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ServiceView from '../views/ServiceView.vue'
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/:serviceName/:eventType',
      name: 'service-event',
      component: ServiceView,
      props: (route) => ({ 
        serviceName: route.params.serviceName,
        eventType: route.params.eventType
      })
    },
  ]
})

export default router
