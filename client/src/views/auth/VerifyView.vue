<script setup lang="ts">
    import router from '@/router';
    import { inject, onMounted } from 'vue';
    import type { Axios } from 'axios';
    import type { OkResponse } from '@/models/response';
    import { useToastStore } from '@/stores/toast';
    import { useConnectionStore } from '@/stores/connection';

    const toast = useToastStore()
    const connection = useConnectionStore()

    const http: Axios = inject('http') as Axios
    const token = router.currentRoute.value.query.token;


    onMounted(async () => {
        if (token) {
            const response = await http.get<OkResponse<string>>('/verify?token=' + token)
    
            connection.setToken(response.data.data)
            router.push({ name: 'home' })
        } else {
            toast.addToast('Link is invalid', 'error')
            router.push({ name: 'login' })
        }
    })
</script>

<template>
    <div class="text-center">
        <h1 class="text-2xl">
            Verification in progress...
        </h1>
        <RouterLink to="/" class="btn btn-sm btn-primary my-2">
            <i class="fa-solid fa-arrow-left"></i>
            Return to app
        </RouterLink>
    </div>
</template>