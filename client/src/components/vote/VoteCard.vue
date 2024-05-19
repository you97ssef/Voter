<script setup lang="ts">
    import type { Vote } from '@/models/vote';
    import { defineProps } from 'vue'

    defineProps<{
        vote: Vote
    }>()
</script>

<template>
    <div class="card bg-base-200 text-base-content flex-none w-64">
        <div class="card-body justify-center">
            <div>
                <p class="text-xs">Voter</p>
                <p class="text-xl font-bold">
                    <span>{{ vote.guest ?? vote.user_id }}</span>
                    <span v-if="vote.guest"> (guest)</span>
                </p>
            </div>
            <div>
                <p class="text-xs">Option</p>
                <h3 class="text-xl font-bold">{{ vote.option_id }}</h3>
            </div>
            <div v-if="vote.valid !== undefined">
                <p v-if="vote.valid" class="font-bold text-success">
                    <i class="fa-regular fa-circle-check"></i>
                    Valid
                </p>
                <p v-else class="font-bold text-error">
                    <i class="fa-regular fa-circle-xmark"></i>
                    Invalid
                </p>
            </div>
        </div>
    </div>
</template>