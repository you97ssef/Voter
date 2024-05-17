import { LiveService } from '@/services/live-service'
import { type App } from 'vue'

export default {
    install: (app: App, host: string, topic: string) => {
        const live = new LiveService(host, topic)

        app.provide('live', live)
    }
}
