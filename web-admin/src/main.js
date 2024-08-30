import {createApp} from 'vue'
import {createPinia} from 'pinia'

import App from './App.vue'
import router from './router'
import piniaPersist from 'pinia-plugin-persist'
import 'uno.css'
import '@/assets/reset.css';
import '@/assets/public.css';

const app = createApp(App)
const pinia = createPinia()
pinia.use(piniaPersist)
app.use(pinia)
app.use(router)
app.mount('#app')
