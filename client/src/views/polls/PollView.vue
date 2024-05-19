<script setup lang="ts">
    import { ref, onMounted, inject } from 'vue'
    import { addVote, buildResults, type Poll, type PollResults, type Results } from '@/models/poll'
    import type { Axios } from 'axios'
    import { useRoute } from 'vue-router'
    import type { OkResponse } from '@/models/response';
    import { useToastStore } from '@/stores/toast';
    import router from '@/router';
    import { useConnectionStore } from '@/stores/connection';
    import { buildValidations, type NewGuestVote, type NewVote, type ValidatedVote, type Vote } from '@/models/vote';
    import VoteCard from '@/components/vote/VoteCard.vue';

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
            validated.value = false
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
            validated.value = false
        }
    }

    function deletePoll() {
        if (window.confirm('Are you sure you want to delete this poll?')) {
            http.delete(`/polls/${results.value!.poll.id}`)
            router.push({ name: 'my-polls' })
        }
    }

    async function closePoll() {
        if (window.confirm('Are you sure you want to close this poll?')) {
            const response = await http.put<OkResponse<Poll>>(`/polls/${results.value!.poll.id}`)
            results.value!.poll.finished_at = response.data.data.finished_at
        }
    }

    async function validateVotes() {
        const response = route.params.id ?
                await http.get<OkResponse<ValidatedVote[]>>(`/validate-poll/${route.params.id}`) :
            await http.get<OkResponse<ValidatedVote[]>>(`/validate-poll-by-code/${route.params.code}`)

        results.value!.votes = buildValidations(results.value!.votes, response.data.data)
        validated.value = true
    }

    const validated = ref(false)
</script>

<template>
    <div v-if="results">
        <div class="flex flex-wrap items-center justify-evenly gap-4 my-16">
            <div class="text-center">
                <h1 class="text-6xl mb-2 font-bold">{{ results.poll.description }}</h1>
                <p class="text-xs" v-if="results.poll.finished_at !== null">Closed {{ new Date(results.poll.finished_at).toLocaleString() }}</p>
                <p class="text-sm"><i class="fa-solid fa-users"></i> {{ results.count }} Voters</p>
                <button class="btn btn-error m-1" @click="closePoll" v-if="connection.isAuthenticated && results.poll.user_id === connection.user.id && results.poll.finished_at === null">
                    <i class="fa-solid fa-ban"></i>
                    Close Poll
                </button>
                <button class="btn btn-error m-1" @click="deletePoll" v-if="connection.isAuthenticated && results.poll.user_id === connection.user.id">
                    <i class="fa-solid fa-trash"></i>
                    Delete Poll
                </button>
            </div>
            <div class="flex flex-col gap-2">
                <div class="card card-compact bg-primary text-secondary-content">
                    <div class="card-body flex-row justify-between items-center">
                        <div class="flex justify-center items-center font-bold gap-2"> 
                            <i class="fa-solid fa-link"></i>
                            {{ shareLink }}
                        </div>
                        <div class="btn btn-sm btn-accent" @click="copyToClipboard">
                            Share
                            <i class="fa-solid fa-share-nodes"></i>
                        </div>
                    </div>
                </div>
                <div class="flex justify-center mb-2" v-if="!connection.isAuthenticated">
                    <label class="input input-bordered flex items-center gap-2">
                        <i class="fa-solid fa-person-circle-question opacity-70"></i>
                        <input type="text" class="grow" placeholder="Your name to vote" v-model="connection.guestUsername" />
                    </label>
                </div>
            </div>
        </div>
        <p class="text-center text-sm mb-4">Options</p>
        <div class="flex flex-wrap justify-center items-center gap-2 mb-16">
            <div class="card w-80 bg-primary text-primary-content" v-for="option in results.options" :key="option.details.id">
                <div class="card-body justify-center items-center">
                    <p>{{ option.details.description }}</p>
                    <div class="radial-progress bg-base-100 text-base-content border-4 border-base-100" :style="'--value:' + option.percentage + ';'" role="progressbar">{{ option.percentage }} %</div>
                    <p>Count: {{ option.count }}</p>
                    <button class="btn btn-sm btn-success" v-if="!alreadyVoted && results.poll.finished_at === null" @click="vote(option.details.id)">
                        <i class="fa-solid fa-vote-yea"></i>
                        Vote
                    </button>
                </div>
            </div>    
        </div>
        <p class="text-center text-sm mb-4">Votes</p>
        <div class="flex flex-wrap justify-center mb-4" v-if="!validated">
            <button type="button" class="btn btn-primary flex-none" @click="validateVotes">
                <i class="fa-solid fa-check"></i>
                Check if votes are valid
            </button>
        </div>
        <div class="flex justify-center flex-wrap gap-5" v-if="results.votes.length === 0">
            <div class="my-2">
                <img src="@/assets/illustrations/empty.svg" class="max-w-64 w-full" />
                <div class="text-center my-2 text-md">
                    No votes found
                </div>
            </div>
        </div>
        <div class="flex flex-wrap justify-center gap-4">
            <VoteCard v-for="vote in results.votes" :key="vote.id" :vote="vote" />
        </div>
    </div>
</template>
