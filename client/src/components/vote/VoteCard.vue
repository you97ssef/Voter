<script setup lang="ts">
    import type { Option } from '@/models/option';
import type { Vote } from '@/models/vote';

    defineProps<{
        vote: Vote,
        options: {
            details: Option;
            percentage: number;
            count: number;
        }[]
    }>()
</script>

<template>
    <div class="card bg-base-200 text-base-content flex-none w-64">
        <div class="card-body justify-center">
            <div>
                <p class="text-xs">Voter</p>
                <p class="text-xl font-bold">
                    <span>{{ vote.user }}</span>
                    <span v-if="vote.is_guest"> (guest)</span>
                </p>
            </div>
            <div>
                <p class="text-xs">Option</p>
                <h3 class="text-xl font-bold">
                    {{ options.find(option => option.details.id === vote.option_id)?.details.description }}
                </h3>
            </div>
            <p class="text-xs">{{ new Date(vote.timestamp * 1000).toLocaleString() }}</p>
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