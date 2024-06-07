<template>
<div class="flex flex-col h-full w-full p-2">
  <div class="flex w-full">
    <div class="flex flex-col mx-auto">
      <img class="w-48 h-48" src="/logo.webp" >
      <div class="mx-auto font-black"> CasaOS Debugger </div>
    </div>
    <div></div>
  </div>

  <Divider />

  <div class="">
  </div>

  <Card>
    <template #title>Tips</template>
    <template #content>
        <p class="m-0">
          请选择左侧的服务以开始
        </p>
    </template>
  </Card>

  <Divider />

  <Card>
    <template #title>错误中心</template>
    <template #content>
      <div>
        错误的事件
      </div>
      <div>
        错误的日志
      </div>
  
    </template>
  </Card>

</div>
</template>
<script setup lang="ts">
import Divider from 'primevue/divider';
import Card from 'primevue/card';

</script>