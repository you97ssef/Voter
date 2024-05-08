<script setup lang="ts">
    import Hero from '@/components/Hero.vue';
    import PrivatePollCard from '@/components/poll/PrivatePollCard.vue';
    import PublicPolls from '@/components/poll/PublicPolls.vue';
</script>

<template>
    <Hero></Hero>
    <PrivatePollCard></PrivatePollCard>
    <PublicPolls></PublicPolls>
</template>
