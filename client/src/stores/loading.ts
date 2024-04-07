import { ref, type Ref } from 'vue'
import { defineStore } from 'pinia'

export const useLoadingStore = defineStore('loading', () => {
    const loadings: Ref<number> = ref(0)

    function addLoading() {
        loadings.value++
    }

    function removeLoading() {
        loadings.value--
    }

    return { loadings, addLoading, removeLoading }
})
