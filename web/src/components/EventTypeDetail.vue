<script setup lang="ts">
import { useDebuggerStore } from "@/stores/debugger";

const debuggerStore = useDebuggerStore();

</script>

<template>
    <div class="flex flex-col w-full bg-slate-200 p-2 rounded-lg">
        <div>事件定义</div>
        <div>
            <div>{{debuggerStore.selectedEventType.name}}</div>
            <div>
            属性定义
            <div>
                {{ 
                    debuggerStore.selectedEventType?.propertyTypeList
                 }}
            </div>
            </div>
        </div>
    </div>
</template>