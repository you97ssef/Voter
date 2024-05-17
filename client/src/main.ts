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
app.use(http, "http://localhost:8888")
app.use(live, "ws://broker.emqx.io:8083/mqtt", "voterapp/")

app.mount('#app')
