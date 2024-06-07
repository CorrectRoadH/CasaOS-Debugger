<script setup lang="ts">
// import Button from 'primevue/button';

import EventTypeList from "@/components/EventTypeList.vue"
import EventTypeDetail from "@/components/EventTypeDetail.vue";
import EventHistory from "@/components/EventHistory.vue";
// import { watch } from "vue";

const props = defineProps<{
  sourceID: string;
  eventType: string; 
}>()

// watch(() => props.sourceID, (newSourceID) => {
//   console.log("sourceID", newSourceID);
// });
</script>

<template>
    <div class="flex m-auto">
      <!-- header -->
      <Button label="事件" />
      <Button label="日志" disabled />
    </div>

    <div class="flex gap-2 h-full" :key="props.sourceID">
      <div class="flex flex-col h-full">
        <Suspense>
          <EventTypeList :sourceID="props.sourceID" :eventType="props.eventType" />
        </Suspense>
      </div>

      <div class="flex flex-col w-full h-full gap-2">
        <EventTypeDetail :sourceID="props.sourceID" :eventType="props.eventType" />
        <EventHistory :sourceID="props.sourceID" :eventType="props.eventType" />
      </div>
    </div>
</template>

