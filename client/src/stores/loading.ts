import { ref, type Ref } from 'vue'
import { defineStore } from 'pinia'

export const useLoadingStore = defineStore('loading', () => {
    const loadings: Ref<number> = ref(0)

    async function addLoading() {
        loadings.value++

        await new Promise((resolve) => setTimeout(resolve, 10000))
    }

    function removeLoading() {
        loadings.value--
    }

    return { loadings, addLoading, removeLoading }
})
