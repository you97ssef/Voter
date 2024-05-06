<script setup lang="ts">
    import { useLoadingStore } from '@/stores/loading';

    const loading = useLoadingStore()
</script>

<template>
    <Transition>
        <div class="absolute backdrop-blur-sm" v-if="loading.loadings > 0">
            <div class="flex justify-center items-center w-screen h-screen">
                <span class="loading loading-spinner loading-lg"></span>
            </div>
        </div>
    </Transition>
</template>


<style scoped>
.v-enter-active,
.v-leave-active {
  transition: opacity 0.5s ease;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}
</style>