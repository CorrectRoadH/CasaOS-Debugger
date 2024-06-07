<script setup lang="ts">
import { useDebuggerStore } from "@/stores/debugger";

const props = defineProps<{
  sourceID: string;
  eventType?: string; 
}>()

</script>

<template>
    <div class="flex flex-col w-full bg-slate-200 p-2 rounded-lg">
        <div class="font-black">事件定义</div>
        <div>
            <div>{{props.eventType}}</div>
           
            <div>
                <div class="font-black">属性定义</div>
                <div>
                    <!-- {{ 
                        debuggerStore.selectedEventType?.propertyTypeList
                    }} -->
                    <!-- <div v-for="item in debuggerStore.selectedEventType.propertyTypeList">
                        <div>item.name</div>
                    </div>  -->
                    <!-- <div v-if="debuggerStore.selectedEventType?.propertyTypeList?.length==0">
                        <div>❗ 警告! Event应该有Propertise定义。请修改代码 </div>
                        <div>无定义</div>
                    </div> -->
                </div>
            </div>
        </div>
    </div>
</template>