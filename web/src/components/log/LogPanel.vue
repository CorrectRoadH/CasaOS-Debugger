<template>
<main class="flex flex-col gap-2 h-full">
  <div v-if="!error">
    {{error}}
  </div>

  <Dropdown v-model="selectedLevel" :options="levelOptions"  placeholder="select a level" checkmark :highlightOnSelect="false" class="w-full md:w-14rem" />


  <div class="flex flex-col gap-2">
    
    <div v-if="isLoading">Loading</div>
    <div v-for="item in logs" v-else class="p-card rounded-lg p-4"
    :class="{
      'bg-red-500': typeof item !== 'string' && item.level === 'error',
      'bg-yellow-500': typeof item === 'string' || item.level === 'warn',
    }"
    >
      <div v-if="typeof item === 'string'">
        <div class="bg-yellow-200 text-black p-2 rounded-lg">警告❗ 应该使用logger来输出结构化的日志</div>
        {{item}}
      </div>

      <div v-else>
        <TabView>
          <TabPanel header="结构化">
            <div>时间: {{item.timestamp}}</div>
            <div>级别: {{item.level}}</div>
            <div>内容: {{item.message}}</div>
            <json-viewer class="!bg-white" :value="item.data" copyable></json-viewer>
          </TabPanel>
          <TabPanel header="Raw">
              {{item.raw}}
          </TabPanel>
        </TabView>
      </div>
    </div>
  </div> 
</main>
</template>


<script setup lang="ts">
import ServiceMap from '@/lib/utils';
// @ts-ignore
import JsonViewer from 'vue-json-viewer'
import TabView from 'primevue/tabview';
import TabPanel from 'primevue/tabpanel';
import Dropdown from 'primevue/dropdown';
import { ref, watch } from 'vue';
import useLog from '@/lib/hook/useLog';


const props = defineProps<{
  serviceName: string;
  eventType: string; 
}>()

// @ts-ignore
const logName = ref(ServiceMap[props.serviceName].logName);

watch(() => props.serviceName, () => {
  console.log('props.serviceName',props.serviceName);
  // @ts-ignore
  logName.value = ServiceMap[props.serviceName].logName;
});


const selectedLevel = ref<string>('all');
const levelOptions = ['all', 'info', 'warn', 'error', 'debug', 'trace'];
const { logs,error,isLoading }= useLog(selectedLevel,logName); 

</script>