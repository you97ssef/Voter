<script setup>
    import Loading from './components/Loading.vue'
    import Toast from './components/Toast.vue'
    import Header from './components/Header.vue'
</script>

<template>
    <Loading />
    <Toast />
    <Header />
    <div class="p-4">
        <RouterView />
    </div>
</template>
