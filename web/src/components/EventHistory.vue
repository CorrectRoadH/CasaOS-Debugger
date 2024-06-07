<script setup lang="ts">
import createClient from "openapi-fetch";
import type { paths } from "../api/openapi" 
import { useDebuggerStore } from "@/stores/debugger";
import { ref, watch } from "vue";
import EventDetail from "./EventDetail.vue";


const props = defineProps<{
  sourceID: string;
  eventType?: string; 
}>()

const client = createClient<paths>({ baseUrl: "/v2/debugger/" });

const eventType = props.eventType == "all" ? undefined : props.eventType;

const { data, error } = await client.GET("/events",{
    params: {
        query: {
            sourceId: props.sourceID,
            eventType,
            offset: 0,
            length: 100
        }  
    }
})
</script>

<template>
<div>事件历史</div>

<div class="flex flex-col w-full h-full overflow-scroll rounded-lg gap-2">
    {{props.eventType}}
    <div v-for="item in data?.data">
        <EventDetail :event="item" :sourceId="props.sourceID" :eventType="props.eventType"  />
    </div>
</div>
</template>