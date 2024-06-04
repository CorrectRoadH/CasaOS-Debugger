<script setup lang="ts">
// define props
import { defineProps } from 'vue';

const props = defineProps<{
    event: {
        name: string;
        properties: string;
        uuid: string;
    };
}>();

const propObjKeys = Object.keys(props.event.properties);

</script>

<template>
    <div class="flex flex-col w-full bg-gray-200 p-2 rounded-lg">
        <div>{{props.event.name}}</div>
        <div>{{props.event.uuid}}</div>

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