<script setup lang="ts">
import { useRouter } from 'vue-router';
import { ref } from "vue";
import ServiceMap from "@/lib/utils";

const beSelectServiceName = ref<string | null>(null);
const router = useRouter();

const handleSourceIDbeClicked = (sourceID:string) => {
    router.push(`/${sourceID}/all`);
    beSelectServiceName.value = sourceID;
}

const handleHomebeClicked = () => {
    router.push(`/`);
    beSelectServiceName.value = null;
}


</script>

<template>
     <div class="flex flex-col gap-2 p-2">
        <div
        @click="handleHomebeClicked()"
        :class="{ 
           'p-2 rounded-lg ': true,
           'p-card shadow-lg' : beSelectServiceName === null,
        }"
        
        >Home</div>

        <div v-for="service in ServiceMap" @click="handleSourceIDbeClicked(service.name)" class="cursor-pointer">
            <div
                 :class="{ 
                    'p-2 rounded-lg ': true,
                    'p-card	 shadow-lg' : beSelectServiceName === service.name,
                 }"
            >{{ service.name }}</div>
        </div>

     </div>
</template>

<style>
highlighted {
    background-color: var(--highlight-bg)
}
</style>