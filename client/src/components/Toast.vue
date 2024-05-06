<script setup lang="ts">
import { useToastStore } from '../stores/toast';
const toastStore = useToastStore()
</script>

<template>
    <TransitionGroup name="list" tag="div" class="toast z-50">
        <div 
            v-for="toast in toastStore.toasts" 
            class="alert shadow-lg"
            :class="toast.type === 'error' ? 'alert-error' : 'alert-success'"
            :key="toast.message"
        >
            <i :class="toast.type === 'error' ? 'fa-circle-exclamation' : 'fa-circle-check'" class="fa-solid"></i>
            <span>{{ toast.message }}</span>
        </div>
    </TransitionGroup>
</template>

<style scoped>
.list-enter-active,
.list-leave-active {
  transition: all 0.5s ease;
}
.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateX(30px);
}
</style>