<script setup lang="ts">
import type { Axios } from 'axios';
import type { OkResponse } from '@/models/response';
import { useConnectionStore } from '../stores/connection';
import { inject } from 'vue';

const http: Axios = inject('http') as Axios
const connection = useConnectionStore()

function logout() {
    connection.removeToken()
}

async function resendVerificationEmail() {
    await http.get<OkResponse<null>>('/resend-verification', {
        params: {
            email: connection.user.email
        }
    })
}
</script>


<template>
    <div class="navbar bg-base-100 justify-between">
        <RouterLink class="btn btn-ghost font-bold text-4xl font-courgette text-center" to="/">Voter</RouterLink>
        <div class="flex gap-2" v-if="!connection.isAuthenticated">
            <RouterLink class="btn btn-primary btn-square" to="/login">
                <i class="fa-solid text-lg fa-arrow-right-to-bracket"></i>
            </RouterLink>
            <RouterLink class="btn btn-success btn-square" to="/register">
                <i class="fa-solid text-lg fa-user-plus"></i>
            </RouterLink>
        </div>
        <div class="flex gap-2" v-else>
            <button class="btn btn-error btn-square" @click="logout">
                <i class="fa-solid text-lg fa-arrow-right-from-bracket"></i>
            </button>
        </div>
    </div>
    <div class="divider font-courgette m-0 mx-4">Hi {{ connection.user?.name ?? 'Guest' }}</div>

    <div class="max-w-screen-md mx-auto p-4" v-if="connection.user?.verified_at === null">
        <div role="alert" class="alert alert-error">
            <i class="fa-solid fa-triangle-exclamation text-xl"></i>
            <span>
                Account is not verified. Please verify your account.
            </span>
            <div>
                <button class="btn btn-sm bg-success-content" @click="resendVerificationEmail">
                    <i class="fa-solid fa-envelope"></i> Resend Verification Email
                </button>
            </div>
        </div>
    </div>
</template>
