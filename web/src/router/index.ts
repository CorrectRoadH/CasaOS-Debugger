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
      path: '/:sourceID/:eventType',
      name: 'service-event',
      component: ServiceView,
      props: (route) => ({ 
        sourceID: route.params.sourceID,
        eventType: route.params.eventType
      })
    },
  ]
})

export default router
