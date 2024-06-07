<script setup lang="ts">
import { useDebuggerStore } from "@/stores/debugger";

const props = defineProps<{
  sourceID: string;
  eventType: string; 
}>()

const debuggerStore = useDebuggerStore();

</script>

<template>
    <div class="flex flex-col w-full p-card	p-2 rounded-lg">
        <div class="font-black">事件定义</div>
        <div>
            <div>{{debuggerStore.selectedEventType.name}}</div>
           
            <div>
                <div class="font-black">属性定义</div>
                <div>
                    <div v-for="item in debuggerStore.selectedEventType.propertyTypeList">
                        <div>{{item.name}}</div>
                    </div> 
                    <div v-if="debuggerStore.selectedEventType?.propertyTypeList?.length==0">
                        <div class="bg-red-200 p-2 rounded-lg">❗ 警告! Event应该有Propertise定义。请修改代码 </div>
                        <div>无定义</div>
                    </div>
                    <div class="bg-yellow-200 p-2 rounded-lg">❗ 注意，如果这里的字段与下面记录对不上，请提醒相关开发人员修正schema</div>
                </div>
            </div>
        </div>
    </div>
</template>