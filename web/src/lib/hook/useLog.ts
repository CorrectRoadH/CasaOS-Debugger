import createClient from "openapi-fetch";
import type { paths } from "../../api/openapi" 
import { computed, ref, watch, type Ref } from "vue";

interface LogData {
    timestamp: string;
    level: string;
    message: string;
    data: Record<string, any>;
    raw: string;
}

function parseLogString(raw: string): LogData|string {
    const timestampRegex = /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}Z/;
    const levelRegex = /\b(info|error|warn|debug|trace)\b/;
    const dataRegex = /\{.*\}$/;

    const timestampMatch = raw.match(timestampRegex);
    const levelMatch = raw.match(levelRegex);
    const dataMatch = raw.match(dataRegex);

    if (!timestampMatch || !levelMatch || !dataMatch) {
      return raw
    }

    const timestamp = timestampMatch[0];
    const level = levelMatch[0];
    const message = raw.substring(timestamp.length + level.length + 2, raw.indexOf(dataMatch[0])).trim();
    const data = JSON.parse(dataMatch[0]);

    return {
        timestamp,
        level,
        message,
        data,
        raw
    };
}

const useLog = (level: Ref<string>, logName: Ref<string>) => {
    const logs = ref<(LogData | string)[]>([]);
    const error = ref<string | null>(null);
    const isLoading = ref<boolean>(true)

    const fetchLogs = async () => {
        try {
            isLoading.value = true
            const client = createClient<paths>({ baseUrl: "/v2/debugger/" });
            const response = await client.GET("/log", {
                params: {
                    query: {
                        service: logName.value,
                        offset: 0,
                        length: 100,
                    }
                }
            });

            if (response.data) {
                logs.value = response.data.data?.map((log: string) => parseLogString(log)).filter((item: string | LogData) => {
                    if (level.value === 'all') {
                        return true;
                    }
                    return typeof item === 'string' || (item as LogData).level === level.value;
                }) || [];
            } else {
                error.value = response.error?.message || 'Unknown error';
            }
        } catch (e) {
            error.value = (e as Error).message;
        }
        isLoading.value = false
    };

    watch([logName, level], () => {
        fetchLogs();
    }, { immediate: true });

    return {
        logs: logs,
        error: error,
        isLoading
    };
}

export default useLog;