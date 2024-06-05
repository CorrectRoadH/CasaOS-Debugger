<script setup lang="ts">
// import Button from 'primevue/button';

import EventTypeList from "@/components/EventTypeList.vue"
import EventTypeDetail from "@/components/EventTypeDetail.vue";
import EventHistory from "@/components/EventHistory.vue";

</script>

<template>
    <div class="flex m-auto">
      <!-- header -->
      <Button label="事件" />
      <Button label="日志" disabled />
    </div>

    <div class="flex gap-2 h-full">
      <div class="flex flex-col h-full">
        <Suspense>
          <EventTypeList />
        </Suspense>
      </div>

      <div class="flex flex-col w-full h-full gap-2">
        <EventTypeDetail />
        <EventHistory />
      </div>
    </div>
</template>

