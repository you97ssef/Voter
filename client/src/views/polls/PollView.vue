<script setup lang="ts">
    import { ref, onMounted, inject } from 'vue'
    import { addVote, buildResults, type PollResults, type Results } from '@/models/poll'
    import type { Axios } from 'axios'
    import { useRoute } from 'vue-router'
    import type { OkResponse } from '@/models/response';
    import { useToastStore } from '@/stores/toast';
    import router from '@/router';
    import { useConnectionStore } from '@/stores/connection';
    import type { NewGuestVote, NewVote, Vote } from '@/models/vote';

    const http: Axios = inject('http') as Axios
    const route = useRoute()
    const toastStore = useToastStore()
    const connection = useConnectionStore()

    const results = ref<Results>()
    const shareLink = ref<string>('')

    const alreadyVoted = ref(false)

    onMounted(async () => {
        try {
            const response = route.params.id ?
                await http.get<OkResponse<PollResults>>(`/polls/${route.params.id}`) :
            await http.get<OkResponse<PollResults>>(`/poll-by-code/${route.params.code}`)
            results.value = buildResults(response.data.data)
            checkIfAlreadyVoted()
            shareLink.value = results.value.poll.private_code === null ? 
                `${window.location.origin}/polls/${results.value.poll.id}` : 
                `${window.location.origin}/poll-by-code/${results.value.poll.private_code}`
        } catch (error) {
            console.error(error)
            router.push({ name: 'not-found' })
        }
    })

    function copyToClipboard() {
        navigator.clipboard.writeText(shareLink.value)
        toastStore.addToast('Link copied to clipboard', 'success')
    }

    function checkIfAlreadyVoted() {
        if (!connection.isAuthenticated) {
            if (!connection.guestUsername) {
                return
            }
            const guest = results.value!.votes.find(v => v.guest === connection.guestUsername)
            if (guest) {
                alreadyVoted.value = true
            }
        } else {
            const user = results.value!.votes.find(v => v.user_id === connection.user.id)
            if (user) {
                alreadyVoted.value = true
            }
        }
    }

    async function vote(option: number) {
        if (connection.isAuthenticated) {
            const vote: NewVote = {
                option_id: option,
                poll_id: results.value!.poll.id,
                poll_code: results.value!.poll.private_code
            }
    
            const response = await http.post<OkResponse<Vote>>('/votes', vote)
            results.value = addVote(results.value!, response.data.data)
            alreadyVoted.value = true
        } else {
            if (!connection.guestUsername) {
                toastStore.addToast('Please enter your name to vote', 'error')
                return
            }

            const guestVote: NewGuestVote = {
                option_id: option,
                poll_id: results.value!.poll.id,
                guest: connection.guestUsername,
                poll_code: results.value!.poll.private_code
            }

            const response = await http.post<OkResponse<Vote>>('/guest-votes', guestVote)
            results.value = addVote(results.value!, response.data.data)
            alreadyVoted.value = true
        }
    }

    function deletePoll() {
        if (window.confirm('Are you sure you want to delete this poll?')) {
            http.delete(`/polls/${results.value!.poll.id}`)
            router.push({ name: 'my-polls' })
        }
    }
</script>

<template>
    <div v-if="results">
        <div class="text-center mb-4">
            <p class="text-sm">Description</p>
            <h1 class="text-6xl mb-2 font-bold">{{ results.poll.description }}</h1>
            <p class="text-sm"><i class="fa-solid fa-users"></i> {{ results.count }} Voters</p>
            <button class="btn btn-error mt-2" @click="deletePoll" v-if="connection.isAuthenticated && results.poll.user_id === connection.user.id">
                <i class="fa-solid fa-trash"></i>
                Delete Poll
            </button>
        </div>
        <div class="flex flex-wrap justify-center items-center gap-2 mb-5">
            <div class="card card-compact bg-secondary text-secondary-content">
                <div class="card-body flex-row justify-between items-center">
                    <div class="flex justify-center items-center font-bold gap-2"> 
                        <i class="fa-solid fa-link"></i>
                        {{ shareLink }}
                    </div>
                    <div class="btn btn-sm btn-primary" @click="copyToClipboard">
                        Share
                        <i class="fa-solid fa-share-nodes"></i>
                    </div>
                </div>
            </div>
        </div>
        <p class="text-center text-sm mb-2">Options</p>
        
        <div class="flex justify-center mb-2" v-if="!connection.isAuthenticated">
            <label class="input input-bordered input-sm flex items-center gap-2">
                <i class="fa-solid fa-person-circle-question opacity-70"></i>
                <input type="text" class="grow" placeholder="Your name to vote" v-model="connection.guestUsername" />
            </label>
        </div>
    
        <div class="flex flex-wrap justify-center items-center gap-2 mb-5">
            <div class="card w-80 bg-primary text-primary-content" v-for="option in results.options" :key="option.details.id">
                <div class="card-body justify-center items-center">
                    <p>{{ option.details.description }}</p>
                    <div class="radial-progress bg-base-100 text-base-content border-4 border-base-100" :style="'--value:' + option.percentage + ';'" role="progressbar">{{ option.percentage }} %</div>
                    <button class="btn btn-sm btn-success" v-if="!alreadyVoted" @click="vote(option.details.id)">
                        <i class="fa-solid fa-vote-yea"></i>
                        Vote
                    </button>
                </div>
            </div>    
        </div>
        <p class="text-center text-sm mb-2">Votes</p>
        <div class="flex justify-center flex-wrap gap-5" v-if="results.votes.length === 0">
            <div class="my-2">
                <img src="@/assets/illustrations/empty.svg" class="max-w-64 w-full" />
                <div class="text-center my-2 text-md">
                    No votes found
                </div>
            </div>
        </div>
        <div class="carousel carousel-center max-w-full space-x-4 rounded-box">
            <div class="carousel-item" v-for="vote in results.votes" :key="vote.id">
                <div class="card bg-base-200 text-base-content w-64">
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
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
