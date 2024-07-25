<script setup lang="ts">
import EventDetail from "./EventDetail.vue";
import useEvent from "@/lib/hook/useEvent";

const props = defineProps<{
  serviceName: string;
  eventType: string; 
}>()

const {isLoading,history,sourceID} =useEvent(props.serviceName, props.eventType)

</script>

<template>
<div>事件历史</div>

<div class="flex flex-col w-full h-full overflow-scroll rounded-lg gap-2">
    <div v-if="!isLoading">
      isLoading
    </div>
    <div v-else v-for="item in history">
        <EventDetail :event="item" :sourceID="sourceID" :eventType="props.eventType"  />
    </div>
    <div class="m-auto font-black" v-if="history.length == 0">
      ❗ 无事件
    </div>
  </div>
</template>