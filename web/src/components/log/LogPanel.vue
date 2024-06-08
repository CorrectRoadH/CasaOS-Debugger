<template>
<main class="flex gap-2 h-full">
  <div v-if="!error">
    {{error}}
  </div>

  <div>
    <!-- <logViewer :log="data?.data[0]" :loading="false" /> -->
     {{ data?.data[0] }}
  </div>
</main>
</template>


<script setup lang="ts">
import ServiceMap from '@/lib/utils';
import createClient from "openapi-fetch";
import type { paths } from "../../api/openapi" 
// import logViewer from "@femessage/log-viewer"

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
          }  
      }
})
</script>