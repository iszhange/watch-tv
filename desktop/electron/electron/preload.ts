import { contextBridge, ipcRenderer } from 'electron'

contextBridge.exposeInMainWorld('WatchTVAPI', {
  countries: (params) => ipcRenderer.invoke('net:api:countries', params)
})