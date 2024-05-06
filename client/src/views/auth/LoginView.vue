<script setup lang="ts">
    import router from '@/router';
    import { useConnectionStore } from '../../stores/connection';
    import { inject, ref } from 'vue';
    import type { Axios } from 'axios';
    import type { OkResponse } from '@/models/response';

    const connection = useConnectionStore()
    const http: Axios = inject('http') as Axios


    async function login() {
        const response = await http.post<OkResponse<string>>('/login', credentials.value)

        connection.setToken(response.data.data)
        router.push({ name: 'home' })
    }

    const credentials = ref({
        username_or_email: '',
        password: '',
    })

    const showPassword = ref(false)
</script>

<template>
    <div class="flex align-center justify-center">
        <div class="card bg-base-200 text-neutral-content w-full max-w-4xl">
            <form class="card-body" >
                <div class="text-center text-primary">
                    <i class="text-8xl fa-regular fa-circle-user"></i>
                    <p class="text-xl my-2">Login</p>
                </div>

                <label class="form-control w-full" for="username_email">
                    <div class="label">
                        <span class="label-text">
                            Username or Email *
                        </span>
                    </div>
                    <input 
                        class="input input-bordered w-full"
                        id="username_email"
                        name="username_email" 
                        placeholder="Username or Email"
                        v-model="credentials.username_or_email"
                        required
                    />
                </label>

                <label class="form-control w-full" for="password">
                    <div class="label">
                        <span class="label-text">
                            Password *
                        </span>
                    </div>
                    <input 
                        class="input input-bordered w-full"
                        id="password"
                        name="password" 
                        placeholder="Password"
                        v-bind:type="showPassword ? 'text' : 'password'"
                        v-model="credentials.password"
                        required
                    />
                </label>

                <div class="flex justify-end">
                    <div class="form-control">
                        <label class="label cursor-pointer" for="show_password">
                            <span class="label-text mr-2">
                                Show Password
                            </span> 
                            <input 
                                class="toggle" 
                                id="show_password"
                                name="show_password"
                                type="checkbox"
                                v-model="showPassword"
                            />
                        </label>
                    </div>
                </div>

                <div class="flex mt-5 justify-end">
                    <button 
                        type="button" 
                        class="btn btn-primary" 
                        @click="login"
                        :disabled="!credentials.username_or_email || !credentials.password"
                    >
                        <i class="fa-solid fa-arrow-right-to-bracket"></i>
                        Login
                    </button>
                </div>
                <div class="divider"></div>
                <div class="flex justify-center flex-wrap">
                    <RouterLink to="/register" class="btn btn-xs btn-success m-1">
                        <i class="fa-solid fa-user-plus"></i>
                        Register
                    </RouterLink>
                </div>
            </form>
        </div>
    </div>
</template>
