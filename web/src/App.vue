<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import HelloWorld from './components/HelloWorld.vue'
import HomeView from './views/HomeView.vue';
import SourceList from './components/SourceList.vue';
</script>

<template>
      
  <div class="flex w-screen h-screen overflow-hidden p-2"> 
    <div class="flex flex-col w-1/12 p-2">
      <!-- <div class="p-2">tab</div>
      <div class="bg-gray-300 rounded-lg p-2"> casaos app management </div>
      <div class="p-2"> zimaos </div> -->
      <Suspense> 
        <SourceList />
      </Suspense>

    </div>

    
    <div class="w-11/12 bg-gray-300 rounded-lg overflow-hidden p-2">
      <Suspense> 
        <HomeView />
      </Suspense>  
    </div>
  </div>
  <!-- <RouterView /> -->
</template>

<style scoped>
</style>
