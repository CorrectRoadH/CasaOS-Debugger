<script setup lang="ts">
import { useDebuggerStore } from "@/stores/debugger";

const debuggerStore = useDebuggerStore();

</script>

<template>
    <div class="flex flex-col bg-slate-200 p-2 rounded-lg">
        <div>事件类型</div>
        <div>
            <div>{{debuggerStore.selectedEventType.name}}</div>
            <div>
            属性
            <div>
                {{ 
                    debuggerStore.selectedEventType?.propertyTypeList
                 }}
            </div>
            </div>
        </div>
    </div>
</template>