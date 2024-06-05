<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import HomeView from './views/HomeView.vue';
import SourceList from './components/SourceList.vue';
</script>

<template>
      
  <div class="flex w-screen h-screen overflow-hidden p-2"> 
    <div class="flex flex-col w-1/12 p-2">
      <Suspense> 
        <SourceList />
      </Suspense>

    </div>

    
    <div class="flex flex-col w-11/12 bg-slate-100 shadow-lg rounded-lg  p-2">
      <Suspense> 
        <HomeView />
      </Suspense>  
    </div>
  </div>
</template>

<style scoped>
</style>
