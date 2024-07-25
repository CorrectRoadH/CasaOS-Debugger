import createClient from "openapi-fetch";
import type { paths } from "../../api/openapi" 
import { computed, ref, watch, type Ref } from "vue";


const client = createClient<paths>({ baseUrl: "/v2/debugger/" });

const useEvent = (serviceName:string,eventTypeName:string) => {
    const isLoading = ref<boolean>(true)


    let history = ref<any[]>([])
    // @ts-ignore
    const sourceID = ServiceMap[serviceName.value].eventName


    const fetchHistory = async () => {
        const eventType = eventTypeName == "all" ? undefined : eventTypeName;
        isLoading.value = true
        const { data, error } = await client.GET("/events",{
            params: {
                query: {
                    sourceId: sourceID,
                    eventType,
                    offset: 0,
                    length: 100
                }  
            }
        })
        // @ts-ignore
        history.value = data?.data
        isLoading.value = false
    }

    watch([eventTypeName,serviceName], () => {
        fetchHistory()
    },{immediate: true})

    return {
        isLoading,
        history
    }
}

export default useEvent;