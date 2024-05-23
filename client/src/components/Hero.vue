<script setup lang="ts">
    import { useConnectionStore } from '@/stores/connection';

    const connection = useConnectionStore()
</script>

<template>
    <div class="my-24 gap-9">
        <div class="flex justify-center">
            <img src="@/assets/illustrations/votes.svg" class="max-w-64 w-full" />
        </div>
        <div class="mt-10 text-center">
            <h1 class="text-5xl font-bold font-courgette">Voter</h1>
            <p class="text-xl py-2">
                Simple and easy to use polling application. Secured with the blockchain technology. Create, validate, watch live polls and share them with your friends.
            </p>
            <div class="flex flex-wrap justify-center gap-2" v-if="!connection.isAuthenticated">
                <RouterLink class="btn btn-primary text-xl m-1" to="/login">
                    <i class="fa-solid fa-arrow-right-to-bracket"></i>
                    Login
                </RouterLink>
                <RouterLink class="btn btn-success text-xl m-1" to="/register">
                    <i class="fa-solid fa-user-plus"></i>
                    Register
                </RouterLink>
            </div>
            <div class="flex flex-wrap justify-center gap-2" v-else>
                <RouterLink class="btn btn-primary text-xl m-1" to="/my-polls">
                    <i class="fa-solid fa-poll"></i>
                    My Polls
                </RouterLink>
                <RouterLink class="btn btn-success text-xl m-1" to="/new-poll">
                    <i class="fa-solid fa-plus"></i>
                    New Poll
                </RouterLink>
            </div>
        </div>
        <div class="text-center my-4 text-md">
            <i class="fa-solid fa-arrow-down-long"></i> 
            Scroll down to view polls
            <i class="fa-solid fa-arrow-down-long"></i>
        </div>
    </div>
</template>
