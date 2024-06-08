<template>
<div class="flex flex-col h-full w-full p-4">
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
    <template #title>不正确的事件</template>
    <template #content>
      WIP
    </template>
  </Card>

  <Divider />

  <Card>
    <template #title>报错的日志</template>
    <template #content>
      WIP
    </template>
  </Card>


</div>
</template>
<script setup lang="ts">
import Divider from 'primevue/divider';
import Card from 'primevue/card';

</script>