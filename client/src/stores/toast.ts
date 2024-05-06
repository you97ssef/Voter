import { ref, type Ref } from 'vue'
import { defineStore } from 'pinia'

const DURATION = 5000

interface Toast {
    message: string
    type: 'success' | 'error'
}

export const useToastStore = defineStore('toast', () => {
    const toasts: Ref<Toast[]> = ref([])

    function addToast(message: string, type: 'success' | 'error') {
        toasts.value.push({ message, type })
        setTimeout(() => toasts.value.shift(), DURATION)
    }

    return { toasts, addToast }
})
