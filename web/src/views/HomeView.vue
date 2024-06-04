<script setup lang="ts">
import EventTypeList from "@/components/EventTypeList.vue"
import EventTypeDetail from "@/components/EventTypeDetail.vue";
import EventHistory from "@/components/EventHistory.vue";

</script>

<template>
    <div class="flex m-auto">
      <!-- header -->
      <div>事件</div>
      <div>日志</div>
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

