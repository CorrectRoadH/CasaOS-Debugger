<script setup lang="ts">
import EventTypeList from "@/components/EventTypeList.vue"
import EventTypeDetail from "@/components/EventTypeDetail.vue";
import EventHistory from "@/components/EventHistory.vue";

</script>

<template>
  <main>
    <div class="flex m-auto">
      <!-- header -->
      <div>事件</div>
      <div>日志</div>
    </div>

    <div class="flex gap-2">
      <div class="flex flex-col bg-slate-400 rounded-lg p-2">
        <div class="p-2">all</div>
        <Suspense>
          <EventTypeList />
        </Suspense>
      </div>

      <div class="flex flex-col gap-2">
        <EventTypeDetail />
        <EventHistory />
      </div>
    </div>
  </main>
</template>

