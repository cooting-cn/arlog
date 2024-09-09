import {createApp} from 'vue'
import {createPinia} from 'pinia'
import 'uno.css'
import '@/assets/reset.css'
import '@/assets/public.css'
import '@/utils/message.js'
import App from './App.vue'
import router from './router/index.js'
import piniaPersist from 'pinia-plugin-persist'
import ECharts from "vue-echarts";

const app = createApp(App)

const pinia = createPinia()
pinia.use(piniaPersist)
app.component('v-chart', ECharts)
app.use(pinia)
app.use(router)
app.mount('#app')

