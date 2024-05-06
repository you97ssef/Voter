import { ref, computed, type Ref } from 'vue'
import { defineStore } from 'pinia'

const TOKEN_KEY = 'auth-token'

export const useConnectionStore = defineStore('connection', () => {
    const token: Ref<null | string> = ref(localStorage.getItem(TOKEN_KEY))

    function setToken(newToken: string) {
        token.value = newToken
        localStorage.setItem(TOKEN_KEY, newToken)
    }

    function removeToken() {
        token.value = null
        localStorage.removeItem(TOKEN_KEY)
    }

    const isAuthenticated = computed(() => token.value !== null)

    const user = computed<User>(() => {
        if (isAuthenticated.value) {
            const payload = token.value!.split('.')[1]
            const decoded = atob(payload)
            
            return JSON.parse(decoded)
        }

        return null
    })

    return { token, setToken, removeToken, isAuthenticated, user }
})

interface User {
    id: number
    name: string
    email: string
    verified_at: string | null
    username: string
}