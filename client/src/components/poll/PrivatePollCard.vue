<script setup lang="ts">
import router from '@/router';

    let code = ''

    function viewPoll() {
        router.push({ name: 'private-poll', params: { code } })
    }
</script>

<template>
    <div class="flex flex-wrap justify-center mb-20">
        <div class="indicator min-w-80">
            <span class="indicator-item indicator-center badge badge-lg badge-primary">
                <i class="fa-solid fa-lock"></i>
            </span>
            <div class="card min-w-80 bg-primary-content shadow-xl h-full">
                <div class="card-body text-center">
                    <p class="text-xl">
                        Access private poll by entering it's private code.
                    </p>
                    <label class="input input-bordered flex items-center gap-2">
                        <i class="fa-solid fa-lock opacity-70"></i>
                        <input type="text" class="grow" placeholder="Private code..." v-model="code" />
                    </label>
                    <button class="btn btn-primary text-xl mt-2" @click="viewPoll">
                        <i class="fa-solid fa-eye"></i> View Poll
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>