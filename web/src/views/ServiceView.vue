<script setup lang="ts">
import TabView from 'primevue/tabview';
import TabPanel from 'primevue/tabpanel';

import EventTypeList from "@/components/event/EventTypeList.vue"
import EventTypeDetail from "@/components/event/EventTypeDetail.vue";
import EventHistory from "@/components/event/EventHistory.vue";
import LogPanel from '@/components/log/LogPanel.vue';

const props = defineProps<{
  serviceName: string;
  eventType: string; 
}>()

</script>

<template>
  <TabView>
    <TabPanel header="事件">
      <div class="flex gap-2 h-full" :key="props.serviceName">
        <div class="flex flex-col h-full">
          <Suspense>
            <EventTypeList :serviceName="props.serviceName" :eventType="props.eventType" />
          </Suspense>
        </div>
  
        <div class="flex flex-col w-full h-full gap-2 p-4">
          <EventTypeDetail :serviceName="props.serviceName" :eventType="props.eventType" />
          <EventHistory :serviceName="props.serviceName" :eventType="props.eventType" />
        </div>
      </div>
      </TabPanel>
    <TabPanel header="日志">
      <LogPanel :serviceName="props.serviceName" :eventType="props.eventType" />
    </TabPanel>
  </TabView>
</template>

