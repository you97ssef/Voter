<script setup lang="ts">
    import { buildPollsWithOptions, type Poll, type PollWithOptions, type Polls } from '@/models/poll';
    import PollCard from '@/components/poll/PollCard.vue';
    import type { Axios } from 'axios';
    import { inject, onMounted, ref, watch } from 'vue';
    import type { OkResponse } from '@/models/response';

    const http: Axios = inject('http') as Axios

    let polls: PollWithOptions[] = []
    let viewedPolls = ref<PollWithOptions[]>([])
    const search = ref('')

    watch(search, () => {
        viewedPolls.value = polls.filter(p => p.description.toLowerCase().includes(search.value.toLowerCase()))
    })

    onMounted(async () => {
        const response = await http.get<OkResponse<Polls>>('/my-polls')
        if (response.data.data.polls === null || response.data.data.options === null) return
        polls = buildPollsWithOptions(response.data.data.polls, response.data.data.options)
        viewedPolls.value = polls
    })
</script>

<template>
    <div class="flex justify-between max-w-screen-sm items-center mx-auto">
        <h1 class="text-center text-4xl py-4">My Polls</h1>
        <RouterLink class="btn btn-success" to="/new-poll">
            <i class="fa-solid fa-plus"></i>
            New Poll
        </RouterLink>
    </div>
    <label class="input input-bordered flex items-center gap-2 mb-4 max-w-screen-sm mx-auto">
        <i class="fa-solid fa-magnifying-glass opacity-70"></i>
        <input type="text" class="grow" placeholder="Search..." v-model="search" />
    </label>
    <div class="flex justify-center flex-wrap gap-5">
        <PollCard v-for="p in viewedPolls" :key="p.id" :poll="p"></PollCard>
        <div class="my-16" v-if="viewedPolls.length === 0">
            <img src="@/assets/illustrations/empty.svg" class="max-w-64 w-full" />
            <div class="text-center my-2 text-md">
                No polls found
            </div>
        </div>
    </div>
</template>