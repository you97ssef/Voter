import { useConnectionStore } from '@/stores/connection'
import { useLoadingStore } from '@/stores/loading'
import { useToastStore } from '@/stores/toast'
import axios from 'axios'
import type { App } from 'vue'

export default {
    install: (app: App, url: string) => {
        const connection = useConnectionStore()
        const toast = useToastStore()
        const loading = useLoadingStore()

        const http = axios.create({
            baseURL: url,
            headers: {
                Authorization: connection.isAuthenticated ? `Bearer ${connection.token}` : '',
            },
        })

        http.interceptors.request.use(
            (config) => {
                loading.addLoading()

                return config
            },
            (error) => {
                loading.removeLoading()

                return Promise.reject(error)
            },
        )

        http.interceptors.response.use(
            (response) => {
                if (response.data?.message) {
                    toast.addToast(response.data.message, 'success')
                }
            
                loading.removeLoading()
                return response
            },
            (error) => {
                if (error.response?.data?.message) {
                    toast.addToast(error.response?.data.message, 'error')
                }

                loading.removeLoading()
                return Promise.reject(error)
            },
        )

        app.provide('http', http)
    }
}
