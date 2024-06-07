<template>
    <div
      class="my-1 min-h-[58px] overflow-hidden rounded-xl border border-gray-500 transition-all duration-300 ease-in-out"
      ref="mainPanel"
      :class="flexClass"
    >
      <ScrollPanel>
        <!-- Header Start -->
        <header
          class="top-0 z-10 flex items-center w-full px-4 transition duration-200 border-b shrink-0 backdrop-blur"
          v-if="$slots.header"
          :class="[
            { sticky: !d_collapsed },
            highlightHeader ? 'border-gray-500 bg-gray-200/80' : 'border-white/80 bg-white/80',
            extendHeight ? 'h-[104px]' : 'h-14',
          ]"
        >
          <slot name="header">
            <div class="grow">Header</div>
            <div class="justify-self-end" @click="toggle">toggle</div>
          </slot>
        </header>
        <!-- Header End -->
  
        <div
          class="relative elements grow"
          v-show="!d_collapsed"
          :id="id + '_content'"
          role="region"
          :aria-labelledby="id + '_header'"
        >
          <!-- Content Start -->
          <slot name="content" :grid="grid"></slot>
          <!-- Content End -->
          <!-- Footer Start -->
          <div v-if="$slots.footer">
            <slot name="footer"></slot>
          </div>
          <!-- Footer End -->
        </div>
      </ScrollPanel>
    </div>
  </template>
  
<script setup lang="ts">
import { useResizeObserver } from "@vueuse/core";
import ScrollPanel from "primevue/scrollpanel";
import { computed, ref, useAttrs, watch } from "vue";
  
  // Props
  const props = defineProps({
    toggleable: Boolean,
    collapsed: {
      type: Boolean,
      default: false,
    },
    highlightHeader: {
      type: Boolean,
      default: false,
    },
    extendHeight: {
      type: Boolean,
      default: false,
    },
  });
  
  // Emit
  const emit = defineEmits(["toggle", "update:collapsed"]);
  
  // Data
  const itemWidth = 144;
  const attrs = useAttrs();
  const id = attrs.id;
  const mainPanel = ref();
  const scrollPanel = ref();
  const scrollPanelStyle = ref({
    height: "auto",
  });
  const grid = ref("grid-cols-6");
  const d_collapsed = ref(props.collapsed);
  
  // Computed
  const flexClass = computed(() => {
    return d_collapsed.value ? "h-[58px]" : "flex-1";
  });
  
  // Watchs
  watch(
    () => props.collapsed,
    (value) => {
      d_collapsed.value = value;
    },
  );
  
  // Methods
  const toggle = (event: MouseEvent) => {
    d_collapsed.value = !d_collapsed.value;
    emit("update:collapsed", d_collapsed.value);
    emit("toggle", { originalEvent: event, value: d_collapsed.value });
  };
  
  const scrollTop = (y: number) => {
    scrollPanel.value.scrollTop(y);
  };
  
  // Hooks
  useResizeObserver(mainPanel, (entries) => {
    const entry = entries[0];
    const { width, height } = entry.contentRect;
    scrollPanelStyle.value = {
      height: `${height}px`,
    };
    const col = Math.floor((width - 16) / itemWidth);
    grid.value = `grid-cols-${col}`;
  });
  
  // Expose
  defineExpose({
    scrollTop,
    toggle,
  });
</script>