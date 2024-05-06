<script setup lang="ts">
    import router from '@/router';
    import { useConnectionStore } from '../../stores/connection';
    import { inject, ref } from 'vue';
    import type { Axios } from 'axios';
    import type { OkResponse } from '@/models/response';

    const connection = useConnectionStore()
    const http: Axios = inject('http') as Axios


    async function register() {
        console.log(form.value)
        const response = await http.post<OkResponse<string>>('/register', form.value)

        connection.setToken(response.data.data)
        router.push({ name: 'home' })
    }

    const form = ref({
        name: '',
        email: '',
        username: '',
        password: '',
    })

    const showPassword = ref(false)
</script>

<template>
    <div class="flex align-center justify-center">
        <div class="card bg-base-200 text-neutral-content w-full max-w-4xl">
            <form class="card-body">
                <div class="text-center text-success">
                    <i class="text-8xl fa-regular fa-circle-user"></i>
                    <p class="text-xl my-2">Register</p>
                </div>

                <label class="form-control w-full" for="name">
                    <div class="label">
                        <span class="label-text">
                            Name *
                        </span>
                    </div>
                    <input 
                        class="input input-bordered w-full"
                        id="name"
                        name="name" 
                        placeholder="John doe"
                        v-model="form.name"
                    />
                </label>

                <label class="form-control w-full" for="email">
                    <div class="label">
                        <span class="label-text">
                            Email *
                        </span>
                    </div>
                    <input 
                        class="input input-bordered w-full"
                        id="email"
                        name="email" 
                        placeholder="john@doe.com"
                        v-model="form.email"
                    />
                </label>

                <label class="form-control w-full" for="username">
                    <div class="label">
                        <span class="label-text">
                            Username *
                        </span>
                    </div>
                    <input 
                        class="input input-bordered w-full"
                        id="username"
                        name="username" 
                        placeholder="Username"
                        v-model="form.username"
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
                        v-model="form.password"
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
                    <button type="button" class="btn btn-success" @click="register">
                        <i class="fa-solid fa-arrow-right-to-bracket"></i>
                        Register
                    </button>
                </div>
                <div class="divider"></div>
                <div class="flex justify-center flex-wrap">
                    <RouterLink to="/login" class="btn btn-xs btn-primary m-1">
                        <i class="fa-solid fa-user-plus"></i>
                        Login
                    </RouterLink>
                </div>
            </form>
        </div>
    </div>
</template>
