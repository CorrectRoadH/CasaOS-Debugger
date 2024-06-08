/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */


export interface paths {
  "/": {
    /** Get information about the host */
    get: operations["HelloWorld"];
  };
  "/sources": {
    /** Get all sources */
    get: operations["GetAllSources"];
  };
  "/log": {
    /** Query log */
    get: operations["QueryLog"];
  };
  "/events": {
    /** Query messages */
    get: operations["GetAllMessages"];
  };
  "/events/types": {
    /** return event type list by sourceId */
    get: operations["GetAllEventType"];
  };
}

export type webhooks = Record<string, never>;

export interface components {
  schemas: {
    base_response: {
      /** @description message returned by server side if there is any */
      message?: string;
    };
    sourceID: string;
    /** @description Event Type */
    EventType: Record<string, never>;
    /** @description Event */
    Event: Record<string, never>;
    Log: string;
  };
  responses: {
    /** @description OK */
    response_ok: {
      content: {
        "application/json": components["schemas"]["base_response"];
      };
    };
    /** @description Conflict */
    response_conflict: {
      content: {
        "application/json": components["schemas"]["base_response"];
      };
    };
    /** @description Internal Server Error */
    response_internal_server_error: {
      content: {
        "application/json": components["schemas"]["base_response"];
      };
    };
    /** @description Not Found */
    response_not_found: {
      content: {
        "application/json": components["schemas"]["base_response"];
      };
    };
    /** @description Bad Request */
    response_bad_request: {
      content: {
        "application/json": components["schemas"]["base_response"];
      };
    };
    /** @description Unauthorized */
    response_unauthorized: {
      content: {
        "application/json": components["schemas"]["base_response"];
      };
    };
    /** @description OK */
    response_get_source_list_ok: {
      content: {
        "application/json": components["schemas"]["base_response"] & {
          data?: components["schemas"]["sourceID"][];
        };
      };
    };
    /** @description OK */
    response_get_event_type_list_ok: {
      content: {
        "application/json": components["schemas"]["base_response"] & {
          data?: components["schemas"]["EventType"][];
        };
      };
    };
    /** @description OK */
    response_query_message_ok: {
      content: {
        "application/json": components["schemas"]["base_response"] & {
          data?: components["schemas"]["Event"][];
        };
      };
    };
    /** @description OK */
    response_query_log_ok: {
      content: {
        "application/json": components["schemas"]["base_response"] & {
          data?: components["schemas"]["Log"][];
        };
      };
    };
  };
  parameters: {
    name: string;
  };
  requestBodies: never;
  headers: never;
  pathItems: never;
}

export type $defs = Record<string, never>;

export type external = Record<string, never>;

export interface operations {

  /** Get information about the host */
  HelloWorld: {
    responses: {
      200: components["responses"]["response_ok"];
      500: components["responses"]["response_internal_server_error"];
    };
  };
  /** Get all sources */
  GetAllSources: {
    responses: {
      200: components["responses"]["response_get_source_list_ok"];
      500: components["responses"]["response_internal_server_error"];
    };
  };
  /** Query log */
  QueryLog: {
    parameters: {
      query: {
        /** @example casaos-installer */
        service: string;
        /** @example 0 */
        offset: number;
        /** @example 10 */
        length: number;
      };
    };
    responses: {
      200: components["responses"]["response_query_log_ok"];
      500: components["responses"]["response_internal_server_error"];
    };
  };
  /** Query messages */
  GetAllMessages: {
    parameters: {
      query: {
        /** @example casaos */
        sourceId?: string;
        /** @example app:install */
        eventType?: string;
        /** @example 0 */
        offset: number;
        /** @example 10 */
        length: number;
      };
    };
    responses: {
      200: components["responses"]["response_query_message_ok"];
      500: components["responses"]["response_internal_server_error"];
    };
  };
  /** return event type list by sourceId */
  GetAllEventType: {
    parameters: {
      query: {
        /** @example 847d7fde */
        sourceId: string;
      };
    };
    responses: {
      200: components["responses"]["response_get_event_type_list_ok"];
      500: components["responses"]["response_internal_server_error"];
    };
  };
}
