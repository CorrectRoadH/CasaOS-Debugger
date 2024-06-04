import { fileURLToPath, URL } from 'node:url'

import { CommonServerOptions, defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import VueDevTools from 'vite-plugin-vue-devtools'
import { ui } from "../raw/usr/share/casaos/modules/casaos_debugger.json";

const devBase = ui.entry.substring(0, ui.entry.lastIndexOf("/"));
const prodBase = "./";

// Enable the proxy for the development server.
const useProxy = false;
// Enable the global proxy for the development server.
const useGlobalProxy = true;

// Custom proxy settings
const proxy = {
  // "^/chat.*": { target: "http://10.0.0.65:8001", changeOrigin: true },
} as CommonServerOptions["proxy"];

// globalProxyTarget:
// The target server for the global proxy. except the custom proxy settings.
const globalProxyTarget = "http://10.0.0.83";

const globalProxy = (() => {
  const result = {} as CommonServerOptions["proxy"] || {}; // Initialize result as an empty object
  const excludePaths = [
    devBase,
    ...Object.keys(proxy || {}).map((key) => key.replace("^", "").replace(".*", "")),
  ].join("|");
  result[`^(?!${excludePaths})/.*`] = {
    target: globalProxyTarget,
    changeOrigin: true,
    ws: true,
  };
  return result;
})();

const isDevMode = process.env.NODE_ENV === "development";
const base = isDevMode ? devBase : prodBase;

// https://vitejs.dev/config/
export default defineConfig({
  base,
  plugins: [
    vue(),
    vueJsx(),
    VueDevTools(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {
    proxy: {
      ...(useGlobalProxy ? globalProxy : {}),
      ...(useProxy ? proxy : {}),
    },
  },
})
