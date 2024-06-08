<template>
<main class="flex flex-col gap-2 h-full">
  <div v-if="!error">
    {{error}}
  </div>

  <Dropdown v-model="selectedLevel" :options="levelOptions"  placeholder="select a level" checkmark :highlightOnSelect="false" class="w-full md:w-14rem" />


  <div class="flex flex-col gap-2">

    <div v-for="item in logs"  class="p-card rounded-lg p-4"
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
import createClient from "openapi-fetch";
import type { paths } from "../../api/openapi" 
// @ts-ignore
import JsonViewer from 'vue-json-viewer'
import TabView from 'primevue/tabview';
import TabPanel from 'primevue/tabpanel';
import Dropdown from 'primevue/dropdown';
import { computed, ref } from 'vue';

interface LogData {
    timestamp: string;
    level: string;
    message: string;
    data: Record<string, any>;
    raw: string;
}

function parseLogString(raw: string): LogData|string {
    const timestampRegex = /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}Z/;
    const levelRegex = /\b(info|error|warn|debug|trace)\b/;
    const dataRegex = /\{.*\}$/;

    const timestampMatch = raw.match(timestampRegex);
    const levelMatch = raw.match(levelRegex);
    const dataMatch = raw.match(dataRegex);

    if (!timestampMatch || !levelMatch || !dataMatch) {
      return raw
    }

    const timestamp = timestampMatch[0];
    const level = levelMatch[0];
    const message = raw.substring(timestamp.length + level.length + 2, raw.indexOf(dataMatch[0])).trim();
    const data = JSON.parse(dataMatch[0]);

    return {
        timestamp,
        level,
        message,
        data,
        raw
    };
}

const props = defineProps<{
  serviceName: string;
  eventType: string; 
}>()

// @ts-ignore
const logName = ServiceMap[props.serviceName].logName;
const client = createClient<paths>({ baseUrl: "/v2/debugger/" });

const { data, error } = await client.GET("/log",{
      params: {
          query: {
            service: logName,
            offset: 0,
            length: 100,
          }  
      }
})


const selectedLevel = ref<string>('all');
const levelOptions = ['all', 'info', 'warn', 'error', 'debug', 'trace'];

const logs = computed(() => {
  return data?.data?.map((log: string) => parseLogString(log)).filter((item:string|LogData)=>{
    if (selectedLevel.value === 'all') {
      return true;
    }
    return typeof item === 'string' || item.level === selectedLevel.value;
  }) || [];
})

</script>