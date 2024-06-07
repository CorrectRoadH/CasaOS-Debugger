<script setup lang="ts">
// define props
import { defineProps } from 'vue';
import { useRouter } from 'vue-router';
// @ts-ignore
import JsonViewer from 'vue-json-viewer'

interface EventProperties {
  [key: string]: any; // 你可以定义具体的类型，而不是 `any`
}

const props = defineProps<{
    event: {
        name: string;
        properties: EventProperties;
        uuid: string;
        timestamp: string; // like `2024-06-05T09:24:56Z` 
    };
    sourceID: string;
    eventType: string; 
}>();

const propObjKeys = Object.keys(props.event.properties);

const convertToLocalTime = (ts:string) => {
    const date = new Date(ts);
    return date.toLocaleString();
};
const router = useRouter();

const eventNamebeClicked = () => {
    router.push(`/${props.sourceID}/${props.event.name}`);
};

</script>

<template>
    <div class="flex flex-col w-full bg-gray-200 p-2 rounded-lg">
        <div class="flex gap-2">

            <div class="text-blue-500 font-bold cursor-pointer"
                @click="eventNamebeClicked"
            >{{props.event.name}}</div>

            <div class="flex gap-2">
                时间: <div>{{convertToLocalTime(props.event.timestamp)}}</div>
            </div>
            <div class="flex gap-2">
                UUID: <div>{{props.event.uuid}}</div>
            </div>

    
        </div>

        <div v-for="key in propObjKeys" :key="key">
            <div class="flex flex-col flex-wrap gap-2">
                <div class="font-extrabold">{{key}}</div>
                <json-viewer :value="props.event.properties[key]" copyable></json-viewer>

            </div>
        </div>
    </div>
</template>

<style>
</style>