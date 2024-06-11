import './assets/main.css'
import 'primevue/resources/themes/md-dark-indigo/theme.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import PrimeVue from 'primevue/config';
import Button from 'primevue/button';
//@ts-ignore
import timeago from 'vue-timeago3'

// define options


import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(PrimeVue);
app.component('Button', Button);

const timeagoOptions = {
    converterOptions: {
        includeSeconds: false,
    }
  }
  
  app.use(timeago,  timeagoOptions) // register timeago with options
  
app.mount('#app')
