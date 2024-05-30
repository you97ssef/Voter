<script setup lang="ts">
    import { ref, onMounted, inject, onUnmounted, computed } from 'vue'
    import { addVote, buildResults, type Poll, type PollResults, type Results } from '@/models/poll'
    import type { Axios } from 'axios'
    import { useRoute } from 'vue-router'
    import type { OkResponse } from '@/models/response';
    import { useToastStore } from '@/stores/toast';
    import router from '@/router';
    import { useConnectionStore } from '@/stores/connection';
    import { buildValidations, type NewGuestVote, type NewVote, type ValidatedVote, type Vote } from '@/models/vote';
    import VoteCard from '@/components/vote/VoteCard.vue';
    import type { LiveService } from '@/services/live-service';

    const http: Axios = inject('http') as Axios
    const route = useRoute()
    const toastStore = useToastStore()
    const connection = useConnectionStore()
    const live = inject('live') as LiveService

    const results = ref<Results>()
    const shareLink = ref<string>('')

    const alreadyVoted = computed(() => {
        return connection.isAuthenticated ? 
            results.value!.votes.some(v => v.user === connection.user.name) : 
            connection.guestUsername ? 
                results.value!.votes.some(v => v.user === connection.guestUsername) : 
                false    
    })

    onMounted(async () => {
        try {
            await getPoll()
            buildShareLink()
            subscribeToVotes()
        } catch (error) {
            router.push({ name: 'not-found' })
        }
    })

    async function getPoll() {
        const response = route.params.id ?
                await http.get<OkResponse<PollResults>>(`/polls/${route.params.id}`) :
            await http.get<OkResponse<PollResults>>(`/poll-by-code/${route.params.code}`)
            results.value = buildResults(response.data.data)
    }

    function buildShareLink() {
        shareLink.value = results.value!.poll.private_code === null ? 
            `${window.location.origin}/polls/${results.value!.poll.id}` : 
            `${window.location.origin}/poll-by-code/${results.value!.poll.private_code}`
    }

    function subscribeToVotes() {
        live.connect()
        live.subscribe(`poll-${results.value!.poll.id}`, (message: string) => {
            try {
                const vote = JSON.parse(message) as Vote
                results.value = addVote(results.value!, vote)
                validated.value = false
                toastStore.addToast('New vote received! Check its validity', 'success')
            } catch (error) {
                toastStore.addToast('Error receiving new vote', 'error')
            }
        })
    }

    function copyToClipboard() {
        navigator.clipboard.writeText(shareLink.value)
        toastStore.addToast('Link copied to clipboard', 'success')
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
            validated.value = false
            publishVote(response.data.data)
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
            validated.value = false
            publishVote(response.data.data)
        }
    }

    onUnmounted(() => {
        if (!results.value) return
        live.unsubscribe(`poll-${results.value.poll.id}`)
        live.disconnect()
    })

    function publishVote(vote: Vote) {
        live.publish(`poll-${results.value!.poll.id}`, JSON.stringify(vote))
    }

    async function deletePoll() {
        if (window.confirm('Are you sure you want to delete this poll?')) {
            await http.delete(`/polls/${results.value!.poll.id}`)
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
                <h1 class="text-6xl mb-2 font-bold">
                    {{ results.poll.description }}
                    <span v-if="results.poll.finished_at === null" class="loading loading-ring loading-lg text-error"></span>
                </h1>
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
                    <div class="card-body flex-row justify-between items-center flex-wrap">
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
            <VoteCard v-for="vote in results.votes" :key="vote.id" :vote="vote" :options="results.options" />
        </div>
    </div>
</template>
