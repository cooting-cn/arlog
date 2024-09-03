import {createApp} from 'vue'
import {createPinia} from 'pinia'
import 'uno.css'
import '@/assets/reset.css';
import '@/assets/public.css';
import App from './App.vue'
import index from './router/index.js'
import piniaPersist from 'pinia-plugin-persist'


const app = createApp(App)
const pinia = createPinia()
pinia.use(piniaPersist)
app.use(pinia)
app.use(index)
app.mount('#app')
