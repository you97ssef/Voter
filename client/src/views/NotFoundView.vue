<template>
    <div class="my-20 gap-9">
        <div class="flex justify-center">
            <img src="@/assets/illustrations/404.svg" class="max-w-64 w-full" />
        </div>
        <div class="mt-10 text-center">
            <h1 class="text-5xl font-bold font-courgette">
                404 Not Found
            </h1>
            <p class="text-xl py-2">
                The page you are looking for does not exist.
            </p>
            <RouterLink class="btn btn-primary text-xl m-1" to="/">
                <i class="fa-solid fa-arrow-left"></i>
                Go Home
            </RouterLink>
        </div>
    </div>
</template>