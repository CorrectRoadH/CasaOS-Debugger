<script setup lang="ts">
import createClient from "openapi-fetch";
import type { paths } from "../../api/openapi" 
import { ref, watch } from "vue";
import EventDetail from "./EventDetail.vue";
import ServiceMap from "@/lib/utils";

const props = defineProps<{
  serviceName: string;
  eventType: string; 
}>()

const client = createClient<paths>({ baseUrl: "/v2/debugger/" });


let history = ref([])
// @ts-ignore
const sourceID = ServiceMap[props.serviceName].eventName


const fetchHistory = async () => {
    const eventType = props.eventType == "all" ? undefined : props.eventType;

  const { data, error } = await client.GET("/events",{
      params: {
          query: {
              sourceId: sourceID,
              eventType,
              offset: 0,
              length: 100
          }  
      }
  })

  // @ts-ignore
  history.value = data?.data
}

console.log("refresh eventType")
fetchHistory()

watch(() => props.eventType, () => {
  fetchHistory()
})
</script>

<template>
<div>事件历史</div>

<div class="flex flex-col w-full h-full overflow-scroll rounded-lg gap-2">
    <div v-for="item in history">
        <EventDetail :event="item" :sourceID="sourceID" :eventType="props.eventType"  />
    </div>
    <div class="m-auto font-black" v-if="history.length == 0">
      ❗ 无事件
    </div>
  </div>
</template>