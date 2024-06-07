
<template>
    <div class="flex w-full h-full overflow-hidden lg:pr-3">
      <!-- Sidebar Start -->
      <div class="ease -mr-1 flex h-screen w-0 flex-col overflow-y-hidden bg-[#F2F4F5] transition-all duration-200"
        id="menu-bar" :class="[
          { 'fixed -left-64 top-0 z-[4000] w-64': isMobile && !showLeftSidebar },
          { 'fixed -left-64 top-0 z-[4000] w-64 translate-x-64': isMobile && showLeftSidebar },
          { 'lg:w-64': !isMobile },
        ]" @click="onSideBarClick">
        <slot name="sidebar"></slot>
      </div>
      <!-- Sidebar End -->
  
      <!-- Content Area Start -->
      <div class="relative flex flex-col flex-1 w-0 mx-2 my-1 overflow-hidden duration-200 ease lg:mx-1">
        <slot name="content"></slot>
      </div>
      <!-- Content Area End -->
  
      <!-- Right Sidebar Start -->
      <div class="flex flex-col my-1 overflow-hidden transition-all duration-200 ease" id="right-sidebar" :class="[
          { 'invisible mx-0 w-0 opacity-0': (!showRightSidebar && !isMobile) || isMobile },
          { 'visible mx-1 w-72 opacity-100': showRightSidebar && !isMobile },
        ]">
        <slot name="rightbar"></slot>
      </div>
      <!-- Right Sidebar End -->
    </div>
</template>
  
<script setup lang="ts">
  defineProps({
    showRightSidebar: {
      type: Boolean,
      default: false,
    },
    isMobile: {
      type: Boolean,
      default: false,
    },
    showLeftSidebar: {
      type: Boolean,
      default: false,
    },
  });
  
  // Emit
  const emit = defineEmits(["onSidebarClick"]);
  
  const onSideBarClick = (e: MouseEvent) => {
    emit("onSidebarClick", e);
  }
  
</script>
  