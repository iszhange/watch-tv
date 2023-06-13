import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// import * as path from 'path'
import electron from 'vite-plugin-electron'
// import electronRenderer from "vite-plugin-electron/renderer";
// import polyfillExports from "vite-plugin-electron/polyfill-exports";



// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    electron([
      {
        entry: 'electron/main.ts',
      },
      {
        entry: 'electron/preloads/index.ts',
        onstart(options) {
          options.reload()
        }
      }
    ]),
  ],
  build: {
    emptyOutDir: false,
  }
})
