<script setup lang="ts">
    import type { PollWithOptions } from '@/models/poll';
    import { defineProps } from 'vue'

    defineProps<{
        poll: PollWithOptions
    }>()
</script>

<template>
    <div class="card w-80 bg-primary text-primary-content">
        <div class="card-body justify-between">
            <div class="pb-3">
                <p class="text-sm">
                    Description
                    <span v-if="poll.private_code">(private)</span>
                </p> 

                <h3 class="text-3xl font-bold">{{ poll.description }}</h3>
            </div>
            <div>
                <p class="text-sm">Options</p>
                <div class="flex flex-wrap gap-2">
                    <div class="badge badge-outline border-2 font-bold flex-none" v-for="option in poll.options" :key="option.id">
                        {{ option.description }}
                    </div>
                </div>
            </div>
            <div class="card-actions justify-end">
                <RouterLink v-if="poll.private_code" :to="{ name: 'private-poll', params: { code: poll.private_code }}" class="btn btn-accent">
                    <i class="fa-solid fa-check-to-slot text-lg"></i>    
                    Show poll
                </RouterLink>
                <RouterLink v-else :to="{ name: 'poll', params: { id: poll.id }}" class="btn btn-accent">
                    <i class="fa-solid fa-check-to-slot text-lg"></i>    
                    Show poll
                </RouterLink>
            </div>
        </div>
    </div>
</template>