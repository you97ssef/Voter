<script setup lang="ts">
    import type { NewPoll, Poll } from '@/models/poll';
    import type { OkResponse } from '@/models/response';
    import router from '@/router';
    import { useToastStore } from '@/stores/toast';
    import type { Axios } from 'axios';
    import { inject, ref } from 'vue';

    const toast = useToastStore()
    const http: Axios = inject('http') as Axios
    

    const newPoll = ref<NewPoll>({
        description: '',
        options: [],
        private: false
    })

    let option = ref('')

    function addOption() {
        if (option.value.length < 1) {
            toast.addToast('Option should not be empty.', 'error')
            return
        }
        newPoll.value.options.push(option.value)
        option.value = ''
    }

    function deleteOption(index: number) {
        newPoll.value.options.splice(index, 1)
    }

    async function createPoll() {
        if (newPoll.value.options.length < 2) {
            toast.addToast('You should add at least two options.', 'error')
            return
        }
        
        const response = await http.post<OkResponse<Poll>>('/polls', newPoll.value)
        
        response.data.data.private_code ?
            router.push({ name: 'private-poll', params: { code: response.data.data.private_code } }) :
            router.push({ name: 'poll', params: { id: response.data.data.id } })
    }
</script>

<template>
    <h1 class="text-center text-4xl py-4">New Poll</h1>
    <form @submit.prevent="createPoll">
        <div class="flex flex-wrap gap-4 justify-evenly items-center my-10">
            <div class="flex flex-col gap-2">
                <label class="input input-bordered flex items-center max-w-screen-sm gap-2">
                    <i class="fa-solid fa-pen opacity-70"></i>
                    <input type="text" placeholder="Poll description..." required v-model="newPoll.description" />
                </label>
                <div class="form-control">
                    <label class="cursor-pointer label gap-4">
                        <span>
                            Should this poll be private?
                        </span> 
                        <input type="checkbox" class="toggle toggle-primary" v-model="newPoll.private" />
                    </label>
                </div>
            </div>
            <div class="flex flex-col gap-2">
                <div class="text-center text-xl">Options</div>
                <p v-if="newPoll.options.length < 2" class="text-center text-error">
                    You should add at least two options.
                </p>
                <div class="flex gap-2 flex-wrap justify-center" v-if="newPoll.options.length > 0">
                    <div class="btn btn-sm badge badge-outline border-2 font-bold gap-2" v-for="(o, i) in newPoll.options" :key="o" @click="deleteOption(i)">
                        {{ o }}
                        <i class="fa-solid fa-xmark"></i>
                    </div>
                </div>
                <label class="input input-bordered flex items-center max-w-screen-sm gap-2">
                    <i class="fa-solid fa-note-sticky opacity-70"></i>
                    <input type="text" class="grow"  placeholder="Option..." v-model="option" />
                </label>
                <button type="button" class="btn btn-success btn-sm" @click="addOption">
                    <i class="fa-solid fa-plus"></i>
                    Add Option
                </button>
            </div>
        </div>
        <div class="flex justify-center">
            <button class="btn btn-primary" type="submit">
                <i class="fa-solid fa-poll"></i>
                Create New Poll
            </button>
        </div>
    </form>
</template>