<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import HomeView from './views/HomeView.vue';

import Layout from './components/kit/Layout.vue';
import BasePanel from './components/kit/BasePanel.vue';

import SourceList from './components/SourceList.vue';
</script>

<template>
  <Layout class="flex w-full h-screen overflow-hidden" id="sidebar" :show-right-sidebar="false"
    :is-mobile="false" :show-left-sidebar="true">
    <template #sidebar>
      <Suspense> 
        <SourceList />
      </Suspense>
    </template>

    <template #content>
      <BasePanel id="panel1" :highlight-header="false">
        <template #content> 
          <div class="flex flex-col">
            <Suspense> 
              <!-- <HomeView /> -->
              <RouterView />
            </Suspense>  
          </div>
        </template>
      </BasePanel>
    </template>
  </Layout>
</template>

<style scoped>
</style>
