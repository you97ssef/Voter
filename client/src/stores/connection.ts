import { ref, computed, type Ref, watch } from 'vue'
import { defineStore } from 'pinia'

const TOKEN_KEY = 'auth-token'
const GUEST_USERNAME_KEY = 'guest-username'

export const useConnectionStore = defineStore('connection', () => {
    const token: Ref<null | string> = ref(localStorage.getItem(TOKEN_KEY))
    const guestUsername = ref<string | null>(localStorage.getItem(GUEST_USERNAME_KEY))

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

    watch(guestUsername, (newGuestUsername) => {
        if (newGuestUsername !== '' && newGuestUsername !== null) {
            localStorage.setItem(GUEST_USERNAME_KEY, newGuestUsername)
        } else {
            localStorage.removeItem(GUEST_USERNAME_KEY)
        }
    })

    return { token, setToken, removeToken, isAuthenticated, user, guestUsername }
})

interface User {
    id: number
    name: string
    email: string
    verified_at: string | null
    username: string
}