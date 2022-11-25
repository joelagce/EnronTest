import { createApp } from 'vue'
import router from './Router.js'
import App from './App.vue'
import './assets/css/tailwind.css'



createApp(App).use(router).mount('#app')
