import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useDebuggerStore = defineStore('debugger', () => {
  const selectedSourceID = ref("")
  const selectedEventType = ref({
    name:"all",
    propertyTypeList:[
      {name:""}
    ]
  })

  return { selectedSourceID,selectedEventType }
})
