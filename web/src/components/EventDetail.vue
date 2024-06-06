<script setup lang="ts">
// define props
import { useDebuggerStore } from '@/stores/debugger';
import { defineProps } from 'vue';

const props = defineProps<{
    event: {
        name: string;
        properties: string;
        uuid: string;
        timestamp: string; // like `2024-06-05T09:24:56Z` 
    };
}>();

const propObjKeys = Object.keys(props.event.properties);

const debuggerStore = useDebuggerStore();

const convertToLocalTime = (ts:string) => {
    const date = new Date(ts);
    return date.toLocaleString();
};

const eventNamebeClicked = () => {
    console.log("eventNamebeClicked", props.event.name);
    debuggerStore.selectedEventType = {name: props.event.name};
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
            <div class="flex gap-2">
                <div class="font-extrabold">{{key}}</div>

                <!-- not overflow the witdh -->
                <div class="event-property">{{props.event.properties[key]}}</div>
            </div>
        </div>
    </div>
</template>
<style>
.event-property {
    width: 100%;
    overflow: hidden; /* 隐藏超出部分 */
    text-overflow: ellipsis; /* 添加省略号 */
  }
</style>