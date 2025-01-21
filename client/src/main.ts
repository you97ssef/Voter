import '@fortawesome/fontawesome-free/css/all.min.css'
import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import http from './plugins/http'
import live from './plugins/live'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(http, process.env.NODE_ENV === 'production' ? 'https://api.voter.youssefbahi.com' : 'http://localhost:9000')
app.use(live, "wss://broker.emqx.io:8084/mqtt", "voterapp/")

app.mount('#app')
