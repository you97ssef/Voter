import { createRouter, createWebHistory, type NavigationGuardNext, type RouteLocationNormalized } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '@/views/auth/LoginView.vue'
import RegisterView from '@/views/auth/RegisterView.vue'
import VerifyView from '@/views/auth/VerifyView.vue'
import { useConnectionStore } from '@/stores/connection'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomeView
        },
        {
            path: '/login',
            name: 'login',
            component: LoginView,
            beforeEnter: authGuard
        },
        {
            path: '/register',
            name: 'register',
            component: RegisterView,
            beforeEnter: authGuard
        },
        {
            path: '/verify',
            name: 'verify',
            component: VerifyView,
        }
    ]
})

function authGuard(to: RouteLocationNormalized, from: RouteLocationNormalized, next: NavigationGuardNext) {
    if (useConnectionStore().isAuthenticated) {
        next({ name: 'home' })
    } else next()
}


export default router
